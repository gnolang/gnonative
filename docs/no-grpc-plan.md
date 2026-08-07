# Add a buf/grpc(connect)-free option to gnonative

## Context

Today every gnonative consumer is forced through buf/connect-gRPC:

- **Go**: all 48 RPC handlers (40 unary + 8 server-streaming) are implemented directly on `gnoNativeService` in `service/api.go` with connect signatures; business logic is inline in the handlers. There is no plain-function API.
- **Mobile**: `framework/service/bridge.go` starts an in-process connect HTTP server (UDS/TCP) *and* a loopback connect client; every JS call does a full connect round-trip in-process via the reflection dispatcher in `framework/service/service_client.go`.
- **Expo/JS**: `@gnolang/gnonative` uses `@connectrpc/connect(-web)` + `@bufbuild/protobuf` for serialization/dispatch, even though the actual transport is already just `GoBridge.invokeGrpcMethod(methodName, jsonString) → base64(json)`.

Goal: let apps opt out of buf/grpc entirely while keeping the existing grpc path 100% backward compatible. The gRPC handlers become thin wrappers over plain Go API functions; the expo package gains a `@gnolang/gnonative/native` entry point that never touches buf/connect; a new example app `examples/js/expo/hello-native` demonstrates it.

**User decisions**: keep `@bufbuild/*`/`@connectrpc/*` in `dependencies` (no breaking change; Metro won't bundle them for `/native` apps). Naming: `/native` subpath + `hello-native` example.

**Verified detail that drives the design**: the current bridge JSON encoding is asymmetric — requests `protojson.Unmarshal` (service_client.go:79), responses `encoding/json.Marshal` of pb structs (service_client.go:98, snake_case, int64→number). The new path uses **protojson both directions** (camelCase, int64→string, bytes→base64 string, zero values omitted). The legacy path is left untouched.

---

## Phase 1 — Go: extract plain API (`service/`)

### `service/api.go` (modify in place)
Convert every handler to a plain signature, keeping the proto method names on `*gnoNativeService`:

```go
// unary (40)
func (s *gnoNativeService) SetRemote(ctx context.Context, req *api_gen.SetRemoteRequest) (*api_gen.SetRemoteResponse, error)
// server-streaming (8): Call, Send, Run, CreateSession, RevokeSession,
// RevokeAllSessions, BroadcastTxCommit, HelloStream
func (s *gnoNativeService) Call(ctx context.Context, req *api_gen.CallRequest, send func(*api_gen.CallResponse) error) error
```

Mechanical per method: `req.Msg.X` → `req.X`; `connect.NewResponse(&Y{...})` → `&Y{...}`; `stream.Send(...)` → `send(...)`. Drop the `connectrpc.com/connect` import. Existing helpers (`ConvertKeyInfo`, `convert*`, `getGrpcError`, `s.ClientSignTx`, `s.estimateGasWanted`) are already connect-free — untouched.

Behavioral fix while here: make `HelloStream` respect `ctx` (`select { case <-ctx.Done(): return ctx.Err(); case <-time.After(2*time.Second): }`) so a closed direct stream doesn't leak its goroutine.

### New `service/api_connect.go`
Thin connect wrappers implementing `_goconnect.GnoNativeServiceHandler` (48 wrappers, 2–5 lines each):

```go
type connectHandler struct{ svc *gnoNativeService }
var _ _goconnect.GnoNativeServiceHandler = (*connectHandler)(nil)

func (h *connectHandler) SetRemote(ctx context.Context, req *connect.Request[api_gen.SetRemoteRequest]) (*connect.Response[api_gen.SetRemoteResponse], error) {
    res, err := h.svc.SetRemote(ctx, req.Msg)
    if err != nil { return nil, err }
    return connect.NewResponse(res), nil
}
func (h *connectHandler) Call(ctx context.Context, req *connect.Request[api_gen.CallRequest], stream *connect.ServerStream[api_gen.CallResponse]) error {
    return h.svc.Call(ctx, req.Msg, stream.Send)
}
```
Error values (`ErrCode` from `api/gen/go/error.go`) pass through unchanged → identical wire behavior.

### New `service/api_interface.go`
`GnoNativeApi` interface listing all 48 plain methods.

### `service/service.go` (modify)
- Embed `GnoNativeApi` into the exported `GnoNativeService` interface (alongside `GetUDSPath/GetTcpAddr/GetTcpPort/Close`). Pure-Go consumers can now call the API directly, with both listeners disabled (`WithDisableUdsListener()` + no `WithUseTcpListener()` already yields "no server").
- In `runGRPCServer`: register `_goconnect.NewGnoNativeServiceHandler(&connectHandler{svc: s}, compress1KB)` instead of `s`.

No changes to `goserver/` or `examples/go/goclient`.

---

## Phase 2 — `framework/service/`: direct dispatcher (bypasses connect)

### New `framework/service/dispatcher.go`
Hand-written dispatch maps + small generic helpers (no reflection):

```go
func unary[...](fn func(context.Context, PReq) (PRes, error)) unaryHandler       // protojson in/out
func streaming[...](fn func(context.Context, PReq, func(PRes) error) error) streamHandler

unary: map[string]unaryHandler{ "SetRemote": unary(svc.SetRemote), /* 40 entries */ }
streams: map[string]streamHandler{ "Call": streaming(svc.Call), /* 8 entries */ }
```

### New `framework/service/service_dispatch.go`
Gomobile-facing promise API (exported, mirrors `ServiceClient` — pattern reference: `service_client.go`):
- `InvokeMethodWithPromiseBlock(promise, method, jsonMessage)` → resolves `base64(protojson)`.
- `CreateStreamWithPromiseBlock` / `StreamReceiveWithPromiseBlock` / `CloseStreamWithPromiseBlock`.

Stream adapter (push→pull): on create, run the plain streaming method in a goroutine with a cancellable context; `send` blocks on a cap-1 `chan []byte` (backpressure); on method return deliver `io.EOF` (raw, not wrapped — JS matches the literal `'EOF'` message, transport_native.ts:164) or the error. Registry: map + atomic id counter, same as `serviceClient`.

### `framework/service/bridge.go` (modify)
- `BridgeConfig`: add `DisableGrpcServers bool` (zero-value = today's behavior).
- When set: pass `service.WithDisableUdsListener()`, skip TCP, skip the loopback connect client; `ServiceClient` becomes a stub whose methods reject with "grpc servers disabled; use InvokeMethod".
- Always attach the dispatcher: `b.ServiceDispatcher = newServiceDispatcher(b.serviceServer)` (works in both modes; embed the interface in `Bridge` like `ServiceClient`).

### New `framework/service/dispatcher_test.go`
- **Completeness test**: iterate `api_gen.File_rpc_proto.Services().Get(0).Methods()`; assert every method appears in exactly one map with matching streaming-ness (`IsStreamingServer()`). This is the drift guard between the two paths.
- Smoke tests using `service.NewGnoNativeService(service.WithDisableUdsListener())` (no sockets): `Hello` unary, `HelloStream` (messages then EOF), one error case asserting the `ErrCode` string, protojson camelCase output shape.

---

## Phase 3 — Native modules (additive)

`expo/ios/GnonativeModule.swift` and `expo/android/src/main/java/land/gno/gnonative/GnonativeModule.kt`, both:
- `AsyncFunction("initBridgeWithOptions")` — like `initBridge` but sets `config.disableGrpcServers` from the JS option; when grpc disabled, skip the iOS-simulator force-TCP block (no sockets needed at all). `initBridge` untouched.
- `AsyncFunction("invokeMethod")`, `createStream`, `streamReceive`, `closeStream` — copies of the existing four calling the new Bridge dispatcher methods.

`PromiseBlock.swift` / `JavaPromiseBlock.kt` reused as-is (`io.EOF` surfaces as message `EOF` through the same `CallReject` plumbing).

---

## Phase 4 — TS types (generated, buf-free)

1. `buf.gen.yaml`: es plugin opt → `target=ts,json_types=true`. `make api.generate` → `api/gen/es/*_pb.ts` gain pure `export type XxxJson = {...}` aliases (int64→`string`, bytes→base64 `string`) and the plain `ErrCode` enum. Existing exports unchanged → grpc path unaffected.
2. New `expo/scripts/extract-json-types.mjs` (node, no deps): extract all `*Json` type blocks + `ErrCode` enum + doc comments from the generated files into `expo/src/native/types.gen.ts` (zero imports, "generated" header, checked in).
3. `expo/Makefile`: run the script as part of `build.api` (after vendoring `api/gen/es` → `expo/src/api/vendor/`).

---

## Phase 5 — Expo JS: `@gnolang/gnonative/native`

### New `expo/src/native/` (must not import `../grpc/*`, `../api/vendor/*`, or `../index`)
- `types.gen.ts` — from Phase 4.
- `types.ts` — `Config`, `BridgeStatus`, `GnoNativeClientApi` interface (mirror of `GnoKeyApi` in `src/api/types.ts`, Json-typed: bytes are base64 strings, int64 accept/return `string`).
- `encoding.ts` — base64↔string helpers via `base64-js` + `fast-text-encoding` (already deps of both paths).
- `error.ts` — `GnoNativeError`: same `errCode()/hasErrCode()` regex parsing of `Name(#123)` as `src/grpc/error.ts`, minus `ConnectError`, using the `ErrCode` enum from `types.gen.ts`.
- `stream.ts` — `AsyncIterable` over `GoBridge.createStream/streamReceive/closeStream`, terminating on `'EOF'` (same logic as the generator in `src/grpc/transport_native.ts`).
- `client.ts` — `GnoNativeClient implements GnoNativeClientApi`: `initClient()` = `GoBridge.initBridgeWithOptions({useGrpcServers:false})` → `setRemote` → `setChainID` (no ports, no connect). Methods: build camelCase request object, `JSON.stringify` → `invokeMethod` → decode base64 → `JSON.parse`. Method-for-method parity with `GnoNativeApi` (`src/api/GnoNativeApi.ts` is the surface reference) so migration is an import swap.
- `provider.tsx` — `GnoNativeProvider` + `useGnoNativeContext` clones of `src/provider/gnonative-provider.tsx` typed to `GnoNativeClient`.
- `index.ts` — barrel.

### Modified
- `expo/src/GoBridge.ts`: add `initBridgeWithOptions`, `invokeMethod`, `createStream`, `streamReceive`, `closeStream` (file is buf-free, shared by both paths).
- `expo/package.json`: add `exports` map — `"."` → build/index (unchanged behavior), `"./native"` → `build/native/index`, plus `"./package.json"` and a `"./build/*"` escape hatch so existing deep imports keep working under Metro's exports resolution. **Dependencies unchanged** (user decision: keep `@bufbuild/*`/`@connectrpc/*` in `dependencies`).
- `expo/README.md`: document the two paths + the protojson type mapping (camelCase, int64→string, bytes→base64, omitted defaults) + the `/native` selling point: no streaming polyfills needed.

---

## Phase 6 — Example app `examples/js/expo/hello-native/`

Mirror `examples/js/expo/hello` content + `expo/example` local-module wiring:
- `package.json` — expo ~53 / RN 0.79 / React 19 (match `expo/example`); no `@gnolang/gnonative` dep; `"expo": { "autolinking": { "nativeModulesDir": "../../../../expo" } }`; no buf/connect/polyfill packages anywhere.
- `metro.config.js` — copy of `expo/example/metro.config.js` with `extraNodeModules['@gnolang/gnonative']` → `../../../../expo`, matching `watchFolders` + blockList.
- `App.tsx` — same flow as `hello/App.tsx` (listKeyInfo, getRemote, getChainID, generateRecoveryPhrase, addressFromMnemonic, addressToBech32, hello, helloStream loop) but importing from `'@gnolang/gnonative/native'`, and **no polyfill imports**.
- `app.json` (id `land.gno.hellonative`), `babel.config.js`, `tsconfig.json`, `assets/` from `hello`.
- `Makefile` — gnoboard convention (`examples/js/expo/gnoboard/Makefile`): build gomobile framework via root `make framework.ios|framework.android`, copy `GnoCore.xcframework` → `ios/Frameworks` / `gnocore.aar` → `android/libs` (they land inside the linked `expo/` module dir), build the expo module (`cd expo && npm install && npm run prepare`), then `npx expo run:ios|android`; keep an `android.reverse_tcp` helper.
- `README.md` — what it demonstrates, prerequisites, run instructions.

---

## Implementation order

1. Phase 1 → `go build ./... && go test ./service/...` (pure refactor; connect path proven by compile-time `var _ GnoNativeServiceHandler` assert).
2. Phase 2 → dispatcher tests.
3. Phases 3 + 4 (independent).
4. Phase 5 → `cd expo && npm run build`; check `grep -r "@bufbuild\|@connectrpc" expo/build/native` is empty.
5. Phase 6 — end-to-end.

## Risks / gotchas

- **Gomobile rebuild required** for the new JS to work; example Makefile handles it; npm publish must be built after the Go changes.
- **Two-path drift**: guarded by the descriptor completeness test + compile-time handler assert. New-RPC checklist: proto → generate → plain method → connect wrapper → dispatcher entry → TS client method.
- **JSON dialect**: legacy responses are snake_case/int64-number, new path camelCase/int64-string — never mix decoders; documented in README.
- **Metro `exports`**: default-on for RN ≥0.79; `/native` unavailable to very old RN — acceptable (new feature), root entry untouched.
- **Stream lifecycle**: Close cancels the context; HelloStream ctx fix prevents goroutine leaks; cap-1 send channel gives backpressure.
- Pre-existing issues out of scope (flag only): `runGRPCServer` reads `s.server` in a goroutine before assignment; unguarded `s.remote` writes in `SetRemote`.

## Design clarifications (from planning discussion, 2026-07-08)

- **Streaming without gRPC**: fully preserved. The JS↔native boundary is already pull-based (`createStream` → repeated `streamReceive` → reject `'EOF'`); the plain Go streaming methods use a `send` callback, the dispatcher bridges push→pull with a cap-1 channel. App-facing API stays `AsyncIterable`. Note: 7 of 8 streaming RPCs send exactly one message today.
- **Events/EventEmitter considered and rejected** for stream delivery: module-global events would need streamId enveloping + JS-side demux/buffering (rebuilding the pull queue in JS), lose backpressure, and add listener-attach races — for no gain given single-message streams. The seam (`expo/src/native/stream.ts`, `framework/service/service_dispatch.go`) allows swapping to push later without app-code changes if chatty subscriptions ever appear.
- **Protobuf scope**: apps on `/native` have zero protobuf/buf at runtime and in dependencies (types.gen.ts is pure type aliases, plain JSON on the wire). The Go core keeps pb structs + protojson internally (compiled into GnoCore, invisible to apps), and buf.gen.yaml remains the maintainers' dev-time codegen source of truth for both paths — deliberately, to prevent type drift between the two client options.
- **Packaging**: keep `@bufbuild/*`/`@connectrpc/*` in `dependencies` (user decision — no breaking change for grpc-path users; Metro never bundles them for `/native` apps).

## Verification

1. `go build ./... && go vet ./... && go test ./service/... ./framework/...`
2. `make framework.ios` (and `framework.android` if the Android SDK is available).
3. `cd expo && npm install && npm run build` — tsc passes for both entries; `build/native/` has no buf/connect imports.
4. Regression (grpc path): `cd expo/example && npm install && npx expo run:ios` — existing App.tsx unchanged and working.
5. New path: `cd examples/js/expo/hello-native && make ios` — unary calls work, `helloStream` yields then terminates cleanly on EOF, an error case parses to an `ErrCode`, no `@bufbuild`/`@connectrpc` reachable from the bundle.
