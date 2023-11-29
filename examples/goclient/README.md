# A Go client

This example show how to implement a go client in few lines and make a request
to the Gno blockchain. It calls the render function of the Boards realm.

## Run the server service

Go to the gnomobile root directory.
`cd ../..`

Run the server.
`go run ./goserver`

## Run the client

In another terminal, execute this is the goclient directory (the default port is `26658`):
`go run . http://localhost:26658`

The client prints the raw Boards' messages.
