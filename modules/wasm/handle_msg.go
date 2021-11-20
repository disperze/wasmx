package wasm

import (
	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v2/types"
)

// HandleMsg allows to handle the different utils related to the gov module
func HandleMsg(
	tx *juno.Tx, index int, msg sdk.Msg, db *database.Db,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *wasmtypes.MsgStoreCode:
		return handleMsgStoreCode(tx, index, cosmosMsg, db)
	case *wasmtypes.MsgInstantiateContract:
		return handleMsgInstantiateContract(tx, index, cosmosMsg, db)
	}

	return nil
}

func handleMsgStoreCode(tx *juno.Tx, index int, msg *wasmtypes.MsgStoreCode, db *database.Db) error {
	event, err := tx.FindEventByType(index, wasmtypes.EventTypeStoreCode)
	if err != nil {
		return err
	}

	codeID, err := tx.FindAttributeByKey(event, wasmtypes.AttributeKeyCodeID)
	if err != nil {
		return err
	}

	code := types.NewCode(codeID, msg.Sender, tx.Timestamp, tx.Height)

	return db.SaveCode(code)
}

func handleMsgInstantiateContract(tx *juno.Tx, index int, msg *wasmtypes.MsgInstantiateContract, db *database.Db) error {
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

	return db.SaveContract(contract)
}
