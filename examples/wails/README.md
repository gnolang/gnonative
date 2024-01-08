# GnoNativeKit on Wails

This is a tutorial on how to use GnoNativeKit with Wails. It is based on the official Wails React-TS template.

[Wails](https://wails.app/) it is a framework for building desktop apps by embedding Go and Web runtimes.

## Prerequisites
- [Wails](https://wails.io/docs/gettingstarted/installation)

## Project Generation


### Generate a new Wails React TS project

To generate a new project, run `wails init -n myproject -t react-ts`. This will generate a new project in the current directory. Please, navigate to the project directory:

```bash
cd myproject
```

Don't worry, you can check out this [link](https://wails.io/docs/gettingstarted/firstproject#project-layout) and become more familiar with this bunch of generated files.


### Configure the GnoNativeKit GRPC Server

We need to change the `app.go` file to start the GnoNative GRPC server:

```go
package main

import (
	"context"
	"fmt"

	framework "github.com/gnolang/gnonative/framework/service"
)

type App struct {
	ctx    context.Context
	bridge *framework.Bridge
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	config := framework.NewBridgeConfig()
	config.UseTcpListener = true
	config.RootDir = "."
 	config.DisableUdsListener = true
	bridge, err := framework.NewBridge(config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	a.bridge = bridge
	fmt.Println("GnoNative GRPC Server created")
}
```

### Install the Frontend Dependencies

Now we have the GRPC server ready to run, we need to install the GnoNative Kit dependencies in the frontend.
Let's also copy the protobufs files we already generated for Typescript.

```bash
cd frontend
npm install @bufbuild/buf @bufbuild/protobuf @bufbuild/protoc-gen-es @connectrpc/connect @connectrpc/connect-web @connectrpc/protoc-gen-connect-es buffer

mkdir -p src/hooks
cp ../../../../templates/es/use-gno-web.ts ./src/hooks/use-gno.ts
cp ../../../../templates/images/logo-universal.png ./src/assets/images 
```
### Set up `@api` alias

The Typescript compiler must be able to resolve import paths starting by `@api` to find the API files.

Copy and paste the following content into a patch file (e.g. `tsconfig.patch`):

```diff
@@ -18,7 +18,11 @@
     "resolveJsonModule": true,
     "isolatedModules": true,
     "noEmit": true,
-    "jsx": "react-jsx"
+    "jsx": "react-jsx",
+    "baseUrl": ".",
+    "paths": {
+      "@api/*": ["../../../../api/gen/es/*"]
+    }
   },
   "include": [
     "src"
```
Copy and paste the following content into an other patch file (e.g. `vite.patch`):

```diff
@@ -1,7 +1,16 @@
 import {defineConfig} from 'vite'
 import react from '@vitejs/plugin-react'
+import path from 'path'
 
 // https://vitejs.dev/config/
 export default defineConfig({
-  plugins: [react()]
+  plugins: [react()],
+  resolve: {
+    alias: [
+      {
+        find: '@api',
+        replacement: path.resolve(__dirname, '../../../../api/gen/es'),
+      },
+    ],
+  },
 })
```

Apply the patches:

```batch
patch tsconfig.json < tsconfig.patch
patch vite.config.ts < vite.patch
```

### Customize the render function

Open `src/App.tsx` and replace the contents with the following code:

```typescript
import { useEffect, useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { useGno } from "./hooks/use-gno";

function App() {
  const [resultText, setResultText] = useState(
    "Start GnoNative Kit on Wails to see result"
  );
  const updateResultText = (result: string) => setResultText(result);

  const gno = useGno();
  useEffect(() => {
    gno.getRemote().then((res) => updateResultText(res));
  }, []);

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo" />
      <div className="input-box"></div>
      <div id="result" className="result">
        {resultText}
      </div>
    </div>
  );
}

export default App;
```

### Run the project

Now we can run the project:

```bash
# go to the project root directory and run:
wails dev
```

### Summary

If you followed all the steps, you should see a native window with the GnoNativeKit logo and the text displayed by the GRPC server.
