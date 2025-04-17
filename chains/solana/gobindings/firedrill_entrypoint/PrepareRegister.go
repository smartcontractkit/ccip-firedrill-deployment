// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package firedrill_entrypoint

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PrepareRegister is the `prepareRegister` instruction.
type PrepareRegister struct {

	// [0] = [WRITE] entrypoint
	//
	// [1] = [WRITE] offramp
	//
	// [2] = [SIGNER] owner
	//
	// [3] = [] offrampProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPrepareRegisterInstructionBuilder creates a new `PrepareRegister` instruction builder.
func NewPrepareRegisterInstructionBuilder() *PrepareRegister {
	nd := &PrepareRegister{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetEntrypointAccount sets the "entrypoint" account.
func (inst *PrepareRegister) SetEntrypointAccount(entrypoint ag_solanago.PublicKey) *PrepareRegister {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(entrypoint).WRITE()
	return inst
}

// GetEntrypointAccount gets the "entrypoint" account.
func (inst *PrepareRegister) GetEntrypointAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOfframpAccount sets the "offramp" account.
func (inst *PrepareRegister) SetOfframpAccount(offramp ag_solanago.PublicKey) *PrepareRegister {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(offramp).WRITE()
	return inst
}

// GetOfframpAccount gets the "offramp" account.
func (inst *PrepareRegister) GetOfframpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *PrepareRegister) SetOwnerAccount(owner ag_solanago.PublicKey) *PrepareRegister {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *PrepareRegister) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOfframpProgramAccount sets the "offrampProgram" account.
func (inst *PrepareRegister) SetOfframpProgramAccount(offrampProgram ag_solanago.PublicKey) *PrepareRegister {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(offrampProgram)
	return inst
}

// GetOfframpProgramAccount gets the "offrampProgram" account.
func (inst *PrepareRegister) GetOfframpProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst PrepareRegister) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PrepareRegister,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PrepareRegister) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PrepareRegister) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Entrypoint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Offramp is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OfframpProgram is not set")
		}
	}
	return nil
}

func (inst *PrepareRegister) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PrepareRegister")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    entrypoint", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       offramp", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("offrampProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj PrepareRegister) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *PrepareRegister) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewPrepareRegisterInstruction declares a new PrepareRegister instruction with the provided parameters and accounts.
func NewPrepareRegisterInstruction(
	// Accounts:
	entrypoint ag_solanago.PublicKey,
	offramp ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	offrampProgram ag_solanago.PublicKey) *PrepareRegister {
	return NewPrepareRegisterInstructionBuilder().
		SetEntrypointAccount(entrypoint).
		SetOfframpAccount(offramp).
		SetOwnerAccount(owner).
		SetOfframpProgramAccount(offrampProgram)
}
