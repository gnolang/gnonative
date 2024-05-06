# gnonative

Develop for Gno using your app's native language

# API documentation

We created a Javascript object helper (useGno) that wraps the RPC API. The
useGno object is defined in the GnoNative repo at `/expo/src/hooks/use-gno.ts`.

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
npx create-expo-app my-app
cd my-app
```

## Add the package to your npm dependencies

```
npm install @gnolang/gnonative
```

## Customize the app

We prepared for you an example Hello World code.

Open App.js and replace the content with this:

```tsx
import React, { useEffect, useState } from 'react';
import { StyleSheet, Text, View } from 'react-native';

import * as Gnonative from '@gnolang/gnonative';

export default function App() {
  const gno = Gnonative.useGno();
  const [greeting, setGreeting] = useState('');

  useEffect(() => {
    const greeting = async () => {
      try {
        // sync Hello function
        setGreeting(await gno.hello('Gno'));

        // async Hello function
        for await (const res of await gno.helloStream('Gno')) {
          console.log(res.greeting);
        }
      } catch (error) {
        console.log(error);
      }
    };
    greeting();
  }, []);

  return (
    <View style={styles.container}>
      <Text>Hey {greeting}</Text>
    </View>
  );
}

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
make pack
```

or

```shell
make publish
```
