package wasm

import (
	"github.com/disperze/wasmx/database"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/juno/modules"
	juno "github.com/desmos-labs/juno/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/profiles module handler
type Module struct {
	db *database.Db
}

// NewModule allows to build a new Module instance
func NewModule(db *database.Db) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "wasm"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	return HandleMsg(tx, index, msg, m.db)
}
