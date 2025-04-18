// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package firedrill_compound

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialize is the `initialize` instruction.
type Initialize struct {
	ChainSelector *uint64
	FeeQuoter     *ag_solanago.PublicKey
	Token         *ag_solanago.PublicKey

	// [0] = [WRITE] compound
	//
	// [1] = [WRITE] config
	//
	// [2] = [WRITE] destChain
	//
	// [3] = [WRITE, SIGNER] authority
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetChainSelector sets the "chainSelector" parameter.
func (inst *Initialize) SetChainSelector(chainSelector uint64) *Initialize {
	inst.ChainSelector = &chainSelector
	return inst
}

// SetFeeQuoter sets the "feeQuoter" parameter.
func (inst *Initialize) SetFeeQuoter(feeQuoter ag_solanago.PublicKey) *Initialize {
	inst.FeeQuoter = &feeQuoter
	return inst
}

// SetToken sets the "token" parameter.
func (inst *Initialize) SetToken(token ag_solanago.PublicKey) *Initialize {
	inst.Token = &token
	return inst
}

// SetCompoundAccount sets the "compound" account.
func (inst *Initialize) SetCompoundAccount(compound ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(compound).WRITE()
	return inst
}

// GetCompoundAccount gets the "compound" account.
func (inst *Initialize) GetCompoundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *Initialize) SetConfigAccount(config ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Initialize) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDestChainAccount sets the "destChain" account.
func (inst *Initialize) SetDestChainAccount(destChain ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(destChain).WRITE()
	return inst
}

// GetDestChainAccount gets the "destChain" account.
func (inst *Initialize) GetDestChainAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ChainSelector == nil {
			return errors.New("ChainSelector parameter is not set")
		}
		if inst.FeeQuoter == nil {
			return errors.New("FeeQuoter parameter is not set")
		}
		if inst.Token == nil {
			return errors.New("Token parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Compound is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.DestChain is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ChainSelector", *inst.ChainSelector))
						paramsBranch.Child(ag_format.Param("    FeeQuoter", *inst.FeeQuoter))
						paramsBranch.Child(ag_format.Param("        Token", *inst.Token))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     compound", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    destChain", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `FeeQuoter` param:
	err = encoder.Encode(obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `FeeQuoter`:
	err = decoder.Decode(&obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	chainSelector uint64,
	feeQuoter ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	// Accounts:
	compound ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	destChain ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetChainSelector(chainSelector).
		SetFeeQuoter(feeQuoter).
		SetToken(token).
		SetCompoundAccount(compound).
		SetConfigAccount(config).
		SetDestChainAccount(destChain).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
