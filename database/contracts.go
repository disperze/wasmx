package database

import (
	"github.com/disperze/wasmx/types"
)

// SaveContract allows to save the given contract into the database.
func (db Db) SaveContract(contract types.Contract) error {
	stmt := `
INSERT INTO contracts (code_id, address, creator, admin, label, creation_time, height) 
VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Sql.Exec(stmt, contract.CodeID, contract.Address, contract.Creator, contract.Admin, contract.Label, contract.CreatedTime, contract.Created.BlockHeight)
	return err
}

// UpdateContractStats update stats by contract call.
func (db Db) UpdateContractStats(contract string, gas int64, fees int64) error {
	stmt := `
UPDATE contracts SET gas=gas+$2, fees=fees+$3 WHERE address = $1`
	_, err := db.Sql.Exec(stmt, contract, gas, fees)
	return err
}
