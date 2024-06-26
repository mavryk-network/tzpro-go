// Copyright (c) 2020-2024 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"github.com/mavryk-network/mvpro-go/internal/client"
	"github.com/mavryk-network/mvpro-go/internal/util"
)

type ContractAPI interface {
	Get(context.Context, Address, Query) (*Contract, error)
	GetScript(context.Context, Address, Query) (*ContractScript, error)
	GetStorage(context.Context, Address, Query) (*ContractValue, error)
	ListCalls(context.Context, Address, Query) (OpList, error)
	GetConstant(context.Context, ExprHash, Query) (*Constant, error)
	GetBigmap(context.Context, int64, Query) (*Bigmap, error)
	GetBigmapValue(context.Context, int64, string, Query) (*BigmapValue, error)
	ListBigmapValues(context.Context, int64, Query) (BigmapValueList, error)
	ListBigmapKeyUpdates(context.Context, int64, string, Query) (BigmapUpdateList, error)
	ListBigmapUpdates(context.Context, int64, Query) (BigmapUpdateList, error)
	ListTickets(context.Context, Address, Query) (TicketList, error)
	ListTicketBalances(context.Context, Address, Query) (TicketBalanceList, error)
	ListTicketEvents(context.Context, Address, Query) (TicketEventList, error)

	NewQuery() *ContractQuery
	NewEventQuery() *EventQuery
	NewConstantQuery() *ConstantQuery
	NewBigmapQuery() *BigmapQuery
	NewBigmapValueQuery() *BigmapValueQuery
	NewBigmapUpdateQuery() *BigmapUpdateQuery
}

func NewContractAPI(c *client.Client) ContractAPI {
	return &contractClient{client: c}
}

type contractClient struct {
	client *client.Client
}

type Contract struct {
	RowId         uint64               `json:"row_id,omitempty"`
	AccountId     uint64               `json:"account_id,omitempty"`
	Address       Address              `json:"address"`
	CreatorId     uint64               `json:"creator_id,omitempty"`
	Creator       Address              `json:"creator"`
	BakerId       uint64               `json:"baker_id,omitempty"  mvpro:"-"`
	Baker         Address              `json:"baker"               mvpro:"-"`
	FirstSeen     int64                `json:"first_seen"`
	LastSeen      int64                `json:"last_seen"`
	FirstSeenTime time.Time            `json:"first_seen_time"`
	LastSeenTime  time.Time            `json:"last_seen_time"`
	StorageSize   int64                `json:"storage_size"`
	StoragePaid   int64                `json:"storage_paid"`
	TotalFeesUsed float64              `json:"total_fees_used"     mvpro:"-"`
	Script        *Script              `json:"script,omitempty"    mvpro:",hex"`
	Storage       *Prim                `json:"storage,omitempty"   mvpro:",hex"`
	InterfaceHash util.HexBytes        `json:"iface_hash"`
	CodeHash      util.HexBytes        `json:"code_hash"`
	StorageHash   util.HexBytes        `json:"storage_hash"`
	Features      util.StringList      `json:"features"`
	Interfaces    util.StringList      `json:"interfaces"`
	CallStats     map[string]int       `json:"call_stats"          mvpro:"-"`
	NCallsIn      int                  `json:"n_calls_in"          mvpro:"-"`
	NCallsOut     int                  `json:"n_calls_out"         mvpro:"-"`
	NCallsFailed  int                  `json:"n_calls_failed"      mvpro:"-"`
	Bigmaps       map[string]int64     `json:"bigmaps,omitempty"   mvpro:"-"`
	Metadata      map[string]*Metadata `json:"metadata,omitempty"  mvpro:"-"`
}

func (c *Contract) Meta() *Metadata {
	m, ok := c.Metadata[c.Address.String()]
	if !ok {
		m = NewMetadata(c.Address)
		if c.Metadata == nil {
			c.Metadata = make(map[string]*Metadata)
		}
		c.Metadata[c.Address.String()] = m
	}
	return m
}

type ContractList []*Contract

func (l ContractList) Len() int {
	return len(l)
}

func (l ContractList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].RowId
}

type ContractQuery = client.TableQuery[*Contract]

func (a contractClient) NewQuery() *ContractQuery {
	return client.NewTableQuery[*Contract](a.client, "contract")
}

func (c *contractClient) Get(ctx context.Context, addr Address, params Query) (*Contract, error) {
	cc := &Contract{}
	u := params.WithPath(fmt.Sprintf("/explorer/contract/%s", addr)).Url()
	if err := c.client.Get(ctx, u, nil, cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *contractClient) GetScript(ctx context.Context, addr Address, params Query) (*ContractScript, error) {
	cc := &ContractScript{}
	u := params.WithPath(fmt.Sprintf("/explorer/contract/%s/script", addr)).Url()
	if err := c.client.Get(ctx, u, nil, cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *contractClient) GetStorage(ctx context.Context, addr Address, params Query) (*ContractValue, error) {
	cc := &ContractValue{}
	u := params.WithPath(fmt.Sprintf("/explorer/contract/%s/storage", addr)).Url()
	if err := c.client.Get(ctx, u, nil, cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *contractClient) ListCalls(ctx context.Context, addr Address, params Query) (OpList, error) {
	calls := make(OpList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/contract/%s/calls", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &calls); err != nil {
		return nil, err
	}
	return calls, nil
}

func (c *contractClient) ListTickets(ctx context.Context, addr Address, params Query) (TicketList, error) {
	list := make(TicketList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/contract/%s/tickets", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *contractClient) ListTicketBalances(ctx context.Context, addr Address, params Query) (TicketBalanceList, error) {
	list := make(TicketBalanceList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/contract/%s/ticket_balances", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *contractClient) ListTicketEvents(ctx context.Context, addr Address, params Query) (TicketEventList, error) {
	list := make(TicketEventList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/contract/%s/ticket_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
