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
make android.reverse_tcp # for Android, after you have connected an Android device (simulator or real device)
yarn start
```

and then run the following command in another terminal in the repo folder:

```console
cd $(go list -m -f '{{.Dir}}') # go to the root of the repo
cd examples/js/react-native/MyApp
npx react-native [run-android|run-ios]
```

## Using Gno in your App

To use Gno in your app, you can import the `useGno` hook from
`@gno/hooks/use-gno`:

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
```

## Regenerate dependencies

If you changed some Go code, or updated the React-Native dependencies, you have to build them again:
```console
cd $(go list -m -f '{{.Dir}}') # go to the root of the repo
APP_NAME=MyApp make build.ios # for iOS
APP_NAME=MyApp make build.android # for Android
```
