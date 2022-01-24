package database

import (
	"github.com/disperze/wasmx/types"
)

// SaveCode allows to save the given code into the database.
func (db Db) SaveCode(code types.Code) error {
	stmt := `
INSERT INTO codes (code_id, creator, creation_time, height) 
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Sql.Exec(stmt, code.CodeID, code.Creator, code.CreatedTime, code.Height)
	return err
}

// HasCode verify if code is registered.
func (db Db) HasCode(id uint64) (bool, error) {
	var count int
	stmt := `SELECT COUNT(1) FROM codes WHERE code_id = $1`
	row := db.Sql.QueryRow(stmt, id)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, err
}

// SetCodeVersion update code version.
func (db Db) SetCodeVersion(id uint64, version string) error {
	stmt := `UPDATE codes SET version=$2 WHERE code_id = $1`
	_, err := db.Sql.Exec(stmt, id, version)
	return err
}
