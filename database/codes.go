package database

import (
	"github.com/disperze/wasmx/types"
)

// SaveCode allows to save the given code into the database.
func (db Db) SaveCode(code types.Code) error {
	stmt := `
INSERT INTO codes (code_id, creator, hash, size, creation_time, height) 
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Sql.Exec(stmt, code.CodeID, code.Creator, code.Hash, code.Size, code.CreatedTime, code.Height)
	return err
}

// GetCodeData get code data.
func (db Db) GetCodeData(codeID uint64) (*types.CodeData, error) {
	var data types.CodeData
	stmt := `SELECT version, ibc, cw20 FROM codes WHERE code_id = $1 LIMIT 1`
	err := db.Sql.QueryRow(stmt, codeID).Scan(&data.Version, &data.IBC, &data.CW20)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// SetCodeData update code data.
func (db Db) SetCodeData(data types.CodeData) error {
	stmt := `UPDATE codes SET version=$2, ibc=$3, cw20=$4 WHERE code_id = $1`
	_, err := db.Sql.Exec(stmt, data.CodeID, data.Version, data.IBC, data.CW20)
	return err
}
