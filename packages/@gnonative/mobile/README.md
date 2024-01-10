### Development

You can use a local registry to test the package.

```bash
docker pull verdaccio/verdaccio:nightly-master
docker run -it --rm --name verdaccio -p 4873:4873 verdaccio/verdaccio

npm adduser --registry http://0.0.0.0:4873/
npm publish --registry http://0.0.0.0:4873/
```

To test the package in a project, you can use:

```bash
npm install @gnonative/hooks --registry http://localhost:4873
```