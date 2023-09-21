package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gno/tm2/pkg/amino/genproto"
	"github.com/gnolang/gno/tm2/pkg/commands"

	// TODO: move these out.
	"github.com/gnolang/gnomobile/service/gnomobiletypes"
)

func main() {
	cmd := commands.NewCommand(
		commands.Metadata{
			LongHelp: "Generates proto bindings for Amino packages",
		},
		commands.NewEmptyConfig(),
		execGen,
	)

	if err := cmd.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)

		os.Exit(1)
	}
}

func execGen(_ context.Context, _ []string) error {
	pkgs := []*amino.Package{
		gnomobiletypes.Package,
	}

	for _, pkg := range pkgs {
		genproto.WriteProto3Schema(pkg)
	}

	return nil
}
