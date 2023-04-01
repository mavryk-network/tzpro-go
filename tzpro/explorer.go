// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Tip struct {
	Name               string             `json:"name"`
	Network            string             `json:"network"`
	Symbol             string             `json:"symbol"`
	ChainId            tezos.ChainIdHash  `json:"chain_id"`
	GenesisTime        time.Time          `json:"genesis_time"`
	Hash               tezos.BlockHash    `json:"block_hash"`
	Height             int64              `json:"height"`
	Cycle              int64              `json:"cycle"`
	Timestamp          time.Time          `json:"timestamp"`
	Protocol           tezos.ProtocolHash `json:"protocol"`
	TotalAccounts      int64              `json:"total_accounts"`
	TotalContracts     int64              `json:"total_contracts"`
	TotalRollups       int64              `json:"total_rollups"`
	FundedAccounts     int64              `json:"funded_accounts"`
	DustAccounts       int64              `json:"dust_accounts"`
	DustDelegators     int64              `json:"dust_delegators"`
	TotalOps           int64              `json:"total_ops"`
	Delegators         int64              `json:"delegators"`
	Bakers             int64              `json:"bakers"`
	Rolls              int64              `json:"rolls"`
	RollOwners         int64              `json:"roll_owners"`
	NewAccounts30d     int64              `json:"new_accounts_30d"`
	ClearedAccounts30d int64              `json:"cleared_accounts_30d"`
	FundedAccounts30d  int64              `json:"funded_accounts_30d"`
	Inflation1Y        float64            `json:"inflation_1y"`
	InflationRate1Y    float64            `json:"inflation_rate_1y"`
	Health             int                `json:"health"`
	Supply             *Supply            `json:"supply"`
	Status             Status             `json:"status"`
}

type Deployment struct {
	Protocol    string `json:"protocol"`
	Version     int    `json:"version"`      // protocol version sequence on indexed chain
	StartHeight int64  `json:"start_height"` // first block on indexed chain
	EndHeight   int64  `json:"end_height"`   // last block on indexed chain or -1
}

type Status struct {
	Status    string  `json:"status"` // loading, connecting, stopping, stopped, waiting, syncing, synced, failed
	Blocks    int64   `json:"blocks"`
	Finalized int64   `json:"finalized"`
	Indexed   int64   `json:"indexed"`
	Progress  float64 `json:"progress"`

	columns []string
}

func (s *Status) WithColumns(cols ...string) *Status {
	s.columns = cols
	return s
}

func (s *Status) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if len(data) == 2 {
		return nil
	}
	if data[0] == '[' {
		return s.UnmarshalJSONBrief(data)
	}
	type Alias *Status
	return json.Unmarshal(data, Alias(s))
}

func (s *Status) UnmarshalJSONBrief(data []byte) error {
	st := Status{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	unpacked := make([]interface{}, 0)
	err := dec.Decode(&unpacked)
	if err != nil {
		return err
	}
	for i, v := range s.columns {
		f := unpacked[i]
		if f == nil {
			continue
		}
		switch v {
		case "status":
			st.Status = f.(string)
		case "blocks":
			st.Blocks, err = strconv.ParseInt(f.(json.Number).String(), 10, 64)
		case "indexed":
			st.Indexed, err = strconv.ParseInt(f.(json.Number).String(), 10, 64)
		case "progress":
			st.Progress, err = f.(json.Number).Float64()
		}
		if err != nil {
			return err
		}
	}
	*s = st
	return nil
}

func (c *Client) GetStatus(ctx context.Context) (*Status, error) {
	s := &Status{}
	if err := c.get(ctx, "/explorer/status", nil, s); err != nil {
		return nil, err
	}
	return s, nil
}

func (c *Client) GetTip(ctx context.Context) (*Tip, error) {
	tip := &Tip{}
	if err := c.get(ctx, "/explorer/tip", nil, tip); err != nil {
		return nil, err
	}
	return tip, nil
}

func (c *Client) ListProtocols(ctx context.Context) ([]Deployment, error) {
	protos := make([]Deployment, 0)
	if err := c.get(ctx, "/explorer/protocols", nil, &protos); err != nil {
		return nil, err
	}
	return protos, nil
}

type BlockchainConfig struct {
	Name                   string  `json:"name"`
	Network                string  `json:"network"`
	Symbol                 string  `json:"symbol"`
	ChainId                string  `json:"chain_id"`
	Deployment             int     `json:"deployment"`
	Version                int     `json:"version"`
	Protocol               string  `json:"protocol"`
	StartHeight            int64   `json:"start_height"`
	EndHeight              int64   `json:"end_height"`
	Decimals               int     `json:"decimals"`
	MinimalStake           float64 `json:"minimal_stake"`
	PreservedCycles        int64   `json:"preserved_cycles"`
	MinimalBlockDelay      int     `json:"minimal_block_delay"`
	DelayIncrementPerRound int     `json:"delay_increment_per_round"`
}

func (c *Client) GetConfig(ctx context.Context) (*BlockchainConfig, error) {
	config := &BlockchainConfig{}
	if err := c.get(ctx, "/explorer/config/head", nil, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Client) GetConfigHeight(ctx context.Context, height int64) (*BlockchainConfig, error) {
	config := &BlockchainConfig{}
	if err := c.get(ctx, "/explorer/config/"+strconv.FormatInt(height, 10), nil, config); err != nil {
		return nil, err
	}
	return config, nil
}

type Supply struct {
	RowId               uint64    `json:"row_id"`
	Height              int64     `json:"height"`
	Cycle               int64     `json:"cycle"`
	Timestamp           time.Time `json:"time"`
	Total               float64   `json:"total"`
	Activated           float64   `json:"activated"`
	Unclaimed           float64   `json:"unclaimed"`
	Circulating         float64   `json:"circulating"`
	Liquid              float64   `json:"liquid"`
	Delegated           float64   `json:"delegated"`
	Staking             float64   `json:"staking"`
	Shielded            float64   `json:"shielded"`
	ActiveStake         float64   `json:"active_stake"`
	ActiveDelegated     float64   `json:"active_delegated"`
	ActiveStaking       float64   `json:"active_staking"`
	InactiveDelegated   float64   `json:"inactive_delegated"`
	InactiveStaking     float64   `json:"inactive_staking"`
	Minted              float64   `json:"minted"`
	MintedBaking        float64   `json:"minted_baking"`
	MintedEndorsing     float64   `json:"minted_endorsing"`
	MintedSeeding       float64   `json:"minted_seeding"`
	MintedAirdrop       float64   `json:"minted_airdrop"`
	MintedSubsidy       float64   `json:"minted_subsidy"`
	Burned              float64   `json:"burned"`
	BurnedDoubleBaking  float64   `json:"burned_double_baking"`
	BurnedDoubleEndorse float64   `json:"burned_double_endorse"`
	BurnedOrigination   float64   `json:"burned_origination"`
	BurnedAllocation    float64   `json:"burned_allocation"`
	BurnedStorage       float64   `json:"burned_storage"`
	BurnedExplicit      float64   `json:"burned_explicit"`
	BurnedSeedMiss      float64   `json:"burned_seed_miss"`
	BurnedAbsence       float64   `json:"burned_absence"`
	BurnedRollup        float64   `json:"burned_rollup"`
	Frozen              float64   `json:"frozen"`
	FrozenDeposits      float64   `json:"frozen_deposits"`
	FrozenRewards       float64   `json:"frozen_rewards"`
	FrozenFees          float64   `json:"frozen_fees"`
	FrozenBonds         float64   `json:"frozen_bonds"`
}
