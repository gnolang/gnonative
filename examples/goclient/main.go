package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/gnolang/gnomobile/service/rpc"
	"github.com/gnolang/gnomobile/service/rpc/rpcconnect"
)

func main() {
	if len(os.Args) < 2 {
		log.Println(
			`usage: ./goclient <url>
url: http://<ip>:<port>`)
		os.Exit(1)
	}

	client := rpcconnect.NewGnomobileServiceClient(
		http.DefaultClient,
		os.Args[1],
	)

	res, err := client.Render(
		context.Background(),
		connect.NewRequest(&rpc.RenderRequest{
			PackagePath: "gno.land/r/demo/boards",
			Args:        "gnomobile/1",
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.GetResult())
}
