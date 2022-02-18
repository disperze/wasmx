package database

import (
	"github.com/disperze/wasmx/types"
)

// SaveCode allows to save the given code into the database.
func (db Db) SaveCode(code types.Code) error {
	stmt := `
INSERT INTO codes (code_id, creator, creation_time, height) 
VALUES ($1, $2, $3, $4)`
	_, err := db.Sql.Exec(stmt, code.CodeID, code.Creator, code.CreatedTime, code.Height)
	return err
}
