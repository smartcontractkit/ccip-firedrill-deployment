package shared

import (
	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/deployment"
)

const FiredrillEntrypointType deployment.ContractType = "FiredrillEntrypoint"

type FiredrillConfig struct {
	Version             semver.Version
	ChainSelector       uint64
	SourceChainSelector uint64
}

type ChainView struct {
	ChainSelector       uint64                             `json:"chainSelector,omitempty"`
	ChainID             string                             `json:"chainID,omitempty"`
	FiredrillEntrypoint map[string]FiredrillEntrypointView `json:"firedrillEntrypoint,omitempty"`
}

func NewChainView(chainSelector uint64) (*ChainView, error) {
	chainID, err := chain_selectors.GetChainIDFromSelector(chainSelector)
	if err != nil {
		return nil, err
	}
	return &ChainView{
		ChainSelector:       chainSelector,
		ChainID:             chainID,
		FiredrillEntrypoint: make(map[string]FiredrillEntrypointView),
	}, nil
}

type ContractMetaData struct {
	TypeAndVersion string `json:"typeAndVersion,omitempty"`
	Address        string `json:"address,omitempty"`
	Owner          string `json:"owner,omitempty"`
}

type FiredrillEntrypointView struct {
	ContractMetaData
	Token     string `json:"token"`
	OnRamp    string `json:"onRamp"`
	OffRamp   string `json:"offRamp"`
	FeeQuoter string `json:"feeQuoter"`
	Compound  string `json:"compound"`
	Receiver  string `json:"receiver"`
}
