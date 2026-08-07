<h2 align="center">⚛️ @gnolang/gnonative ⚛️</h2>

### Bring Your Gno.land (d)Apps to React Native Effortlessly!

## Overview

`@gnolang/gnonative` simplifies the process of access the Gno.land (d)apps to mobile by using gRPC to connect with core blockchain functions.

It helps bypass this complexity by using gRPC to make [calls to the Gno core API](https://buf.build/gnolang/gnonative/docs/main:land.gno.gnonative.v1) and access the blockchain's realm functions on a remote Gno.land node.

## Two entry points

There are two entry points; pick one.

### `@gnolang/gnonative` (default, gRPC path)

The original API. Uses `@connectrpc/connect(-web)` + `@bufbuild/protobuf` for
serialization/dispatch over an in-process connect server. Requests/responses are protobuf message
objects (`Uint8Array` for bytes, `bigint` for int64).

```ts
import { GnoNativeProvider, useGnoNativeContext } from '@gnolang/gnonative';
```

### `@gnolang/gnonative/native` (buf/connect-free path)

A drop-in-shaped client that **never touches buf/connect**. It talks to the same Go core through a
direct dispatcher (`GoBridge.invokeMethod` / `createStream`), so an app using only `/native` bundles
no `@bufbuild/*`/`@connectrpc/*` and needs **no streaming polyfills**. Plain JSON on the wire; values
are strings.

```ts
import { GnoNativeProvider, useGnoNativeContext, GnoNativeClient } from '@gnolang/gnonative/native';
```

The API is method-for-method parallel to the gRPC path, so migrating is mostly an import swap plus
adapting argument types (see the type mapping below). See `examples/js/expo/hello-native` for a full
example.

> Note: `@bufbuild/*`/`@connectrpc/*` remain in `dependencies` so the default path keeps working.
> Metro simply never bundles them for a `/native`-only app.

#### JSON type mapping (`/native`)

The `/native` path uses **protobuf JSON (protojson) in both directions**. The generated types live
in `src/native/types.gen.ts` (the `*Json` aliases). The encoding differs from the message objects on
the gRPC path:

| Proto type       | gRPC path (message object) | `/native` path (`*Json`) |
| ---------------- | -------------------------- | ------------------------ |
| `string`         | `string`                   | `string`                 |
| `bytes`          | `Uint8Array`               | base64 `string`          |
| `int64`/`sint64` | `bigint`                   | `string`                 |
| field names      | camelCase                  | proto `json_name`        |
| default values   | present                    | omitted                  |

Field names follow the proto `json_name`: usually camelCase (`gasFee`, `signerAddress`), sometimes
PascalCase (`Msgs`, `Name`, `Greeting`) or snake_case (`key_info`, `tx_json`). `GnoNativeClient`
builds these for you; if you hand-craft a request, use the `*Json` types — they are the source of
truth. Never decode a `/native` payload with a gRPC-path decoder or vice versa. Use `bytesToBase64` /
`base64ToBytes` (exported from `/native`) to convert addresses and keys.

# API documentation

The RPC API documentation is available in the Buf registry:

- [Documentation](https://buf.build/gnolang/gnonative/docs/main:land.gno.gnonative.v1)

# Installation in Expo projects

## Prerequisites

Please follow the general `Build instructions` in the main
[README](https://github.com/gnolang/gnonative/blob/main/README.md) and then:

```console
make asdf.install_tools
npm config set @buf:registry  https://buf.build/gen/npm/v1/
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
          console.log(res.greeting);
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
