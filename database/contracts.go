package database

import (
	"github.com/disperze/wasmx/types"
)

// SaveContract allows to save the given contract into the database.
func (db Db) SaveContract(contract types.Contract, gas, fees int64) error {
	stmt := `
INSERT INTO contracts (code_id, address, creator, admin, label, creation_time, height, ibc, gas, fees) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := db.Sql.Exec(stmt, contract.CodeID, contract.Address, contract.Creator, contract.Admin, contract.Label, contract.CreatedTime, contract.Created.BlockHeight, contract.Ibc, gas, fees)
	return err
}

// UpdateContractStats update stats by contract call.
func (db Db) UpdateContractStats(contract string, tx, gas, fees int64) error {
	stmt := `
UPDATE contracts SET tx=tx+$2, gas=gas+$3, fees=fees+$4 WHERE address = $1`
	_, err := db.Sql.Exec(stmt, contract, tx, gas, fees)
	return err
}
