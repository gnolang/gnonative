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
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
	"github.com/gnolang/gnonative/v4/api/gen/go/_goconnect"
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
			ShortHelp:   "start a Go client of GnoNative",
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
	path := fs.String("path", "/tmp/gnonative.sock", "absolute path of the socket")

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

			client := _goconnect.NewGnoNativeServiceClient(
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
			client := _goconnect.NewGnoNativeServiceClient(
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

func exampleAction(client _goconnect.GnoNativeServiceClient) error {
	res, err := client.Render(
		context.Background(),
		connect.NewRequest(&api_gen.RenderRequest{
			PackagePath: "gno.land/r/gnoland/pages",
			Args:        "p/partners",
		}),
	)
	if err != nil {
		return err
	}
	log.Println(res.Msg.GetResult())
	return nil
}
