# gnomobile
*experimental* Gno &amp; mobile

## Build instructions

### Get a copy of the repo

```console
git clone git@github.com:gnolang/gnomobile.git
cd gnomobile
```

### Android

#### Build the Go code as a librairy

```console
make build.android
```

#### Start metro

```console
cd gnoboard
yarn start
```

#### Connect your Android phone/emulator and bind its port to metro

You can either connect an Android phone via USB cable, or launch an emulator device from Android Studio.

##### Real device

Connect your device and bind the port to metro:

```console
make android.reverse_tcp
```

##### Emulator device

You can either run Android Studio and open the Android project in `gnoboard/android`.
If you prefer the CLI option:

```console
android-studio ./android
```
Once done, bind the port to metro:

```console
make android.reverse_tcp
```

#### Build with android-studio

Open Android Studio and open the current Android project if it's not already done.
Select the right device in the device list. Open the `Run` menu, and select `Run app`.
See more: https://developer.android.com/studio/run#basic-build-run

### iOS

#### Build the Go code as a librairy

```console
make build.ios
```

#### Start metro

```console
cd gnoboard
yarn start
```

#### Open Xcode and connect your iOS device

Open Xcode and open the GnoBoard Xcode workspace: `gnoboard/ios/gnoboard.xcworkspace`
You can either connect an iOS phone via USB cable, or launch an emulator device from Xcode.
See more: https://developer.apple.com/documentation/xcode/running-your-app-in-simulator-or-on-a-device

#### Select a developper certificate

In Xcode, double click on gnoboard project on the left pane, go to the `Signing & Capabilities` pane.
In the `Signing` section, select your `team` certificate.

#### Build with Xcode

Select the right device in the device list. Open the `Product` menu, and select `Run`.
See more: https://developer.apple.com/documentation/xcode/building-and-running-an-app
