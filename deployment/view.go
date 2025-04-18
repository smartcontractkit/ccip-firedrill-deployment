package deployment

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/common/view/types"

	firedrill_entrypoint_v1_5 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_6/gethwrappers/firedrill_entrypoint"
)

var _ deployment.ViewState = CCIPViewFiredrill

type FiredrillEntrypointView struct {
	types.ContractMetaData
	Active   bool           `json:"active"`
	Token    common.Address `json:"token"`
	OnRamp   common.Address `json:"onRamp"`
	OffRamp  common.Address `json:"offRamp"`
	Compound common.Address `json:"compound"`
	Receiver common.Address `json:"receiver"`
}

type ChainView struct {
	ChainSelector       uint64                             `json:"chainSelector,omitempty"`
	ChainID             string                             `json:"chainID,omitempty"`
	FiredrillEntrypoint map[string]FiredrillEntrypointView `json:"firedrillEntrypoint,omitempty"`
}

func NewChainView(chain deployment.Chain) (*ChainView, error) {
	chainID, err := chain_selectors.GetChainIDFromSelector(chain.Selector)
	if err != nil {
		return nil, err
	}
	return &ChainView{
		ChainSelector:       chain.Selector,
		ChainID:             chainID,
		FiredrillEntrypoint: make(map[string]FiredrillEntrypointView),
	}, nil
}

type CCIPFiredrillView struct {
	Chains map[string]*ChainView `json:"chains,omitempty"`
}

func (v CCIPFiredrillView) MarshalJSON() ([]byte, error) {
	// Alias to avoid recursive calls
	type Alias CCIPFiredrillView
	return json.MarshalIndent(&struct{ Alias }{Alias: Alias(v)}, "", " ")
}

func CCIPViewFiredrill(e deployment.Environment) (json.Marshaler, error) {
	ab, err := e.ExistingAddresses.Addresses()
	if err != nil {
		return nil, err
	}
	chains := make(map[string]*ChainView)
	for chainSel, addresses := range ab {
		chain := e.Chains[chainSel]
		chainView, err := NewChainView(chain)
		if err != nil {
			return nil, err
		}
		for addressStr, typeAndVersion := range addresses {
			address := common.HexToAddress(addressStr)
			switch typeAndVersion.String() {
			case cldf.NewTypeAndVersion(FiredrillEntrypointType, deployment.Version1_5_0).String():
				contract, err := firedrill_entrypoint_v1_5.NewFiredrillEntrypoint(address, chain.Client)
				if err != nil {
					return nil, err
				}
				view, err := contractView(contract, address, typeAndVersion.String())
				if err != nil {
					return nil, err
				}
				chainView.FiredrillEntrypoint[addressStr] = view
			case cldf.NewTypeAndVersion(FiredrillEntrypointType, deployment.Version1_6_0).String():
				contract, err := firedrill_entrypoint.NewFiredrillEntrypoint(address, chain.Client)
				if err != nil {
					return nil, err
				}
				view, err := contractView(contract, address, typeAndVersion.String())
				if err != nil {
					return nil, err
				}
				chainView.FiredrillEntrypoint[addressStr] = view
			}
		}
		if len(chainView.FiredrillEntrypoint) > 0 {
			chains[chain.Name()] = chainView
		}
	}
	return CCIPFiredrillView{
		Chains: chains,
	}, nil
}

func contractView(contract FiredrillEntrypoint, address common.Address, typeAndVersion string) (FiredrillEntrypointView, error) {
	active, err := contract.IsActive(nil)
	if err != nil {
		return FiredrillEntrypointView{}, err
	}
	owner, err := contract.Owner(nil)
	if err != nil {
		return FiredrillEntrypointView{}, err
	}
	token, err := contract.Token(nil)
	if err != nil {
		return FiredrillEntrypointView{}, err
	}
	onramp, err := contract.OnRamp(nil)
	if err != nil {
		return FiredrillEntrypointView{}, err
	}
	offramp, err := contract.OffRamp(nil)
	if err != nil {
		return FiredrillEntrypointView{}, err
	}
	compound, err := contract.Compound(nil)
	if err != nil {
		return FiredrillEntrypointView{}, err
	}
	receiver, err := contract.Receiver(nil)
	if err != nil {
		return FiredrillEntrypointView{}, err
	}
	return FiredrillEntrypointView{
		ContractMetaData: types.ContractMetaData{
			TypeAndVersion: typeAndVersion,
			Address:        address,
			Owner:          owner,
		},
		Active:   active,
		Token:    token,
		OnRamp:   onramp,
		OffRamp:  offramp,
		Compound: compound,
		Receiver: receiver,
	}, nil
}
