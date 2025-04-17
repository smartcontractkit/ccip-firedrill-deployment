// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package firedrill_entrypoint

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DrillPendingCommitQueueTxSpike is the `drillPendingCommitQueueTxSpike` instruction.
type DrillPendingCommitQueueTxSpike struct {
	From *uint8
	To   *uint8

	// [0] = [WRITE] entrypoint
	//
	// [1] = [SIGNER] owner
	//
	// [2] = [WRITE] onramp
	//
	// [3] = [WRITE] offramp
	//
	// [4] = [WRITE] compound
	//
	// [5] = [] onrampProgram
	//
	// [6] = [] offrampProgram
	//
	// [7] = [] compoundProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDrillPendingCommitQueueTxSpikeInstructionBuilder creates a new `DrillPendingCommitQueueTxSpike` instruction builder.
func NewDrillPendingCommitQueueTxSpikeInstructionBuilder() *DrillPendingCommitQueueTxSpike {
	nd := &DrillPendingCommitQueueTxSpike{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetFrom sets the "from" parameter.
func (inst *DrillPendingCommitQueueTxSpike) SetFrom(from uint8) *DrillPendingCommitQueueTxSpike {
	inst.From = &from
	return inst
}

// SetTo sets the "to" parameter.
func (inst *DrillPendingCommitQueueTxSpike) SetTo(to uint8) *DrillPendingCommitQueueTxSpike {
	inst.To = &to
	return inst
}

// SetEntrypointAccount sets the "entrypoint" account.
func (inst *DrillPendingCommitQueueTxSpike) SetEntrypointAccount(entrypoint ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(entrypoint).WRITE()
	return inst
}

// GetEntrypointAccount gets the "entrypoint" account.
func (inst *DrillPendingCommitQueueTxSpike) GetEntrypointAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *DrillPendingCommitQueueTxSpike) SetOwnerAccount(owner ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *DrillPendingCommitQueueTxSpike) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOnrampAccount sets the "onramp" account.
func (inst *DrillPendingCommitQueueTxSpike) SetOnrampAccount(onramp ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(onramp).WRITE()
	return inst
}

// GetOnrampAccount gets the "onramp" account.
func (inst *DrillPendingCommitQueueTxSpike) GetOnrampAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOfframpAccount sets the "offramp" account.
func (inst *DrillPendingCommitQueueTxSpike) SetOfframpAccount(offramp ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(offramp).WRITE()
	return inst
}

// GetOfframpAccount gets the "offramp" account.
func (inst *DrillPendingCommitQueueTxSpike) GetOfframpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCompoundAccount sets the "compound" account.
func (inst *DrillPendingCommitQueueTxSpike) SetCompoundAccount(compound ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(compound).WRITE()
	return inst
}

// GetCompoundAccount gets the "compound" account.
func (inst *DrillPendingCommitQueueTxSpike) GetCompoundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetOnrampProgramAccount sets the "onrampProgram" account.
func (inst *DrillPendingCommitQueueTxSpike) SetOnrampProgramAccount(onrampProgram ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(onrampProgram)
	return inst
}

// GetOnrampProgramAccount gets the "onrampProgram" account.
func (inst *DrillPendingCommitQueueTxSpike) GetOnrampProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetOfframpProgramAccount sets the "offrampProgram" account.
func (inst *DrillPendingCommitQueueTxSpike) SetOfframpProgramAccount(offrampProgram ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(offrampProgram)
	return inst
}

// GetOfframpProgramAccount gets the "offrampProgram" account.
func (inst *DrillPendingCommitQueueTxSpike) GetOfframpProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCompoundProgramAccount sets the "compoundProgram" account.
func (inst *DrillPendingCommitQueueTxSpike) SetCompoundProgramAccount(compoundProgram ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(compoundProgram)
	return inst
}

// GetCompoundProgramAccount gets the "compoundProgram" account.
func (inst *DrillPendingCommitQueueTxSpike) GetCompoundProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst DrillPendingCommitQueueTxSpike) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DrillPendingCommitQueueTxSpike,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DrillPendingCommitQueueTxSpike) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DrillPendingCommitQueueTxSpike) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.From == nil {
			return errors.New("From parameter is not set")
		}
		if inst.To == nil {
			return errors.New("To parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Entrypoint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Onramp is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Offramp is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Compound is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.OnrampProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.OfframpProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CompoundProgram is not set")
		}
	}
	return nil
}

func (inst *DrillPendingCommitQueueTxSpike) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DrillPendingCommitQueueTxSpike")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("From", *inst.From))
						paramsBranch.Child(ag_format.Param("  To", *inst.To))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     entrypoint", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         onramp", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        offramp", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       compound", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  onrampProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta(" offrampProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("compoundProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj DrillPendingCommitQueueTxSpike) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `From` param:
	err = encoder.Encode(obj.From)
	if err != nil {
		return err
	}
	// Serialize `To` param:
	err = encoder.Encode(obj.To)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DrillPendingCommitQueueTxSpike) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `From`:
	err = decoder.Decode(&obj.From)
	if err != nil {
		return err
	}
	// Deserialize `To`:
	err = decoder.Decode(&obj.To)
	if err != nil {
		return err
	}
	return nil
}

// NewDrillPendingCommitQueueTxSpikeInstruction declares a new DrillPendingCommitQueueTxSpike instruction with the provided parameters and accounts.
func NewDrillPendingCommitQueueTxSpikeInstruction(
	// Parameters:
	from uint8,
	to uint8,
	// Accounts:
	entrypoint ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	onramp ag_solanago.PublicKey,
	offramp ag_solanago.PublicKey,
	compound ag_solanago.PublicKey,
	onrampProgram ag_solanago.PublicKey,
	offrampProgram ag_solanago.PublicKey,
	compoundProgram ag_solanago.PublicKey) *DrillPendingCommitQueueTxSpike {
	return NewDrillPendingCommitQueueTxSpikeInstructionBuilder().
		SetFrom(from).
		SetTo(to).
		SetEntrypointAccount(entrypoint).
		SetOwnerAccount(owner).
		SetOnrampAccount(onramp).
		SetOfframpAccount(offramp).
		SetCompoundAccount(compound).
		SetOnrampProgramAccount(onrampProgram).
		SetOfframpProgramAccount(offrampProgram).
		SetCompoundProgramAccount(compoundProgram)
}
