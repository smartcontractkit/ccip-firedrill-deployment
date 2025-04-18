package solana

import "github.com/gagliardetto/solana-go"

func FindFiredrillCompoundPDA(firedrillCompoundProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("compound")}, firedrillCompoundProgram)
}

func FindFiredrillOfframpPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("offramp")}, firedrillOfframpProgram)
}

func FindFiredrillOfframpReferenceAddressesPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("reference_addresses")}, firedrillOfframpProgram)
}

func FindFiredrillOfframpSourceChainPDA(firedrillOfframpProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("source_chain_state")}, firedrillOfframpProgram)
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
