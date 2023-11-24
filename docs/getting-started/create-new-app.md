# Create a new App

## Overview

This guide will walk you through creating a new Gno Mobile App using the `create-app` command.

## Prerequisites

Check the [Build instruction](../../README.md#build-instructions) guide for prerequisites.

## Create a new App

To create a new app, run the following command:

```console
cd gnomobile (root of the repo)
make create-app APP_NAME=MyApp
```

This will create a new app in the `examples/MyApp` directory containing a basic integration with Gno.

## Run the App

To run the app, run the following command:

```console
cd examples/MyApp
yarn start
```
and then run the following command in another terminal:

```console
cd examples/MyApp
npx react-native [run-android|run-ios]
```

## Using Gno in your App

To use Gno in your app, you can import the `useGno` hook from `@gno/hooks/use-gno`:

```ts
import { useGno } from '@gno/hooks/use-gno';

export default function App() {
  const gno = useGno();

  React.useEffect(() => {
    gno.getRemote()
    .then(res => console.log(res))
    .catch(err => console.log(err));
  }, []);

...
