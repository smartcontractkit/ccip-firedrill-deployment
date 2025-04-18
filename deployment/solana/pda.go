package solana

import (
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func FindFiredrillCompoundPDA(firedrillCompoundProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("compound")}, firedrillCompoundProgram)
}

func FindFiredrillCompoundConfigPDA(firedrillCompoundProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, firedrillCompoundProgram)
}

func FindFiredrillCompoundDestChainPDA(firedrillCompoundProgram solana.PublicKey, chainSelector uint64) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("dest_chain_state"), chainSelectorLE}, firedrillCompoundProgram)
}

func FindFiredrillOfframpPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("offramp")}, firedrillOfframpProgram)
}

func FindFiredrillOfframpReferenceAddressesPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("reference_addresses")}, firedrillOfframpProgram)
}

func FindFiredrillOfframpSourceChainPDA(offrampProgram solana.PublicKey, chainSelector uint64) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("source_chain_state"), chainSelectorLE}, offrampProgram)
}

func FindFiredrillOfframpConfigPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, firedrillOfframpProgram)
}

func FindFiredrillFeeQuoterPDA(firedrillFeeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_quoter")}, firedrillFeeQuoterProgram)
}

func FindFiredrillFeeQuoterConfigPDA(firedrillFeeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, firedrillFeeQuoterProgram)
}

func FindFiredrillFeeQuoterDestChainPDA(firedrillFeeQuoterProgram solana.PublicKey, chainSelector uint64) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("dest_chain_state"), chainSelectorLE}, firedrillFeeQuoterProgram)
}

func FindFiredrillEntrypointPDA(firedrillEntrypointProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("entrypoint")}, firedrillEntrypointProgram)
}

func FindFiredrillTokenPDA(firedrillOnrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("token")}, firedrillOnrampProgram)
}
