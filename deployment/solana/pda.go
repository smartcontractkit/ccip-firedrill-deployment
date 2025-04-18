package solana

import (
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func FindFiredrillCompoundPDA(firedrillCompoundProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("compound")}, firedrillCompoundProgram)
}

func FindFiredrillOfframpPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("offramp")}, firedrillOfframpProgram)
}

func FindFiredrillOfframpReferenceAddressesPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("reference_addresses")}, firedrillOfframpProgram)
}

func FindFiredrillOfframpSourceChainPDA(chainSelector uint64, offrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("source_chain_state"), chainSelectorLE}, offrampProgram)
}

func FindFiredrillOfframpConfigPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, firedrillOfframpProgram)
}

func FindFiredrillEntrypointPDA(firedrillEntrypointProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("entrypoint")}, firedrillEntrypointProgram)
}

func FindFiredrillTokenPDA(firedrillOnrampProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("token")}, firedrillOnrampProgram)
}
