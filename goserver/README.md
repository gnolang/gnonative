# Go server

This Go executable run a Gnomobile service. The gRPC server prints the port it
listens to:
`2023-11-29T13:59:22.566+0100    INFO    service/service.go:231  createTcpGrpcServer: gRPC server listens to    {"port": 58748}`

## Run the service

`go run ./goserver`

To close it, press Ctrl+C in the terminal.
