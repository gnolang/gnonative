# hello-native

A minimal Expo app that talks to the Gno core through **`@gnolang/gnonative/native`** — the
buf/connect-free entry point. It demonstrates that the connect-free path needs **no polyfills** and
bundles **no `@bufbuild/*`/`@connectrpc/*`**.

What it does (`App.tsx`): lists key info, reads the remote and chain ID, generates a recovery
phrase, derives its address (base64) and bech32 form, calls `hello`, then consumes `helloStream`
(streaming with no polyfills).

## How it differs from `hello`

- Imports from `@gnolang/gnonative/native` instead of `@gnolang/gnonative`.
- No `import 'react-native-get-random-values'` / text-encoding / streaming polyfills.
- Values are strings: addresses are base64 strings, int64 amounts are strings, and stream fields use
  the proto `json_name` (e.g. `res.Greeting`).

## Prerequisites

Follow the general Build instructions in the repo root
[README](../../../../README.md) (install toolchains via `make asdf.install_tools`). You need the
gomobile toolchain to build the native framework.

This example links the local `@gnolang/gnonative` module at `../../../../expo` (via
`expo.autolinking.nativeModulesDir` and `metro.config.js`) rather than installing it from npm.

## Run

The `Makefile` builds the gomobile framework into the local expo module, builds the module's JS,
installs this app's deps, and launches it:

```console
# iOS
make ios

# Android (physical device: forward metro + gnodev ports first)
make android
make android.reverse_tcp
```
