package evm

import (
	"context"
	"maps"
	"slices"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/environment/memory"
	"github.com/smartcontractkit/chainlink/v2/core/logger"

	firedrill_entrypoint_v1_5 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_off_ramp"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func TestDeployFiredrillContracts(t *testing.T) {
	lggr := logger.TestLogger(t)
	chains, _ := memory.NewMemoryChains(t, 2, 1)
	env := *deployment.NewEnvironment(
		memory.Memory,
		lggr,
		deployment.NewMemoryAddressBook(),
		nil,
		chains,
		map[uint64]deployment.SolChain{},
		map[uint64]deployment.AptosChain{},
		[]string{},
		nil,
		func() context.Context { return tests.Context(t) },
		deployment.XXXGenerateTestOCRSecrets(),
	)
	chainSels := env.AllChainSelectors()
	require.Len(t, chainSels, 2)
	chainSel := chainSels[0]
	sourceChainSel := chainSels[1]
	firedrillChangeset, err := DeployFiredrillContracts(env, shared.FiredrillConfig{
		Version:             deployment.Version1_5_0,
		ChainSelector:       chainSel,
		SourceChainSelector: sourceChainSel,
	})
	require.NoError(t, err)
	sourceChainAddr, err := firedrillChangeset.AddressBook.AddressesForChain(sourceChainSel)
	require.Error(t, err)
	assert.Empty(t, sourceChainAddr)
	chainAddr, err := firedrillChangeset.AddressBook.AddressesForChain(chainSel)
	require.NoError(t, err)
	assert.Len(t, chainAddr, 1)
	typeAndVersionList := slices.Collect(maps.Values(chainAddr))
	assert.Contains(t, typeAndVersionList, deployment.NewTypeAndVersion(shared.FiredrillEntrypointType, deployment.Version1_5_0))
}

func TestRegisterFiredrill(t *testing.T) {
	lggr := logger.TestLogger(t)
	chains, _ := memory.NewMemoryChains(t, 2, 1)
	env := *deployment.NewEnvironment(
		memory.Memory,
		lggr,
		deployment.NewMemoryAddressBook(),
		nil,
		chains,
		map[uint64]deployment.SolChain{},
		map[uint64]deployment.AptosChain{},
		[]string{},
		nil,
		func() context.Context { return tests.Context(t) },
		deployment.XXXGenerateTestOCRSecrets(),
	)
	chainSels := env.AllChainSelectors()
	require.Len(t, chainSels, 2)
	chainSel := chainSels[0]
	sourceChainSel := chainSels[1]
	firedrillChangeset, err := DeployFiredrillContracts(env, shared.FiredrillConfig{
		Version:             deployment.Version1_5_0,
		ChainSelector:       chainSel,
		SourceChainSelector: sourceChainSel,
	})
	require.NoError(t, err)
	err = env.ExistingAddresses.Merge(firedrillChangeset.AddressBook)
	require.NoError(t, err)
	firedrillEntrypointAddrRaw, err := deployment.SearchAddressBook(firedrillChangeset.AddressBook, chainSel, shared.FiredrillEntrypointType)
	require.NoError(t, err)
	firedrillEntrypointAddr := common.HexToAddress(firedrillEntrypointAddrRaw)
	firedrillEntrypoint, err := firedrill_entrypoint_v1_5.NewFiredrillEntrypoint(firedrillEntrypointAddr, env.Chains[chainSel].Client)
	require.NoError(t, err)
	firedrillOnRampAddr, err := firedrillEntrypoint.OnRamp(nil)
	require.NoError(t, err)
	firedrillOffRampAddr, err := firedrillEntrypoint.OffRamp(nil)
	require.NoError(t, err)
	firedrillOffRamp, err := firedrill_off_ramp.NewFiredrillOffRamp(firedrillOffRampAddr, env.Chains[chainSel].Client)
	require.NoError(t, err)
	offRampSetSink := make(chan *firedrill_off_ramp.FiredrillOffRampConfigSet)
	subscription, err := firedrillOffRamp.WatchConfigSet(nil, offRampSetSink)
	require.NoError(t, err)
	defer subscription.Unsubscribe()
	err = FiredrillRegisterContracts(env.Logger, env.ExistingAddresses, env.Chains[chainSel])
	require.NoError(t, err)
	timer := time.NewTimer(1 * time.Second)
	for {
		select {
		case ev := <-offRampSetSink:
			assert.Equal(t, ev.StaticConfig.ChainSelector, chainSel)
			assert.Equal(t, ev.StaticConfig.SourceChainSelector, chainSel)
			assert.Equal(t, ev.StaticConfig.OnRamp, firedrillOnRampAddr)
			assert.Equal(t, ev.StaticConfig.CommitStore, firedrillOffRampAddr)
			assert.Equal(t, ev.StaticConfig.RmnProxy, firedrillEntrypointAddr)
			assert.Equal(t, ev.StaticConfig.TokenAdminRegistry, firedrillEntrypointAddr)
			assert.Equal(t, ev.DynamicConfig.Router, firedrillEntrypointAddr)
			assert.Equal(t, ev.DynamicConfig.PriceRegistry, firedrillEntrypointAddr)
			return
		case <-timer.C:
			require.Fail(t, "haven't received all events")
		}
	}
}
