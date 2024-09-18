# Gno Native Kit

Gno Native Kit is a framework that allows developers to build and port gno.land (d)apps written in the (d)app's native language.

Current Gno (d)apps run on desktop/laptop computers which have Go installed. To run on mobile, the (d)app would need to bundle the Go runtime, which is complicated for most developers.

However, Gno Native Kit helps bypass this complexity by using gRPC to make [calls to the Gno core API](https://buf.build/gnolang/gnonative/docs/main:land.gno.gnonative.v1) and access the blockchain's realm functions on a remote Gno.land node.
These API calls are a programming language-independent wrapper on top of the core supported APIs like [gnoclient](https://github.com/gnolang/gno/tree/master/gno.land/pkg/gnoclient) and [crypto/keys](https://github.com/gnolang/gno/tree/master/tm2/pkg/crypto/keys).

Watch [this Gno Native Kit tutorial](https://www.youtube.com/watch?v=N1HLyQDHGQ0) to easily get started on building and bringing your (d)apps to mobile and desktop.

## Expo module

To use Gno Native Kit, we advice you to use the Expo module in your Expo React-Native project. Please read the [README](expo/README.md) in the `expo` folder.

In the `expo/example` folder, you can find a minimal app using the Expo module.

Other examples are available in the `example/js/expo` folder.

## Bare React-Native project

If you are interested of using Gno Native Kit in a bare React-Native project, please check the `hello` example app in the `example/js/react-native/hello` folder.

## Prerequisites for building Gno Native Kit or example apps

### Install requirements for macOS 13 and macOS 14

(If you are on Ubuntu, see the next section to install requirements.)

Install Xcode. To install the Command Line Developer Tools, in a terminal enter:

```sh
xcode-select --install
```

After the Developer Tools are installed, we need to make sure it is updated. In
System Preferences, click Software Update and update it if needed.

To install asdf using brew, follow instructions at https://asdf-vm.com . In short,
first install brew following the instructions at https://brew.sh . Then, in
a terminal enter:

```sh
brew install asdf gnu-tar gpg
```

If your terminal is zsh, enter:

```sh
echo -e "\n. $(brew --prefix asdf)/libexec/asdf.sh" >> ${ZDOTDIR:-~}/.zshrc
```

If your terminal is bash, enter:

```sh
echo -e "\n. \"$(brew --prefix asdf)/libexec/asdf.sh\"" >> ~/.bash_profile
```

Start a new terminal to get the changes to the environment .

(optional) To install Android Studio, download and install the latest
android-studio-{version}-mac.dmg from https://developer.android.com/studio .
(Tested with Jellyfish 2023.3.1 .)

### Install requirements for Ubuntu 20.04, 22.04 and 24.04

To install asdf, follow instructions at https://asdf-vm.com . In short, in
a terminal enter:

```sh
sudo apt install curl git make
git clone https://github.com/asdf-vm/asdf.git ~/.asdf
echo '. "$HOME/.asdf/asdf.sh"' >> ~/.bashrc
echo 'export ANDROID_HOME="$HOME/Android/Sdk"' >> ~/.bashrc
echo 'export ANDROID_NDK_HOME="$ANDROID_HOME/ndk/23.1.7779620"' >> ~/.bashrc
```

Start a new terminal to get the changes to the environment .

To install Android Studio, download the latest
android-studio-{version}-linux.tar.gz from
https://developer.android.com/studio . (Tested with Jellyfish 2023.3.1 .)
In a terminal, enter the following with the correct {version}:

```sh
sudo tar -C /usr/local -xzf android-studio-{version}-linux.tar.gz
```

To launch Android Studio, in a terminal enter:

```sh
/usr/local/android-studio/bin/studio.sh &
```

### Install the tools with asdf (only need to do once)

```sh
make asdf.install_tools
```

If you get an error like "https://github.com/CocoaPods/CLAide.git (at master@97b765e) is not yet checked out" then reinstall cocoapods like this:

```sh
asdf uninstall cocoapods
make asdf.install_tools
```

### Build for Android

#### Set up the Android NDK

-   Launch Android Studio and accept the default startup options. Create a new
    default project (so that we get the main screen).
-   On the Tools menu, open the SDK Manager.
-   In the "SDK Tools" tab, click "Show Package Details". Expand
    "NDK (Side by side)" and check "23.1.7779620".
-   Click OK to install and close the SDK Manager.
