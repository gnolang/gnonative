package main

import (
	"context"
	"log"
	"os"

	"github.com/gnolang/gnomobile/service"
)

func main() {
	service, err := service.NewGnomobileService(
		service.WithUseTcpListener(),
		service.WithRemote("http://testnet.gno.berty.io:26657"),
	)
	if err != nil {
		log.Fatalf("failed to run GnomobileService: %v", err)
		os.Exit(1)
	}
	defer service.Close()

	<-context.Background().Done()
}
