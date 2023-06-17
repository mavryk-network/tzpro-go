package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func run() error {
	// parse command line flags
	flag.Parse()

	if flag.NArg() == 0 {
		return fmt.Errorf("missing address")
	}

	// use a placeholder calling context
	ctx := context.Background()

	// create a new SDK client
	c, err := tzpro.NewClient("https://api.tzpro.io", nil)
	if err != nil {
		return err
	}

	// parse an address
	addr, err := tezos.ParseAddress(flag.Arg(0))
	if err != nil {
		return err
	}

	// fetch metadata for the address
	md, err := c.GetWalletMetadata(ctx, addr)
	if err != nil {
		// handle 404 NotFound errors in a special way
		if e, ok := tzpro.IsHttpError(err); ok && e.Status == http.StatusNotFound {
			return fmt.Errorf("No metadata for this account")
		}
		return err
	}

	fmt.Printf("Account name: %s\n", md.Alias().Name)
	fmt.Printf("Account kind: %s\n", md.Alias().Kind)

	return nil
}
