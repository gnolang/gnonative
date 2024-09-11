# Create a New App

## Overview

This guide will walk you through creating a new Gno Native Kit App using the
`new-app` command.

## Prerequisites

Check the [Build instructions](../../README.md#build-instructions) guide for
prerequisites.

## Create a new App

To create a new app, run the following command in the repo folder:

```console
cd $(go list -m -f '{{.Dir}}') # go to the root of the repo
APP_NAME=MyApp make new-app
```

This will create a new app in the `examples/js/react-native/MyApp` directory containing a basic
integration with Gno.

## Run the App

To run the app, run the following command:

```console
cd examples/js/react-native/MyApp
npx expo [run:android|run-ios]
```

## Using Gno in your App

To use Gno in your app, you can import the `useGno` hook from
`@gno/hooks/use-gno`:

```ts
import { GnoNativeProvider, useGnoNativeContext } from '@gnolang/gnonative';

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

  useEffect(() => {
    (async () => {
      try {
        const remote = await gnonative.getRemote();
        const chainId = await gnonative.getChainID();
        console.log('Remote %s ChainId %s', remote, chainId);
        ...
}
```
