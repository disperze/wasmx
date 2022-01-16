package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

// Contract type
type Contract struct {
	*wasmtypes.ContractInfo
	Address     string
	CreatedTime string
	Ibc         bool
}

// NewContract instance
func NewContract(contract *wasmtypes.ContractInfo, address string, created string, ibc bool) Contract {
	return Contract{
		ContractInfo: contract,
		Address:      address,
		CreatedTime:  created,
		Ibc:          ibc,
	}
}
