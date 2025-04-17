// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package firedrill_offramp

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initializes the CCIP Offramp Config account.
//
// The initialization of the Offramp is responsibility of Admin, nothing more than calling these
// initialization methods should be done first.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for initialization of the config.
// * `svm_chain_selector` - The chain selector for SVM.
// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
type InitializeConfig struct {
	SvmChainSelector *uint64

	// [0] = [WRITE] config
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [] systemProgram
	//
	// [3] = [] program
	//
	// [4] = [] programData
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeConfigInstructionBuilder creates a new `InitializeConfig` instruction builder.
func NewInitializeConfigInstructionBuilder() *InitializeConfig {
	nd := &InitializeConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetSvmChainSelector sets the "svmChainSelector" parameter.
func (inst *InitializeConfig) SetSvmChainSelector(svmChainSelector uint64) *InitializeConfig {
	inst.SvmChainSelector = &svmChainSelector
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *InitializeConfig) SetConfigAccount(config ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *InitializeConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitializeConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitializeConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProgramAccount sets the "program" account.
func (inst *InitializeConfig) SetProgramAccount(program ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *InitializeConfig) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProgramDataAccount sets the "programData" account.
func (inst *InitializeConfig) SetProgramDataAccount(programData ag_solanago.PublicKey) *InitializeConfig {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(programData)
	return inst
}

// GetProgramDataAccount gets the "programData" account.
func (inst *InitializeConfig) GetProgramDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst InitializeConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SvmChainSelector == nil {
			return errors.New("SvmChainSelector parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Program is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ProgramData is not set")
		}
	}
	return nil
}

func (inst *InitializeConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("SvmChainSelector", *inst.SvmChainSelector))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      program", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  programData", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SvmChainSelector` param:
	err = encoder.Encode(obj.SvmChainSelector)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SvmChainSelector`:
	err = decoder.Decode(&obj.SvmChainSelector)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeConfigInstruction declares a new InitializeConfig instruction with the provided parameters and accounts.
func NewInitializeConfigInstruction(
	// Parameters:
	svmChainSelector uint64,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	program ag_solanago.PublicKey,
	programData ag_solanago.PublicKey) *InitializeConfig {
	return NewInitializeConfigInstructionBuilder().
		SetSvmChainSelector(svmChainSelector).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetProgramAccount(program).
		SetProgramDataAccount(programData)
}
