package shared

import (
	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink/deployment"
)

const FiredrillEntrypointType deployment.ContractType = "FiredrillEntrypoint"

type FiredrillConfig struct {
	Version             semver.Version
	ChainSelector       uint64
	SourceChainSelector uint64
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
	Receiver  string `json:"receiver"`
}
