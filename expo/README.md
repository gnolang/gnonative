<h2 align="center">⚛️ @gnolang/gnonative ⚛️</h2>

### Bring Your Gno.land (d)Apps to React Native Effortlessly!

## Overview

`@gnolang/gnonative` simplifies the process of access the Gno.land (d)apps to mobile by using gRPC to connect with core blockchain functions.

It helps bypass this complexity by using gRPC to make [calls to the Gno core API](https://buf.build/gnolang/gnonative/docs/main:land.gno.gnonative.v1) and access the blockchain's realm functions on a remote Gno.land node.

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
