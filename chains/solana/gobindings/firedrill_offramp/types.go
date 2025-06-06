// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package firedrill_offramp

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type PriceUpdates struct {
	TokenPriceUpdates []TokenPriceUpdate
	GasPriceUpdates   []GasPriceUpdate
}

func (obj PriceUpdates) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenPriceUpdates` param:
	err = encoder.Encode(obj.TokenPriceUpdates)
	if err != nil {
		return err
	}
	// Serialize `GasPriceUpdates` param:
	err = encoder.Encode(obj.GasPriceUpdates)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PriceUpdates) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenPriceUpdates`:
	err = decoder.Decode(&obj.TokenPriceUpdates)
	if err != nil {
		return err
	}
	// Deserialize `GasPriceUpdates`:
	err = decoder.Decode(&obj.GasPriceUpdates)
	if err != nil {
		return err
	}
	return nil
}

type TokenPriceUpdate struct {
	SourceToken ag_solanago.PublicKey
	UsdPerToken [28]uint8
}

func (obj TokenPriceUpdate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SourceToken` param:
	err = encoder.Encode(obj.SourceToken)
	if err != nil {
		return err
	}
	// Serialize `UsdPerToken` param:
	err = encoder.Encode(obj.UsdPerToken)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenPriceUpdate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SourceToken`:
	err = decoder.Decode(&obj.SourceToken)
	if err != nil {
		return err
	}
	// Deserialize `UsdPerToken`:
	err = decoder.Decode(&obj.UsdPerToken)
	if err != nil {
		return err
	}
	return nil
}

type GasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     [28]uint8
}

func (obj GasPriceUpdate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `UsdPerUnitGas` param:
	err = encoder.Encode(obj.UsdPerUnitGas)
	if err != nil {
		return err
	}
	return nil
}

func (obj *GasPriceUpdate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `UsdPerUnitGas`:
	err = decoder.Decode(&obj.UsdPerUnitGas)
	if err != nil {
		return err
	}
	return nil
}

type MerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       []byte
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]uint8
}

func (obj MerkleRoot) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SourceChainSelector` param:
	err = encoder.Encode(obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Serialize `OnRampAddress` param:
	err = encoder.Encode(obj.OnRampAddress)
	if err != nil {
		return err
	}
	// Serialize `MinSeqNr` param:
	err = encoder.Encode(obj.MinSeqNr)
	if err != nil {
		return err
	}
	// Serialize `MaxSeqNr` param:
	err = encoder.Encode(obj.MaxSeqNr)
	if err != nil {
		return err
	}
	// Serialize `MerkleRoot` param:
	err = encoder.Encode(obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MerkleRoot) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SourceChainSelector`:
	err = decoder.Decode(&obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `OnRampAddress`:
	err = decoder.Decode(&obj.OnRampAddress)
	if err != nil {
		return err
	}
	// Deserialize `MinSeqNr`:
	err = decoder.Decode(&obj.MinSeqNr)
	if err != nil {
		return err
	}
	// Deserialize `MaxSeqNr`:
	err = decoder.Decode(&obj.MaxSeqNr)
	if err != nil {
		return err
	}
	// Deserialize `MerkleRoot`:
	err = decoder.Decode(&obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

type ConfigOcrPluginType struct {
	Discriminant uint8
}

func (obj ConfigOcrPluginType) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Discriminant` param:
	err = encoder.Encode(obj.Discriminant)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ConfigOcrPluginType) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Discriminant`:
	err = decoder.Decode(&obj.Discriminant)
	if err != nil {
		return err
	}
	return nil
}

type Ocr3ConfigInfo struct {
	ConfigDigest                   [32]uint8
	F                              uint8
	N                              uint8
	IsSignatureVerificationEnabled uint8
}

func (obj Ocr3ConfigInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ConfigDigest` param:
	err = encoder.Encode(obj.ConfigDigest)
	if err != nil {
		return err
	}
	// Serialize `F` param:
	err = encoder.Encode(obj.F)
	if err != nil {
		return err
	}
	// Serialize `N` param:
	err = encoder.Encode(obj.N)
	if err != nil {
		return err
	}
	// Serialize `IsSignatureVerificationEnabled` param:
	err = encoder.Encode(obj.IsSignatureVerificationEnabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Ocr3ConfigInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ConfigDigest`:
	err = decoder.Decode(&obj.ConfigDigest)
	if err != nil {
		return err
	}
	// Deserialize `F`:
	err = decoder.Decode(&obj.F)
	if err != nil {
		return err
	}
	// Deserialize `N`:
	err = decoder.Decode(&obj.N)
	if err != nil {
		return err
	}
	// Deserialize `IsSignatureVerificationEnabled`:
	err = decoder.Decode(&obj.IsSignatureVerificationEnabled)
	if err != nil {
		return err
	}
	return nil
}

type Ocr3Config struct {
	PluginType   ConfigOcrPluginType
	ConfigInfo   Ocr3ConfigInfo
	Signers      [16][20]uint8
	Transmitters [16][32]uint8
}

func (obj Ocr3Config) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PluginType` param:
	err = encoder.Encode(obj.PluginType)
	if err != nil {
		return err
	}
	// Serialize `ConfigInfo` param:
	err = encoder.Encode(obj.ConfigInfo)
	if err != nil {
		return err
	}
	// Serialize `Signers` param:
	err = encoder.Encode(obj.Signers)
	if err != nil {
		return err
	}
	// Serialize `Transmitters` param:
	err = encoder.Encode(obj.Transmitters)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Ocr3Config) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PluginType`:
	err = decoder.Decode(&obj.PluginType)
	if err != nil {
		return err
	}
	// Deserialize `ConfigInfo`:
	err = decoder.Decode(&obj.ConfigInfo)
	if err != nil {
		return err
	}
	// Deserialize `Signers`:
	err = decoder.Decode(&obj.Signers)
	if err != nil {
		return err
	}
	// Deserialize `Transmitters`:
	err = decoder.Decode(&obj.Transmitters)
	if err != nil {
		return err
	}
	return nil
}

type SourceChainConfig struct {
	IsEnabled                 bool
	IsRmnVerificationDisabled bool
	LaneCodeVersion           CodeVersion
	OnRamp                    OnRampAddress
}

func (obj SourceChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `IsEnabled` param:
	err = encoder.Encode(obj.IsEnabled)
	if err != nil {
		return err
	}
	// Serialize `IsRmnVerificationDisabled` param:
	err = encoder.Encode(obj.IsRmnVerificationDisabled)
	if err != nil {
		return err
	}
	// Serialize `LaneCodeVersion` param:
	err = encoder.Encode(obj.LaneCodeVersion)
	if err != nil {
		return err
	}
	// Serialize `OnRamp` param:
	err = encoder.Encode(obj.OnRamp)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SourceChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `IsEnabled`:
	err = decoder.Decode(&obj.IsEnabled)
	if err != nil {
		return err
	}
	// Deserialize `IsRmnVerificationDisabled`:
	err = decoder.Decode(&obj.IsRmnVerificationDisabled)
	if err != nil {
		return err
	}
	// Deserialize `LaneCodeVersion`:
	err = decoder.Decode(&obj.LaneCodeVersion)
	if err != nil {
		return err
	}
	// Deserialize `OnRamp`:
	err = decoder.Decode(&obj.OnRamp)
	if err != nil {
		return err
	}
	return nil
}

type OnRampAddress struct {
	Bytes [64]uint8
	Len   uint32
}

func (obj OnRampAddress) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bytes` param:
	err = encoder.Encode(obj.Bytes)
	if err != nil {
		return err
	}
	// Serialize `Len` param:
	err = encoder.Encode(obj.Len)
	if err != nil {
		return err
	}
	return nil
}

func (obj *OnRampAddress) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bytes`:
	err = decoder.Decode(&obj.Bytes)
	if err != nil {
		return err
	}
	// Deserialize `Len`:
	err = decoder.Decode(&obj.Len)
	if err != nil {
		return err
	}
	return nil
}

type SourceChainState struct {
	MinSeqNr uint64
}

func (obj SourceChainState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MinSeqNr` param:
	err = encoder.Encode(obj.MinSeqNr)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SourceChainState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MinSeqNr`:
	err = decoder.Decode(&obj.MinSeqNr)
	if err != nil {
		return err
	}
	return nil
}

type OcrPluginType ag_binary.BorshEnum

const (
	OcrPluginTypeCommit OcrPluginType = iota
	OcrPluginTypeExecution
)

func (value OcrPluginType) String() string {
	switch value {
	case OcrPluginTypeCommit:
		return "Commit"
	case OcrPluginTypeExecution:
		return "Execution"
	default:
		return ""
	}
}

type MessageExecutionState ag_binary.BorshEnum

const (
	MessageExecutionStateUntouched MessageExecutionState = iota
	MessageExecutionStateInProgress
	MessageExecutionStateSuccess
	MessageExecutionStateFailure
)

func (value MessageExecutionState) String() string {
	switch value {
	case MessageExecutionStateUntouched:
		return "Untouched"
	case MessageExecutionStateInProgress:
		return "InProgress"
	case MessageExecutionStateSuccess:
		return "Success"
	case MessageExecutionStateFailure:
		return "Failure"
	default:
		return ""
	}
}

type CodeVersion ag_binary.BorshEnum

const (
	CodeVersionDefault CodeVersion = iota
	CodeVersionV1
)

func (value CodeVersion) String() string {
	switch value {
	case CodeVersionDefault:
		return "Default"
	case CodeVersionV1:
		return "V1"
	default:
		return ""
	}
}
