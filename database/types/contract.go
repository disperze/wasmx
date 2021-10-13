package types

// ContractRow represents a single PostgreSQL row containing the data of a contract
type ContractRow struct {
	CodeID       uint64 `db:"code_id"`
	Address      string `db:"address"`
	Creator      string `db:"creator"`
	Admin        string `db:"admin"`
	Label        string `db:"label"`
	CreationTime string `db:"creation_time"`
	Height       int64  `db:"height"`
}
