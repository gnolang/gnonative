# Architecture

Gno Native Kit builds the Gno core into a mobile app and exposes it to JavaScript over a small
in-process bridge, exchanging plain JSON. There is no gRPC, protobuf, or buf codegen anywhere in the
stack.

## The layers

```
┌──────────────────────────────────────────────────────────────┐
│  React Native / Expo app                                       │
│    @gnolang/gnonative  (expo/)                                 │
│      GnoNativeClient  → GoBridge.invokeMethod / createStream   │
└───────────────────────────────┬────────────────────────────────┘
                                 │ JSON strings over the Expo native module
┌───────────────────────────────▼────────────────────────────────┐
│  Native module  (expo/ios/*.swift, expo/android/*.kt)          │
│    initBridge / invokeMethod / createStream / streamReceive /  │
│    closeStream  → the gomobile-built framework                 │
└───────────────────────────────┬────────────────────────────────┘
                                 │ gomobile bind
┌───────────────────────────────▼────────────────────────────────┐
│  Go framework  (framework/service/)                            │
│    Bridge → ServiceDispatcher                                  │
│      dispatcher.go: name → handler maps (unary + streaming)    │
│      encodes/decodes with encoding/json                        │
└───────────────────────────────┬────────────────────────────────┘
                                 │ direct Go calls
┌───────────────────────────────▼────────────────────────────────┐
│  Service  (service/)                                           │
│    GnoNativeApi: 48 plain-Go methods over gnoclient/keys       │
└───────────────────────────────┬────────────────────────────────┘
                                 │
                          gnoclient / crypto/keys → remote Gno.land node
```

## Wire types

The request/response types live in `api/` as plain Go structs (`api/types.go`). Their `encoding/json`
output reproduces the protobuf JSON (protojson) dialect:

- field names use the proto `json_name` (lowerCamelCase, or the explicit overrides `key_info`,
  `Msgs`, `tx_json`, `TotalFee`, `Name`, `Greeting`);
- `int64`/`uint64` are strings (`,string`);
- `[]byte` is standard base64;
- zero values are omitted (`omitempty`).

`api/errcode.go` holds the `ErrCode` enum and the error helpers (`Wrap`, `Is`, `Code`, …). The
`ErrCode.Error()` format `Name(#code)` is what the JS `GnoNativeError.errCode()` parses.

The TypeScript view of the same contract is `expo/src/native/apitypes.ts` (the `*Json` aliases).

## Keeping the two sides in sync

There is no code generator; the Go structs and the TS aliases are hand-maintained. Two Go test
suites guard against drift:

- `api/types_test.go` — a reflect audit (every field has a json tag + `omitempty`, every 64-bit int
  has `,string`, no non-pointer struct fields) plus golden marshal/unmarshal tests pinning the exact
  wire JSON.
- `framework/service/dispatcher_test.go` — asserts every method of `service.GnoNativeApi` appears in
  exactly one dispatch map with the right signature shape, plus Hello/HelloStream/error smoke tests.

When you add or change a wire type, update `api/types.go` and `expo/src/native/apitypes.ts` together
and let the tests confirm the dialect.
