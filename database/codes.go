package database

import (
	"github.com/disperze/wasmx/types"
)

// SaveCode allows to save the given code into the database.
func (db Db) SaveCode(code types.Code) error {
	stmt := `
INSERT INTO codes (code_id, creator, hash, size, creation_time, height) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Sql.Exec(stmt, code.CodeID, code.Creator, code.Hash, code.Size, code.CreatedTime, code.Height)
	return err
}

// HasCodeVersion verify if code has version.
func (db Db) HasCodeVersion(id uint64) (bool, error) {
	var version string
	stmt := `SELECT version FROM codes WHERE code_id = $1`
	row := db.Sql.QueryRow(stmt, id)
	err := row.Scan(&version)
	if err != nil {
		return false, err
	}
	return version != "", err
}

// SetCodeVersion update code version.
func (db Db) SetCodeVersion(id uint64, version string) error {
	stmt := `UPDATE codes SET version=$2 WHERE code_id = $1`
	_, err := db.Sql.Exec(stmt, id, version)
	return err
}
