package deployment

import (
	"encoding/json"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	evm_deployment "github.com/smartcontractkit/ccip-firedrill-deployment/deployment/evm"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
	sol_deployment "github.com/smartcontractkit/ccip-firedrill-deployment/deployment/solana"
)

var _ deployment.ViewState = CCIPViewFiredrill

type CCIPFiredrillView struct {
	Chains    map[string]*shared.ChainView `json:"chains,omitempty"`
	SolChains map[string]*shared.ChainView `json:"solChains,omitempty"`
}

func (v CCIPFiredrillView) MarshalJSON() ([]byte, error) {
	// Alias to avoid recursive calls
	type Alias CCIPFiredrillView
	return json.MarshalIndent(&struct{ Alias }{Alias: Alias(v)}, "", " ")
}

func CCIPViewFiredrill(e deployment.Environment) (json.Marshaler, error) {
	evmView, err := evm_deployment.EVMViewFiredrill(e)
	if err != nil {
		return nil, err
	}
	solView, err := sol_deployment.SolViewFiredrill(e)
	if err != nil {
		return nil, err
	}
	return CCIPFiredrillView{
		Chains:    evmView,
		SolChains: solView,
	}, nil
}
