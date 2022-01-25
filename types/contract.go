package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// Contract type
type Contract struct {
	*wasmtypes.ContractInfo
	Address     string
	CreatedTime string
}

// NewContract instance
func NewContract(contract *wasmtypes.ContractInfo, address string, created string) Contract {
	return Contract{
		ContractInfo: contract,
		Address:      address,
		CreatedTime:  created,
	}
}
