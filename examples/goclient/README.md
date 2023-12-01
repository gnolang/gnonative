# A Go client

This example show how to implement a go client in few lines and make a request
to the Gno blockchain. It calls the render function of the Boards realm.

## Run the server service

Please read the corresponding
[README](https://github.com/gnolang/gnomobile/blob/main/goserver/README.md).

## Run the client

In another terminal, execute the client.

### Unix Domain Socket

The default path is `/tmp/socket`.

`go run . uds` or specify the absolute socket path:
`go run . uds -path /tmp/socket`

### TCP connection

The default port is `26658`.

`go run . tcp` or specify the remote address:
`go run . tcp -addr http://localhost:26658`

The client prints the raw Boards' messages.
