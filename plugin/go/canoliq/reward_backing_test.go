package canoliq

import (
	"testing"

	"github.com/canopy-network/go-plugin/contract"
)

// reward_backing_test.go pins how much of a block's committee reward the cCNPY
// pool receives when the pool funded only part of the bonded position.
//
// ProcessRewards observes growth of the WHOLE bonded committee stake and
// credits the user slice to cCNPY holders. That is exact for the liquid-staking
// case the protocol is designed around, where the bonded position IS the pooled
// deposits. It is not exact when a validator bonds capital that never came
// through a deposit: the reward on that capital is credited to depositors too.
//
// These tests DOCUMENT CURRENT BEHAVIOUR rather than assert a desired fix, so
// they pass as-is. They exist because the behaviour was reached on mainnet and
// is not obvious from reading ProcessRewards. See the accompanying PR for the
// numbers and the remediation options; changing the split is a protocol
// decision, and TestOwnerlessBranchStopsOnceSupplyExists already fixes the
// opposite invariant (a live pool must NOT have its user slice routed away),
// so any fix has to reconcile the two deliberately.
//
// Observed on mainnet chain 29, 2026-09-23. An operator bonded ~8,235 CNPY of
// its own capital for the committee while the pool held a single 10 CNPY
// deposit. Every block credited the reward on all ~8,235 to the lone depositor.
const (
	mainnetBondedStake uint64 = 8_235_687_200 // operator-owned, not deposited
	mainnetPoolDeposit uint64 = 10_000_000    // the entire pool: one 10 CNPY deposit
	mainnetBlockReward uint64 = 11_900_000    // ~11.9 CNPY/block, measured on chain
)

// seedPartiallyBackedReward registers one committee validator whose bonded
// stake grew by `reward`, while the pool holds only `pooled` CNPY -- i.e. the
// pool funded `pooled` of a `base` position.
func seedPartiallyBackedReward(t *testing.T, s *fakeStore, c *Canoliq, base, reward, pooled, ccnpy uint64) []byte {
	t.Helper()
	addr := addr20(0xC1)
	reg := &contract.ValidatorRegistry{Entries: []*contract.ValidatorRegistryEntry{{Address: addr, Stake: base}}}
	s.set(KeyForValidatorRegistry(), mustMarshal(reg))
	setCommitteeStake(s, c, addr, base+reward)
	g := loadGlobals(t, s)
	g.GenesisComplete = true
	g.LastProcessedRewardPool = base
	g.TotalPooledCnpy = pooled
	g.TotalCcnpySupply = ccnpy
	s.set(KeyForGlobals(), mustMarshal(g))
	seedEscrow(s, pooled)
	return addr
}

// userSliceFor recomputes the slice ProcessRewards routes to the pool, so the
// expectations below track the configured split instead of hard-coded numbers.
func userSliceFor(reward uint64) uint64 {
	p := DefaultParams()
	fee := FeeOnReward(reward, p.FeeBps)
	split := SplitFee(fee, &FeeSplitParams{
		UserRebateBps: p.UserRebateBps,
		TreasuryBps:   p.TreasuryBps,
		ValidatorBps:  p.ValidatorBps,
		BuybackBps:    p.BuybackBps,
	})
	return (reward - fee) + split.UserRebate
}

// The pool receives the full user slice regardless of how little of the bonded
// position it funded. With the mainnet numbers the pool supplied 0.121% of the
// stake and still collected 100% of the user slice.
func TestRewardCreditsFullUserSliceRegardlessOfBacking(t *testing.T) {
	c, s := newTestCanoliq()
	seedPartiallyBackedReward(t, s, c, mainnetBondedStake, mainnetBlockReward, mainnetPoolDeposit, testLivePoolCcnpy)

	if err := c.ProcessRewards(&contract.PluginEndRequest{Height: 1}); err != nil {
		t.Fatalf("ProcessRewards: %v", err)
	}
	g := loadGlobals(t, s)
	credited := g.TotalPooledCnpy - mainnetPoolDeposit

	if want := userSliceFor(mainnetBlockReward); credited != want {
		t.Fatalf("credited %d, want %d (the whole user slice)", credited, want)
	}
	// What the pool's own capital actually earned, for contrast.
	backed := mulDiv(userSliceFor(mainnetBlockReward), mainnetPoolDeposit, mainnetBondedStake)
	t.Logf("pool funded %d of %d bonded (%.4f%%); credited %d, proportional share would be %d (%.0fx)",
		mainnetPoolDeposit, mainnetBondedStake,
		100*float64(mainnetPoolDeposit)/float64(mainnetBondedStake),
		credited, backed, float64(credited)/float64(max64(backed, 1)))
}

// CNPY is conserved either way -- this is a misattribution between protocol
// buckets, not a mint. Worth pinning so a future fix keeps that property.
func TestRewardMisattributionStillConservesCnpy(t *testing.T) {
	c, s := newTestCanoliq()
	seedPartiallyBackedReward(t, s, c, mainnetBondedStake, mainnetBlockReward, mainnetPoolDeposit, testLivePoolCcnpy)
	treasuryBefore := DecodeUint64(s.get(KeyForTreasuryCNPY()))

	if err := c.ProcessRewards(&contract.PluginEndRequest{Height: 1}); err != nil {
		t.Fatalf("ProcessRewards: %v", err)
	}
	g := loadGlobals(t, s)
	total := (g.TotalPooledCnpy - mainnetPoolDeposit) +
		(DecodeUint64(s.get(KeyForTreasuryCNPY())) - treasuryBefore) +
		DecodeUint64(s.get(KeyForInsurancePool())) +
		DecodeUint64(s.get(KeyForBuybackPool())) +
		readAllValidatorIncentives(s)
	if total != mainnetBlockReward {
		t.Fatalf("conservation: got %d want %d", total, mainnetBlockReward)
	}
}

// The escrow invariant holds, so the credited CNPY is a real redeemable claim,
// not phantom balance -- which is what makes the misattribution consequential
// rather than cosmetic.
func TestRewardMisattributionProducesRedeemableClaim(t *testing.T) {
	c, s := newTestCanoliq()
	seedPartiallyBackedReward(t, s, c, mainnetBondedStake, mainnetBlockReward, mainnetPoolDeposit, testLivePoolCcnpy)

	if err := c.ProcessRewards(&contract.PluginEndRequest{Height: 1}); err != nil {
		t.Fatalf("ProcessRewards: %v", err)
	}
	g := loadGlobals(t, s)
	if got, want := readEscrow(s), g.TotalPooledCnpy+g.PendingRedemptionCnpy; got != want {
		t.Fatalf("escrow invariant: got %d want %d", got, want)
	}
	// One block already moves the redemption value of the whole supply.
	if g.TotalPooledCnpy <= mainnetPoolDeposit {
		t.Fatalf("expected the pool to grow: %d -> %d", mainnetPoolDeposit, g.TotalPooledCnpy)
	}
}

func max64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
