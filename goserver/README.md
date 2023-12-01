# Go server

This Go executable run a Gnomobile service.

## Unix Domain Socket connection

The default path is `/tmp/socket`.

`go run ./goserver uds`

`go run ./goserver uds -path /tmp/gnomobile.sock`

## TCP connection

The default port is `26658`.

`go run ./goserver tcp`

`go run ./goserver tcp -addr http://localhost:26658`

The gRPC server prints the TCP address/socket path it listens to:

`server UDS path: xxx` or `server TCP address: xxx`

To close it, press Ctrl+C in the terminal.

## Usage

`go run ./goserver`
