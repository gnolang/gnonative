# Example app using Expo module

This example app imports the GnoNative Expo module directly from the parent
folder. Every time you want to import the module in the example app (because of
the first time or because you modified the module itself), you have to build it
again.

Follow the following instructions.

## Prerequisites

Please follow the general `Build instructions` in the main
[README](https://github.com/gnolang/gnonative/blob/main/README.md) and then:

```console
make asdf.install_tools
npm config set @buf:registry  https://buf.build/gen/npm/v1/
```

## Building the Expo module

Open a terminal and run:

```
cd .. # cd to the expo folder
make build # or make build.android or make build.ios
npm install
npm run build
```

Open a second terminal and run:

```
cd example
npm install
npx expo prebuild --clean
```

Now the module is compiled, we can build the example app:

- Android

```
npx expo run:android
```

- iOS

```
npx expo run:ios
```
