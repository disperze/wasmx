package database

import (
	"fmt"

	"github.com/desmos-labs/juno/db"
	"github.com/desmos-labs/juno/db/postgresql"
	juno "github.com/desmos-labs/juno/types"
	"github.com/jmoiron/sqlx"
)

// Db represents a PostgreSQL database with expanded features.
// so that it can properly store posts and other Wasm-related data.
type Db struct {
	*postgresql.Database
	Sqlx *sqlx.DB
}

// Cast casts the given database to be a *Db
func Cast(database db.Database) *Db {
	wasmDb, ok := (database).(*Db)
	if !ok {
		panic(fmt.Errorf("database is not a WasmDB instance"))
	}
	return wasmDb
}

// Builder allows to create a new Db instance implementing the database.Builder type
func Builder(ctx *db.Context) (db.Database, error) {
	database, err := postgresql.Builder(ctx)
	if err != nil {
		return nil, err
	}

	psqlDb, ok := (database).(*postgresql.Database)
	if !ok {
		return nil, fmt.Errorf("invalid database type")
	}

	return &Db{
		Database: psqlDb,
		Sqlx:     sqlx.NewDb(psqlDb.Sql, "postgresql"),
	}, nil
}

// SaveTx overrides postgresql.Database to perform a no-op
func (db *Db) SaveTx(_ *juno.Tx) error {
	return nil
}

// HasValidator overrides postgresql.Database to perform a no-op
func (db *Db) HasValidator(_ string) (bool, error) {
	return true, nil
}

// SaveValidators overrides postgresql.Database to perform a no-op
func (db *Db) SaveValidators(_ []*juno.Validator) error {
	return nil
}

// SaveCommitSignatures overrides postgresql.Database to perform a no-op
func (db *Db) SaveCommitSignatures(_ []*juno.CommitSig) error {
	return nil
}

// SaveMessage overrides postgresql.Database to perform a no-op
func (db *Db) SaveMessage(_ *juno.Message) error {
	return nil
}
