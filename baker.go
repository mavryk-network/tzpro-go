// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
    "context"
    "fmt"
    "strconv"
    "time"

    "blockwatch.cc/tzgo/tezos"
)

type Baker struct {
    Address           tezos.Address    `json:"address"`
    BakerSince        time.Time        `json:"baker_since_time"`
    BakerUntil        *time.Time       `json:"baker_until,omitempty"`
    GracePeriod       int64            `json:"grace_period"`
    BakerVersion      string           `json:"baker_version"`
    TotalBalance      float64          `json:"total_balance"`
    SpendableBalance  float64          `json:"spendable_balance"`
    FrozenBalance     float64          `json:"frozen_balance"`
    DelegatedBalance  float64          `json:"delegated_balance"`
    StakingBalance    float64          `json:"staking_balance"`
    StakingCapacity   float64          `json:"staking_capacity"`
    DepositsLimit     *float64         `json:"deposits_limit"`
    StakingShare      float64          `json:"staking_share"`
    ActiveDelegations int64            `json:"active_delegations"`
    IsFull            bool             `json:"is_full"`
    IsActive          bool             `json:"is_active"`
    Events            *BakerEvents     `json:"events,omitempty"`
    Stats             *BakerStatistics `json:"stats,omitempty"`
    Metadata          *Metadata        `json:"metadata,omitempty"`
}

type BakerStatistics struct {
    TotalRewardsEarned float64 `json:"total_rewards_earned"`
    TotalFeesEarned    float64 `json:"total_fees_earned"`
    TotalLost          float64 `json:"total_lost"`
    BlocksBaked        int64   `json:"blocks_baked"`
    BlocksProposed     int64   `json:"blocks_proposed"`
    SlotsEndorsed      int64   `json:"slots_endorsed"`
    AvgLuck64          int64   `json:"avg_luck_64"`
    AvgPerformance64   int64   `json:"avg_performance_64"`
    AvgContribution64  int64   `json:"avg_contribution_64"`
    NBakerOps          int64   `json:"n_baker_ops"`
    NProposal          int64   `json:"n_proposals"`
    NBallot            int64   `json:"n_ballots"`
    NEndorsement       int64   `json:"n_endorsements"`
    NPreendorsement    int64   `json:"n_preendorsements"`
    NSeedNonce         int64   `json:"n_nonce_revelations"`
    N2Baking           int64   `json:"n_double_bakings"`
    N2Endorsement      int64   `json:"n_double_endorsements"`
    NSetDepositsLimit  int64   `json:"n_set_limits"`
}

type BakerEvents struct {
    LastBakeHeight    int64     `json:"last_bake_height"`
    LastBakeBlock     string    `json:"last_bake_block"`
    LastBakeTime      time.Time `json:"last_bake_time"`
    LastEndorseHeight int64     `json:"last_endorse_height"`
    LastEndorseBlock  string    `json:"last_endorse_block"`
    LastEndorseTime   time.Time `json:"last_endorse_time"`
    NextBakeHeight    int64     `json:"next_bake_height"`
    NextBakeTime      time.Time `json:"next_bake_time"`
    NextEndorseHeight int64     `json:"next_endorse_height"`
    NextEndorseTime   time.Time `json:"next_endorse_time"`
}

type CycleIncome struct {
    Address          tezos.Address `json:"address"`
    Cycle            int64         `json:"cycle"`
    Rolls            int64         `json:"snapshot_rolls"`
    Balance          float64       `json:"own_balance"`
    Delegated        float64       `json:"delegated_balance"`
    Staking          float64       `json:"staking_balance"`
    NDelegations     int64         `json:"n_delegations"`
    NBakingRights    int64         `json:"n_baking_rights"`
    NEndorsingRights int64         `json:"n_endorsing_rights"`
    Luck             float64       `json:"luck"`
    LuckPct          int64         `json:"luck_percent"`
    ContributionPct  int64         `json:"contribution_percent"`
    PerformancePct   int64         `json:"performance_percent"`
    NBlocksBaked     int64         `json:"n_blocks_baked"`
    NBlocksProposed  int64         `json:"n_blocks_proposed"`
    NSlotsEndorsed   int64         `json:"n_slots_endorsed"`
    NSeedsRevealed   int64         `json:"n_seeds_revealed"`
    ExpectedIncome   float64       `json:"expected_income"`
    TotalIncome      float64       `json:"total_income"`
    TotalBonds       float64       `json:"total_bonds"`
    BakingIncome     float64       `json:"baking_income"`
    EndorsingIncome  float64       `json:"endorsing_income"`
    AccusationIncome float64       `json:"accusation_income"`
    SeedIncome       float64       `json:"seed_income"`
    FeesIncome       float64       `json:"fees_income"`
    TotalLoss        float64       `json:"total_loss"`
    AccusationLoss   float64       `json:"accusation_loss"`
    SeedLoss         float64       `json:"seed_loss"`
    EndorsingLoss    float64       `json:"endorsing_loss"`
}

type Delegator struct {
    Address tezos.Address `json:"address"`
    Balance float64       `json:"balance"`
}

type CycleSnapshot struct {
    BakeCycle        int64       `json:"baking_cycle"`
    Height           int64       `json:"snapshot_height"`
    Cycle            int64       `json:"snapshot_cycle"`
    Timestamp        time.Time   `json:"snapshot_time"`
    Index            int         `json:"snapshot_index"`
    Rolls            int64       `json:"snapshot_rolls"`
    StakingBalance   float64     `json:"staking_balance"`
    OwnBalance       float64     `json:"own_balance"`
    DelegatedBalance float64     `json:"delegated_balance"`
    NDelegations     int64       `json:"n_delegations"`
    Delegators       []Delegator `json:"delegators"`
}

type BakerParams struct {
    Params
}

func NewBakerParams() BakerParams {
    return BakerParams{NewParams()}
}

func (p BakerParams) WithLimit(v uint) BakerParams {
    p.Query.Set("limit", strconv.Itoa(int(v)))
    return p
}

func (p BakerParams) WithOffset(v uint) BakerParams {
    p.Query.Set("offset", strconv.Itoa(int(v)))
    return p
}

func (p BakerParams) WithCursor(v uint) BakerParams {
    p.Query.Set("cursor", strconv.Itoa(int(v)))
    return p
}

func (p BakerParams) WithMeta() BakerParams {
    p.Query.Set("meta", "1")
    return p
}

func (c *Client) GetBaker(ctx context.Context, addr tezos.Address, params BakerParams) (*Baker, error) {
    b := &Baker{}
    u := params.AppendQuery(fmt.Sprintf("/explorer/bakers/%s", addr))
    if err := c.get(ctx, u, nil, b); err != nil {
        return nil, err
    }
    return b, nil
}

func (c *Client) ListBakers(ctx context.Context, params BakerParams) ([]*Baker, error) {
    b := make([]*Baker, 0)
    u := params.AppendQuery("/explorer/bakers")
    if err := c.get(ctx, u, nil, &b); err != nil {
        return nil, err
    }
    return b, nil
}

func (c *Client) GetBakerBallot(ctx context.Context, addr tezos.Address, params BakerParams) ([]*Ballot, error) {
    cc := make([]*Ballot, 0)
    u := params.AppendQuery(fmt.Sprintf("/explorer/bakers/%s/votes", addr))
    if err := c.get(ctx, u, nil, &cc); err != nil {
        return nil, err
    }
    return cc, nil
}

func (c *Client) GetBakerEndorsements(ctx context.Context, addr tezos.Address, params OpParams) ([]*Op, error) {
    ops := make([]*Op, 0)
    u := params.AppendQuery(fmt.Sprintf("/explorer/bakers/%s/endorsements", addr))
    if err := c.get(ctx, u, nil, &ops); err != nil {
        return nil, err
    }
    return ops, nil
}

func (c *Client) GetBakerDelegations(ctx context.Context, addr tezos.Address, params OpParams) ([]*Op, error) {
    ops := make([]*Op, 0)
    u := params.AppendQuery(fmt.Sprintf("/explorer/bakers/%s/delegations", addr))
    if err := c.get(ctx, u, nil, &ops); err != nil {
        return nil, err
    }
    return ops, nil
}

func (c *Client) GetBakerRights(ctx context.Context, addr tezos.Address, cycle int64, params OpParams) (*CycleRights, error) {
    var r CycleRights
    u := params.AppendQuery(fmt.Sprintf("/explorer/bakers/%s/rights/%d", addr, cycle))
    if err := c.get(ctx, u, nil, &r); err != nil {
        return nil, err
    }
    return &r, nil
}

func (c *Client) GetBakerIncome(ctx context.Context, addr tezos.Address, cycle int64, params OpParams) (*CycleIncome, error) {
    var r CycleIncome
    u := params.AppendQuery(fmt.Sprintf("/explorer/bakers/%s/income/%d", addr, cycle))
    if err := c.get(ctx, u, nil, &r); err != nil {
        return nil, err
    }
    return &r, nil
}

func (c *Client) GetBakerSnapshot(ctx context.Context, addr tezos.Address, cycle int64, params OpParams) (*CycleSnapshot, error) {
    var r CycleSnapshot
    u := params.AppendQuery(fmt.Sprintf("/explorer/bakers/%s/delegators/%d", addr, cycle))
    if err := c.get(ctx, u, nil, &r); err != nil {
        return nil, err
    }
    return &r, nil
}
