package deployment

import (
	"context"
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	deploy "github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/environment/memory"
	"github.com/smartcontractkit/chainlink/v2/core/logger"

	evm_deployment "github.com/smartcontractkit/ccip-firedrill-deployment/deployment/evm"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func TestCCIPViewFiredrill(t *testing.T) {
	lggr := logger.TestLogger(t)
	chains, _ := memory.NewMemoryChains(t, 3, 1)

	blockChains := make(map[uint64]cldf_chain.BlockChain)
	for _, chain := range chains {
		blockChains[chain.ChainSelector()] = chain
	}

	env := *deployment.NewEnvironment(
		memory.Memory,
		lggr,
		deployment.NewMemoryAddressBook(),
		nil,
		[]string{},
		nil,
		func() context.Context { return tests.Context(t) },
		deployment.XXXGenerateTestOCRSecrets(),
		cldf_chain.NewBlockChains(blockChains),
	)
	chainSels := env.BlockChains.ListChainSelectors(cldf_chain.WithFamily(chainsel.FamilyEVM))
	require.Len(t, chainSels, 3)
	chainSel := chainSels[0]
	sourceChainSel := chainSels[1]
	changeset, err := evm_deployment.DeployFiredrillContracts(env, shared.FiredrillConfig{
		Version:             deploy.Version1_5_0,
		ChainSelector:       chainSel,
		SourceChainSelector: sourceChainSel,
	})
	require.NoError(t, err)
	err = env.ExistingAddresses.Merge(changeset.AddressBook)
	require.NoError(t, err)
	view, err := CCIPViewFiredrill(env)
	require.NoError(t, err)
	_, err = view.MarshalJSON()
	require.NoError(t, err)
}
