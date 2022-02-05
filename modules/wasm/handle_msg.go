package wasm

import (
	"github.com/disperze/wasmx/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v2/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *wasmtypes.MsgStoreCode:
		return m.handleMsgStoreCode(tx, index, cosmosMsg)
	case *wasmtypes.MsgInstantiateContract:
		return m.handleMsgInstantiateContract(tx, index, cosmosMsg)
	}

	return nil
}

func (m *Module) handleMsgStoreCode(tx *juno.Tx, index int, msg *wasmtypes.MsgStoreCode) error {
	event, err := tx.FindEventByType(index, wasmtypes.EventTypeStoreCode)
	if err != nil {
		return err
	}

	codeID, err := tx.FindAttributeByKey(event, wasmtypes.AttributeKeyCodeID)
	if err != nil {
		return err
	}

	code := types.NewCode(codeID, msg.Sender, tx.Timestamp, tx.Height)

	return m.db.SaveCode(code)
}

func (m *Module) handleMsgInstantiateContract(tx *juno.Tx, index int, msg *wasmtypes.MsgInstantiateContract) error {
	event, err := tx.FindEventByType(index, wasmtypes.EventTypeInstantiate)
	if err != nil {
		return err
	}

	contractAddress, err := tx.FindAttributeByKey(event, wasmtypes.AttributeKeyContractAddr)
	if err != nil {
		return err
	}

	createdAt := &wasmtypes.AbsoluteTxPosition{
		BlockHeight: uint64(tx.Height),
		TxIndex:     uint64(index),
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Sender)
	admin, _ := sdk.AccAddressFromBech32(msg.Admin)
	contractInfo := wasmtypes.NewContractInfo(msg.CodeID, creator, admin, msg.Label, createdAt)
	contract := types.NewContract(&contractInfo, contractAddress, tx.Timestamp)

	return m.db.SaveContract(contract)
}
