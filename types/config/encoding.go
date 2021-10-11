package config

import (
	wasmapp "github.com/CosmWasm/wasmd/app"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
)

// MakeEncodingConfig creates an EncodingConfig to properly handle all the messages
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	wasmapp.ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	wasmapp.ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
