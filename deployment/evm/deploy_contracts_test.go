package evm

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	deploy "github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/logger"

	firedrill_entrypoint_v1_5 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_off_ramp"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func TestDeployFiredrillContracts(t *testing.T) {
	lggr := logger.TestLogger(t)
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t,
			[]uint64{chainsel.TEST_90000001.Selector, chainsel.TEST_90000002.Selector}),
		environment.WithLogger(lggr),
	)
	require.NoError(t, err)

	chainSels := env.BlockChains.ListChainSelectors(cldf_chain.WithFamily(chainsel.FamilyEVM))
	require.Len(t, chainSels, 2)
	chainSel := chainSels[0]
	sourceChainSel := chainSels[1]
	_, firedrillChangeset, err := DeployFiredrillContracts(*env, shared.FiredrillConfig{
		Version:             deploy.Version1_5_0,
		ChainSelector:       chainSel,
		SourceChainSelector: sourceChainSel,
	})
	require.NoError(t, err)
	refs := firedrillChangeset.DataStore.Addresses().Filter(func(refs []datastore.AddressRef) []datastore.AddressRef {
		res := make([]datastore.AddressRef, 0, len(refs))
		for _, ref := range refs {
			if ref.Type.String() == shared.FiredrillEntrypointType.String() && ref.Version.Equal(&deploy.Version1_5_0) {
				res = append(res, ref)
			}
		}
		return res
	})
	assert.Len(t, refs, 1)
}

func TestRegisterFiredrill(t *testing.T) {
	lggr := logger.TestLogger(t)
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t,
			[]uint64{chainsel.TEST_90000001.Selector, chainsel.TEST_90000002.Selector}),
		environment.WithLogger(lggr),
	)
	require.NoError(t, err)

	chainSels := env.BlockChains.ListChainSelectors(cldf_chain.WithFamily(chainsel.FamilyEVM))
	require.Len(t, chainSels, 2)
	chainSel := chainSels[0]
	sourceChainSel := chainSels[1]
	firedrillRef, output, err := DeployFiredrillContracts(*env, shared.FiredrillConfig{
		Version:             deploy.Version1_5_0,
		ChainSelector:       chainSel,
		SourceChainSelector: sourceChainSel,
	})
	require.NoError(t, err)
	evmChains := env.BlockChains.EVMChains()
	firedrillEntrypoint, err := firedrill_entrypoint_v1_5.NewFiredrillEntrypoint(common.HexToAddress(firedrillRef.Address), evmChains[chainSel].Client)
	require.NoError(t, err)
	firedrillOnRampAddr, err := firedrillEntrypoint.OnRamp(nil)
	require.NoError(t, err)
	firedrillOffRampAddr, err := firedrillEntrypoint.OffRamp(nil)
	require.NoError(t, err)
	firedrillOffRamp, err := firedrill_off_ramp.NewFiredrillOffRamp(firedrillOffRampAddr, evmChains[chainSel].Client)
	require.NoError(t, err)
	offRampSetSink := make(chan *firedrill_off_ramp.FiredrillOffRampConfigSet)
	subscription, err := firedrillOffRamp.WatchConfigSet(nil, offRampSetSink)
	require.NoError(t, err)
	defer subscription.Unsubscribe()
	err = FiredrillRegisterContracts(env.Logger, output.DataStore.Addresses(), env.BlockChains.EVMChains()[chainSel])
	require.NoError(t, err)
	timer := time.NewTimer(1 * time.Second)
	for {
		select {
		case ev := <-offRampSetSink:
			assert.Equal(t, ev.StaticConfig.ChainSelector, chainSel)
			assert.Equal(t, ev.StaticConfig.SourceChainSelector, chainSel)
			assert.Equal(t, ev.StaticConfig.OnRamp, firedrillOnRampAddr)
			assert.Equal(t, ev.StaticConfig.CommitStore, firedrillOffRampAddr)
			assert.Equal(t, ev.StaticConfig.RmnProxy.String(), firedrillRef.Address)
			assert.Equal(t, ev.StaticConfig.TokenAdminRegistry.String(), firedrillRef.Address)
			assert.Equal(t, ev.DynamicConfig.Router.String(), firedrillRef.Address)
			assert.Equal(t, ev.DynamicConfig.PriceRegistry.String(), firedrillRef.Address)
			return
		case <-timer.C:
			require.Fail(t, "haven't received all events")
		}
	}
}
