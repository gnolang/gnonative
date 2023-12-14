# Go server

This Go executable run a GnoNative service.

## Unix Domain Socket connection

The default path is `/tmp/socket`.

`go run . uds`

`go run . uds -path /tmp/gnonative.sock`

## TCP connection

The default port is `26658`.

`go run . tcp`

`go run . tcp -addr localhost:26658`

The gRPC server prints the TCP address/socket path it listens to:

`server UDS path: xxx` or `server TCP address: xxx`

To close it, press Ctrl+C in the terminal.

## Usage

`go run .`
