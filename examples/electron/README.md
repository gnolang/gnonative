# Installing GNO and Electron

To make our setup easier, let's use the Electron TypeScript boilerplate. It's a simple Electron app that uses TypeScript and TSC compiler.

### Get the boilerplate from GitHub:

```bash
cd gnonative/examples

curl -LO https://github.com/electron/electron-quick-start-typescript/archive/refs/heads/main.zip
```

### Unzip the file:

```bash
unzip main.zip
```

### Go to the directory and install the dependencies:

```bash
cd electron-quick-start-typescript-main

npm install @connectrpc/connect @connectrpc/connect-node @connectrpc/protoc-gen-connect-es @bufbuild/protobuf @bufbuild/buf @bufbuild/protoc-gen-es
```

### Enable the Electron-Nodejs integration:

Open the `main.ts` file and add the following line as shown [in docs](https://www.electronjs.org/docs/latest/tutorial/sandbox#disabling-the-sandbox-for-a-single-process):

```bash
webPreferences: {
    preload: path.join(__dirname, "preload.js"),
    nodeIntegration: true,
},
```

### Update the `package.json` file:

In order to allow the TSC compiler to compile the JS files, we need to add the `--AllowJs` flag to the build script:

```bash
"scripts": {
    "build": "tsc --AllowJs",
    ...
```

### Copy the GNO files to the boilerplate directory:

Now we need to copy the GNO files to the boilerplate directory. We will copy the `gnonative/api/gen/es` directory and the `templates/es/use-gno.ts` file to the `src/api` directory.

```bash
cp -r ../../api/gen/es ./src/api
cp ../../templates/es/use-gno.ts ./src/api/use-gno.ts
```

### Use the GNO API in the boilerplate:

Now we can use the GNO API in the boilerplate. Let's open the `preload.ts` file and add the following lines:

```bash
import { useGno } from "./api/use-gno";


window.addEventListener("DOMContentLoaded", () => {

  const replaceText = (selector: string, text: string) => {
    const element = document.getElementById(selector);
    if (element) {
      element.innerText = text;
    }
  };

  const gno = useGno();
  gno.getChainID().then((res) => {
    replaceText("chainID", res);
  });
});
```

In the `index.html` file, we need to add the following line to display the chainID:

```html
 chainID: <span id="chainID"></span>.
```

### Run:

Now we can start the GnoNative GRPC API:

```bash
cd gnonative/gnoserver
go run . tcp
```

and the boilerplate:

```bash
npm start
```
