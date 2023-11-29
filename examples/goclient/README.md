# A Go client

This example show how to implement a go client in few lines and make a request
to the Gno blockchain. It calls the render function of the Boards realm.

## Run the server service

`go run ./../../goserver`

The server print the port it listen to. For instance, you should have something
like this :
`2023-11-29T13:59:22.566+0100    INFO    service/service.go:231  createTcpGrpcServer: gRPC server listens to    {"port": 58748}`
Note the port is 58748.

## Run the client

In another terminal, execute this (note that we are using the port `58748`):
`go run ./client http://localhost:58748`

The client prints the raw Boards' messages.
