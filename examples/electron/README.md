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

Open the `src/main.ts` file and add the line `nodeIntegration: true` to `webPreferences`.
More details about Electron-Nodejs integration can be found [here](https://www.electronjs.org/docs/latest/tutorial/sandbox#disabling-the-sandbox-for-a-single-process).

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

### Communicating with the Blockchain:

Let's open the `src/preload.ts` file and add change it a little bit.

We'll use the GnoNative Javascript GRPC Client wrapper (aka: `useGno`) to communicate with the GnoNative GRPC Server.

At the top of the `src/preload.ts` file add `import { useGno } from "./api/use-gno"`.

And use the `useGno` function to create a "gno service instance" and become able to call the RPC methods already available.

Check out this sample code bellow. 

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
  gno.getRemote().then((res) => {
    replaceText("gnoRemoteAddress", res);
  });
});
```

In the `index.html` file, add the following line to the ```<body>``` element to display the **Gno Remote Address**:

```html
 GNO Remote Address: <span id="gnoRemoteAddress"></span>.
```

### Run:

In a separate terminal, from the root of the repo, enter the bellow command to start the GnoNative GRPC Server:

```bash
cd gnonative/gnoserver
go run . tcp
```

and in the Electron boilerplate folder, enter:

```bash
npm start
```
