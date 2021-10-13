package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

type Contract struct {
	*wasmtypes.ContractInfo
	Address     string
	CreatedTime string
}

func NewContract(contract *wasmtypes.ContractInfo, address string, created string) Contract {
	return Contract{
		ContractInfo: contract,
		Address:      address,
		CreatedTime:  created,
	}
}
