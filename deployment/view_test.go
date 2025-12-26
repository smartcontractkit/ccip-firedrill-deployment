package deployment

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/stretchr/testify/require"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	changeset2 "github.com/smartcontractkit/chainlink/deployment/common/changeset"

	deploy "github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/logger"

	evm_deployment "github.com/smartcontractkit/ccip-firedrill-deployment/deployment/evm"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func TestCCIPViewFiredrill(t *testing.T) {
	lggr := logger.TestLogger(t)

	newEnv, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t,
			[]uint64{chainsel.TEST_90000001.Selector, chainsel.TEST_90000002.Selector, chainsel.TEST_90000003.Selector}),
		environment.WithLogger(lggr),
	)
	require.NoError(t, err)
	env := *newEnv

	chainSels := env.BlockChains.ListChainSelectors(cldf_chain.WithFamily(chainsel.FamilyEVM))
	require.Len(t, chainSels, 3)
	chainSel := chainSels[0]
	sourceChainSel := chainSels[1]
	env, err = changeset2.Apply(t, env, changeset2.Configure(evm_deployment.FiredrillDeployRegisterChangeSet{}, shared.FiredrillConfig{
		Version:             deploy.Version1_5_0,
		ChainSelector:       chainSel,
		SourceChainSelector: sourceChainSel,
	}))
	require.NoError(t, err)
	view, err := CCIPViewFiredrill(env)
	require.NoError(t, err)
	_, err = view.MarshalJSON()
	require.NoError(t, err)
}
