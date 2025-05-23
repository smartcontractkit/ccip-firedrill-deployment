// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package failing_receiver

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CcipReceive is the `ccipReceive` instruction.
type CcipReceive struct {
	Message *Any2SVMMessage

	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCcipReceiveInstructionBuilder creates a new `CcipReceive` instruction builder.
func NewCcipReceiveInstructionBuilder() *CcipReceive {
	nd := &CcipReceive{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 0),
	}
	return nd
}

// SetMessage sets the "message" parameter.
func (inst *CcipReceive) SetMessage(message Any2SVMMessage) *CcipReceive {
	inst.Message = &message
	return inst
}

func (inst CcipReceive) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CcipReceive,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CcipReceive) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CcipReceive) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Message == nil {
			return errors.New("Message parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
	}
	return nil
}

func (inst *CcipReceive) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CcipReceive")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=0]").ParentFunc(func(accountsBranch ag_treeout.Branches) {})
				})
		})
}

func (obj CcipReceive) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CcipReceive) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	return nil
}

// NewCcipReceiveInstruction declares a new CcipReceive instruction with the provided parameters and accounts.
func NewCcipReceiveInstruction(
	// Parameters:
	message Any2SVMMessage) *CcipReceive {
	return NewCcipReceiveInstructionBuilder().
		SetMessage(message)
}
