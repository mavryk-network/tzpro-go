package main

import (
	"context"

	"github.com/mavryk-network/mvpro-go/mvpro"
)

func TestNft(ctx context.Context, c *mvpro.Client) {
	p := mvpro.NewQuery()
	// dex
	try("ListNftMarkets", func() {
		if _, err := c.Nft.ListMarkets(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListNftEvents", func() {
		if _, err := c.Nft.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})
	addr := mvpro.NewAddress("KT1HbQepzV1nVGg8QVznG7z4RcHseD5kwqBn")
	try("ListMarketEvents", func() {
		if _, err := c.Nft.ListMarketEvents(ctx, addr, p); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListNftPositions", func() {
		if _, err := c.Nft.ListPositions(ctx, p); err != nil {
			panic(err)
		}
	})
	try("ListNftmarketPositions", func() {
		if _, err := c.Nft.ListMarketPositions(ctx, addr, p); err != nil {
			panic(err)
		}
	})

	// trades
	try("ListNftTrades", func() {
		if _, err := c.Nft.ListTrades(ctx, p); err != nil {
			panic(err)
		}
	})
	try("ListNftMarketTrades", func() {
		if _, err := c.Nft.ListMarketTrades(ctx, addr, p); err != nil {
			panic(err)
		}
	})
}
