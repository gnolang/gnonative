<h2 align="center">⚛️ @gnolang/gnonative ⚛️</h2>

### Bring Your Gno.land (d)Apps to React Native Effortlessly!

## Overview

`@gnolang/gnonative` lets React Native / Expo apps talk to the Gno.land blockchain through the Gno
core, built into your app as a native framework (gomobile). The JS client calls the core over a
small in-process bridge and exchanges plain JSON — no gRPC, no protobuf, no polyfills.

```ts
import { GnoNativeProvider, useGnoNativeContext } from '@gnolang/gnonative';
```

The root export and the `./native` subpath expose the same surface (`./native` is kept as an alias).

## Wire dialect

Requests and responses are encoded as JSON in both directions, following the protobuf JSON
(protojson) dialect. The `*Json` type aliases in `src/native/apitypes.ts` are the source of truth for
every request/response shape.

| Value            | Encoding                          |
| ---------------- | --------------------------------- |
| `string`         | `string`                          |
| `bytes`          | base64 `string`                   |
| `int64`/`uint64` | `string`                          |
| `uint32`         | `number`                          |
| field names      | proto `json_name`                 |
| zero values      | omitted                           |

Field names follow the proto `json_name`: usually lowerCamelCase (`gasFee`, `signerAddress`),
sometimes PascalCase (`Msgs`, `Name`, `Greeting`) or snake_case (`key_info`, `tx_json`).
`GnoNativeClient` builds these for you; if you hand-craft a request, use the `*Json` types. Use
`bytesToBase64` / `base64ToBytes` / `base64ToString` (exported from the package) to convert addresses,
keys, and query results.

Errors are `GnoNativeError` (message like `ErrInvalidAddress(#205)`); call `.errCode()` to get the
`ErrCode`.

## Migrating from v4 (the gRPC path) to v5

v5 removes the buf/connect/protobuf gRPC path entirely. If you were on v4:

- Import from `@gnolang/gnonative` (the client is the former `/native` client; `./native` still works).
- Remove `@bufbuild/*` / `@connectrpc/*` and all connect/streaming polyfills
  (`react-native-polyfill-globals`, `web-streams-polyfill`, `text-encoding`, …) from your app.
- Types are the `*Json` aliases: `KeyInfo` → `KeyInfoJson`, `BaseAccount` → `BaseAccountJson`, etc.
- `GRPCError` → `GnoNativeError` (same `errCode()` API).
- Addresses/keys are base64 **strings**, not `Uint8Array`. int64 arguments are **strings**, not
  `bigint`. `bytes` results are base64 strings — decode with `base64ToBytes`/`base64ToString`.
- Fields use the proto `json_name` (e.g. `res.Greeting`, not `res.greeting`).
- `,string` 64-bit fields are strict: pass a JSON string, not a bare number.

See `examples/js/expo/hello` for a complete, minimal example.

# Installation in Expo projects

## Prerequisites

Please follow the general `Build instructions` in the main
[README](https://github.com/gnolang/gnonative/blob/main/README.md) and then:

```console
make asdf.install_tools
```

## Create new Expo app

```
npx create-expo-app my-app --template expo-template-blank-typescript
cd my-app
```

## Add the package to your npm dependencies

```
npm install @gnolang/gnonative
```

## Customize the app

We prepared for you an example Hello World code.

Open App.tsx and replace the content with this:

```tsx
import { GnoNativeProvider, useGnoNativeContext } from '@gnolang/gnonative';
import React, { useEffect, useState } from 'react';
import { StyleSheet, Text, View } from 'react-native';

const config = {
  remote: 'https://gno.berty.io',
  chain_id: 'dev',
};

export default function App() {
  return (
    <GnoNativeProvider config={config}>
      <InnerApp />
    </GnoNativeProvider>
  );
}

const InnerApp = () => {
  const { gnonative } = useGnoNativeContext();
  const [greeting, setGreeting] = useState('');

  useEffect(() => {
    (async () => {
      try {
        const accounts = await gnonative.listKeyInfo();
        console.log(accounts);

        const remote = await gnonative.getRemote();
        const chainId = await gnonative.getChainID();
        console.log('Remote %s ChainId %s', remote, chainId);

        setGreeting(await gnonative.hello('Gno'));

        for await (const res of await gnonative.helloStream('Gno')) {
          // The field is `Greeting` (proto json_name override).
          console.log(res.Greeting);
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return (
    <View style={styles.container}>
      <Text>Gnonative App</Text>
      <Text>{greeting}</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
```

## Run the app

```
# Re-generate the native project directories from scratch
npx expo prebuild --clean
# Run the example app on Android
npx expo run:android
# Run the example app on iOS
npx expo run:ios
```

# Installation in bare React Native projects

For bare React Native projects, you must ensure that you have
[installed and configured the `expo` package](https://docs.expo.dev/bare/installing-expo-modules/)
before continuing.

# Generate new NPM package

You can run one of the following command:

```shell
make npm.pack
```

or

```shell
make npm.publish
```
