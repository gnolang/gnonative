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

	framework "github.com/gnolang/gnomobile/framework/service"
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
	config.RootDir = ""
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
npm install @bufbuild/buf @bufbuild/protobuf @bufbuild/protoc-gen-es @connectrpc/connect @connectrpc/connect-web @connectrpc/protoc-gen-connect-es

mkdir -p src/api
cp -r ../../../../api/gen/es/* ./src/api/
cp ../../../../templates/es/use-gno-web.ts ./src/api/use-gno.ts
cp ../../../../templates/images/logo-universal.png ./src/assets/images 
```

Open `src/App.tsx` and replace the contents with the following code:

```typescript
import { useEffect, useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { useGno } from "./api/use-gno";

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