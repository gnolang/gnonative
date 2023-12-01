package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/gnolang/gnomobile/service/rpc"
	"github.com/gnolang/gnomobile/service/rpc/rpcconnect"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/peterbourgon/unixtransport"
)

func main() {
	err := runMain(os.Args[1:])

	switch {
	case err == nil:
		// noop
	case err == flag.ErrHelp || strings.Contains(err.Error(), flag.ErrHelp.Error()):
		os.Exit(2)
	default:
		fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		os.Exit(1)
	}
}

func runMain(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup flags
	var fs *flag.FlagSet
	{
		fs = flag.NewFlagSet("goclient", flag.ContinueOnError)
	}

	var root *ffcli.Command
	{
		root = &ffcli.Command{
			ShortUsage:  "goclient <subcommand>",
			ShortHelp:   "start a Go client of Gnomobile",
			FlagSet:     fs,
			Subcommands: []*ffcli.Command{uds(), tcp()},
			Exec: func(ctx context.Context, args []string) error {
				return flag.ErrHelp
			},
		}
	}

	if err := root.ParseAndRun(ctx, args); err != nil {
		log.Fatal(err)
	}

	return nil
}

func uds() *ffcli.Command {
	fs := flag.NewFlagSet("goclient uds", flag.ExitOnError)
	path := fs.String("path", "/tmp/gnomobile.sock", "absolute path of the socket")

	return &ffcli.Command{
		Name:       "uds",
		ShortUsage: "goclient uds [flags]",
		ShortHelp:  "Connect via Unix Domain Socket",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			// custom transport with deadline of 2 seconds
			t := &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(network, addr, time.Second*2)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(time.Now().Add(time.Second * 2))
					return conn, nil
				},
			}

			// Register the "http+unix" and "https+unix" protocols in the transport.
			unixtransport.Register(t)

			httpClient := &http.Client{Transport: t}

			// add a trailing colon
			fullPath := fmt.Sprintf("http+unix://%s:", *path)

			client := rpcconnect.NewGnomobileServiceClient(
				httpClient,
				fullPath,
			)
			if err := exampleAction(client); err != nil {
				log.Fatal(err)
				return err
			}

			return nil
		},
	}
}

func tcp() *ffcli.Command {
	fs := flag.NewFlagSet("goclient tcp", flag.ExitOnError)
	addr := fs.String("addr", "http://localhost:26658", "remote TCP address")

	return &ffcli.Command{
		Name:       "tcp",
		ShortUsage: "goclient tcp [flags]",
		ShortHelp:  "Connect via TCP",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			client := rpcconnect.NewGnomobileServiceClient(
				http.DefaultClient,
				*addr,
			)
			if err := exampleAction(client); err != nil {
				log.Fatal(err)
				return err
			}

			return nil
		},
	}
}

func exampleAction(client rpcconnect.GnomobileServiceClient) error {
	res, err := client.Render(
		context.Background(),
		connect.NewRequest(&rpc.RenderRequest{
			PackagePath: "gno.land/r/demo/boards",
			Args:        "gnomobile/1",
		}),
	)
	if err != nil {
		return err
	}
	log.Println(res.Msg.GetResult())
	return nil
}
