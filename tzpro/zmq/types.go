// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package zmq

import (
	"github.com/mavryk-network/tzgo/tezos"
	"github.com/mavryk-network/tzpro-go/tzpro/index"
)

type (
	OpHash    = tezos.OpHash
	BlockHash = tezos.BlockHash
	Op        = index.Op
	Block     = index.Block
	Status    = index.Status
)

var (
	ParseOpHash    = tezos.ParseOpHash
	ParseBlockHash = tezos.ParseBlockHash
)
