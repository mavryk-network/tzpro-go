// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/tzpro/token"
	"github.com/mavryk-network/tzgo/tezos"
)

type (
	Query = client.Query

	OpHash  = tezos.OpHash
	Address = tezos.Address
	Z       = tezos.Z

	Token        = token.Token
	TokenAddress = tezos.Token
)

var (
	AddressTypeContract = tezos.AddressTypeContract
	ParseAddress        = tezos.ParseAddress
	NewAddress          = tezos.NewAddress
)
