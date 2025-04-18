// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package firedrill_compound

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type FiredrillCompound struct {
	ChainSelector uint64
	Owner         ag_solanago.PublicKey
	FeeQuoter     ag_solanago.PublicKey
	Token         ag_solanago.PublicKey
}

var FiredrillCompoundDiscriminator = [8]byte{254, 217, 204, 248, 103, 242, 213, 162}

func (obj FiredrillCompound) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FiredrillCompoundDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
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

func (obj *FiredrillCompound) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FiredrillCompoundDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[254 217 204 248 103 242 213 162]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
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

type DestChain struct {
	Version       uint8
	ChainSelector uint64
	State         DestChainState
	Config        DestChainConfig
}

var DestChainDiscriminator = [8]byte{77, 18, 241, 132, 212, 54, 218, 16}

func (obj DestChain) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DestChainDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `State` param:
	err = encoder.Encode(obj.State)
	if err != nil {
		return err
	}
	// Serialize `Config` param:
	err = encoder.Encode(obj.Config)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChain) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DestChainDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[77 18 241 132 212 54 218 16]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `State`:
	err = decoder.Decode(&obj.State)
	if err != nil {
		return err
	}
	// Deserialize `Config`:
	err = decoder.Decode(&obj.Config)
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Version            uint8
	DefaultCodeVersion CodeVersion
	SvmChainSelector   uint64
	Owner              ag_solanago.PublicKey
	ProposedOwner      ag_solanago.PublicKey
	FeeQuoter          ag_solanago.PublicKey
	RmnRemote          ag_solanago.PublicKey
	LinkTokenMint      ag_solanago.PublicKey
	FeeAggregator      ag_solanago.PublicKey
}

var ConfigDiscriminator = [8]byte{155, 12, 170, 224, 30, 250, 204, 130}

func (obj Config) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ConfigDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	// Serialize `DefaultCodeVersion` param:
	err = encoder.Encode(obj.DefaultCodeVersion)
	if err != nil {
		return err
	}
	// Serialize `SvmChainSelector` param:
	err = encoder.Encode(obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `ProposedOwner` param:
	err = encoder.Encode(obj.ProposedOwner)
	if err != nil {
		return err
	}
	// Serialize `FeeQuoter` param:
	err = encoder.Encode(obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Serialize `RmnRemote` param:
	err = encoder.Encode(obj.RmnRemote)
	if err != nil {
		return err
	}
	// Serialize `LinkTokenMint` param:
	err = encoder.Encode(obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Serialize `FeeAggregator` param:
	err = encoder.Encode(obj.FeeAggregator)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Config) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ConfigDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[155 12 170 224 30 250 204 130]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	// Deserialize `DefaultCodeVersion`:
	err = decoder.Decode(&obj.DefaultCodeVersion)
	if err != nil {
		return err
	}
	// Deserialize `SvmChainSelector`:
	err = decoder.Decode(&obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `ProposedOwner`:
	err = decoder.Decode(&obj.ProposedOwner)
	if err != nil {
		return err
	}
	// Deserialize `FeeQuoter`:
	err = decoder.Decode(&obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Deserialize `RmnRemote`:
	err = decoder.Decode(&obj.RmnRemote)
	if err != nil {
		return err
	}
	// Deserialize `LinkTokenMint`:
	err = decoder.Decode(&obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `FeeAggregator`:
	err = decoder.Decode(&obj.FeeAggregator)
	if err != nil {
		return err
	}
	return nil
}
