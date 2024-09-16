This is a new [**React Native**](https://reactnative.dev) project (without Expo), bootstrapped using [`@react-native-community/cli`](https://github.com/react-native-community/cli).
Then we tweaked the project to add gnonative capabilities. We will give you below the step to reproduce that.

# Getting Started

> **Note**: Make sure you have completed the [React Native - Environment Setup](https://reactnative.dev/docs/environment-setup) instructions till "Creating a new application" step, before proceeding.

## Step 1: Install the dependencies

```bash
npm install
```

## Step 2: Start the Metro Server

First, you will need to start **Metro**, the JavaScript _bundler_ that ships _with_ React Native.

To start Metro, run the following command from the _root_ of your React Native project:

```bash
# using npm
npm start
```

## Step 2: Start your Application

Let Metro Bundler run in its _own_ terminal. Open a _new_ terminal from the _root_ of your React Native project. Run the following command to start your _Android_ or _iOS_ app:

### For Android

```bash
npm run android
```

### For iOS

If you have not install the pod dependencies, you can do it by running:

```bash
cd ios && pod install && cd ..
```

Then you can run the iOS app with the following command:

```bash
npm run ios
```

If everything is set up _correctly_, you should see your new app running in your _Android Emulator_ or _iOS Simulator_ shortly provided you have set up your emulator/simulator correctly.

This is one way to run your app — you can also run it directly from within Android Studio and Xcode respectively.

## Recreate a new React-Native project without Expo

### Create a new React Native project

```bash
npx @react-native-community/cli@latest init test --pm npm
```

### Edit JavaScript files

```bash
cp hello/App.tsx test/App.tsx
cp -r hello/GoBridge test/
cp -r hello/api test/
cp -r hello/grpc test/
cp -r hello/provider test/
```

Edit `package.json`:

```diff
@@ -11,7 +11,18 @@
   },
   "dependencies": {
     "react": "18.3.1",
-    "react-native": "0.75.3"
+    "react-native": "0.75.3",
+    "@buf/gnolang_gnonative.bufbuild_es": "^1.7.2-20240905152811-1254a97aecd9.2",
+    "@buf/gnolang_gnonative.connectrpc_es": "^1.4.0-20240905152811-1254a97aecd9.3",
+    "@bufbuild/protobuf": "^1.7.2",
+    "@connectrpc/connect": "^1.4.0",
+    "@connectrpc/connect-web": "^1.4.0",
+    "base-64": "^1.0.0",
+    "react-native-fetch-api": "^3.0.0",
+    "react-native-polyfill-globals": "^3.1.0",
+    "react-native-url-polyfill": "^2.0.0",
+    "text-encoding": "^0.7.0",
+    "web-streams-polyfill": "^3.2.1"
   },
   "devDependencies": {
     "@babel/core": "^7.20.0",
```

### Modify Android project

Go to the `test` directory:

```bash
cd test
```

Copy files from the `hello/android` project:

```bash
cp -r ../hello/android/app/src/main/java/land android/app/src/main/java/
```

Edit `android/app/src/main/java/com/test/MainApplication.kt`

```patch
@@ -10,6 +10,8 @@ import com.facebook.react.defaults.DefaultNewArchitectureEntryPoint.load
 import com.facebook.react.defaults.DefaultReactHost.getDefaultReactHost
 import com.facebook.react.defaults.DefaultReactNativeHost
 import com.facebook.soloader.SoLoader
+import land.gno.gobridge.GoBridgePackage
+import land.gno.rootdir.RootDirPackage

 class MainApplication : Application(), ReactApplication {

@@ -19,6 +21,8 @@ class MainApplication : Application(), ReactApplication {
             PackageList(this).packages.apply {
               // Packages that cannot be autolinked yet can be added manually here, for example:
               // add(MyReactNativePackage())
+              add(RootDirPackage())
+              add(GoBridgePackage())
             }

         override fun getJSMainModuleName(): String = "index"
```

Edit `android/app/build.gradle`:

```diff
@@ -2,6 +2,8 @@ apply plugin: "com.android.application"
 apply plugin: "org.jetbrains.kotlin.android"
 apply plugin: "com.facebook.react"

+def frameworkDir = "${rootDir.getAbsoluteFile().getParentFile().getParentFile().getParentFile().getParentFile().getParentFile().getAbsolutePath()}/framework"
+
 /**
  * This is the configuration block to customize your React Native Android app.
  * By default you don't need to apply any configuration, just uncomment the lines you need.
@@ -107,7 +109,37 @@ android {
     }
 }

+// Auto-build gomobile.aar by running Makefile rule
+task makeDeps(description: 'Build gnocore.aar (Gno go core)') {
+    outputs.files fileTree(dir: "${frameworkDir}/android", include: ["*.jar", "*.aar"])
+
+    doLast {
+        if (System.properties['os.name'].toLowerCase().contains('windows')) {
+            logger.warn("Warning: can't run make on Windows, you must build gomobile.aar manually")
+            return
+        }
+
+        def checkMakeInPath = exec {
+            standardOutput = new ByteArrayOutputStream() // equivalent to '> /dev/null'
+            ignoreExitValue = true
+            commandLine 'bash', '-l', '-c', 'command -v make'
+        }
+
+        if (checkMakeInPath.getExitValue() == 0) {
+            exec {
+                def makefileDir = "${rootDir.getPath()}/../../../../.."
+                workingDir makefileDir
+                environment 'PWD', makefileDir
+                commandLine 'make', 'build.android'
+            }
+        } else {
+            logger.warn('Warning: make command not found in PATH')
+        }
+    }
+}
+
 dependencies {
+    implementation fileTree(dir: "${frameworkDir}/android", include: ["*.jar"])
     // The version of react-native is set by the React Native Gradle Plugin
     implementation("com.facebook.react:react-android")

@@ -116,4 +148,6 @@ dependencies {
     } else {
         implementation jscFlavor
     }
+
+    implementation makeDeps.outputs.files
 }
```

### Modify iOS project

Go to the `test` directory if it's not the case:

```bash
cd test
```

Copy files from the `hello/ios` project:

```bash
cp -r ../hello/ios/Sources ./ios
cp ../hello/ios/hello/hello-Bridging-Header.h ./ios/test/test-Bridging-Header.h
```

#### Add files to the Xcode project

Open Xcode and open the project `ios/test.xcworkspace`.

In Xcode, in the tree view, click on the `test` folder under the `test` project. Then click on `File` -> `Add Files to "test"...` then add the `ios/Sources` folder. Add file again to add `ios/test/test-Bridging-Header.h`.

#### Link binary with libraries

In Xcode, double-click on the `test` project then open the `Build Phases` tab. Expand the `Link Binary With Libraries` section, click on the `+` and add the following libraries:

-   `libresolv.tbd`
-   click on `Add Other...` and `Add Files...` and add the `../../../../framework/ios/GnoCore.xcframework/` folder

#### Configure the bridging header

Doucle-click on the `test` project and open the `Build Settings` tab. Search for `Objective-C Bridging Header` and set the value to `test/test-Bridging-Header.h`.

### Congratulations! :tada:

You've successfully run and modified your React Native App. :partying_face:

# Troubleshooting

If you can't get this to work, see the [Troubleshooting](https://reactnative.dev/docs/troubleshooting) page.
