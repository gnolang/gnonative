package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/gnolang/gnonative/v4/service"
	"github.com/peterbourgon/ff/v3/ffcli"
)

// path of the remote Gno node
var remote string

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
		fs = flag.NewFlagSet("goserver", flag.ContinueOnError)
	}

	fs.StringVar(&remote, "remote", "https://api.gno.berty.io:443", "address of the Gno node")

	var root *ffcli.Command
	{
		root = &ffcli.Command{
			ShortUsage:  "goserver [flag] <subcommand>",
			ShortHelp:   "start a Go server of GnoNative",
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
	fs := flag.NewFlagSet("goserver uds", flag.ExitOnError)
	path := fs.String("path", "/tmp/gnonative.sock", "path of the socket to listen to")

	return &ffcli.Command{
		Name:       "uds",
		ShortUsage: "goserver uds [flags]",
		ShortHelp:  "Listen on Unix Domain Socket",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			options := []service.GnoNativeOption{
				service.WithRemote(remote),
			}

			if *path != "" {
				options = append(options, service.WithUdsPath(*path))
			}

			service, err := service.NewGnoNativeService(options...)
			if err != nil {
				log.Fatalf("failed to run GnoNativeService: %v", err)
				os.Exit(1)
			}
			defer service.Close()

			fmt.Printf("server UDS path: %s\n", service.GetUDSPath())

			// <-context.Background().Done()
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			<-c
			return nil
		},
	}
}

func tcp() *ffcli.Command {
	fs := flag.NewFlagSet("goserver tcp", flag.ExitOnError)
	addr := fs.String("addr", "", "TCP address to listen to")

	return &ffcli.Command{
		Name:       "tcp",
		ShortUsage: "goserver tcp [flags]",
		ShortHelp:  "Listen on TCP",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			options := []service.GnoNativeOption{
				service.WithRemote(remote),
				service.WithUseTcpListener(),
			}

			if *addr != "" {
				options = append(options, service.WithTcpAddr(*addr))
			}

			service, err := service.NewGnoNativeService(options...)
			if err != nil {
				log.Fatalf("failed to run GnoNativeService: %v", err)
				os.Exit(1)
			}
			defer service.Close()

			fmt.Printf("server TCP address: %s\n", service.GetTcpAddr())

			<-context.Background().Done()
			return nil
		},
	}
}
