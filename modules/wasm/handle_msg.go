package wasm

import (
	"github.com/disperze/wasmx/database"
	"github.com/disperze/wasmx/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/desmos-labs/juno/types"
)

// Listen contracts by code
const CodeID = 3

// HandleMsg allows to handle the different utils related to the gov module
func HandleMsg(
	tx *juno.Tx, index int, msg sdk.Msg,
	wasmClient wasmtypes.QueryClient,
	cdc codec.Marshaler, db *database.Db,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *wasmtypes.MsgInstantiateContract:
		return handleMsgInstantiateContract(tx, index, cosmosMsg, db)
	}

	return nil
}

func handleMsgInstantiateContract(tx *juno.Tx, index int, msg *wasmtypes.MsgInstantiateContract, db *database.Db) error {
	event, err := tx.FindEventByType(index, sdk.EventTypeMessage)
	if err != nil {
		return err
	}

	contractAddress, err := tx.FindAttributeByKey(event, wasmtypes.AttributeKeyContract)
	if err != nil {
		return err
	}

	createdAt := &wasmtypes.AbsoluteTxPosition{
		BlockHeight: uint64(tx.Height),
		TxIndex:     uint64(index),
	}

	contractInfo := wasmtypes.NewContractInfo(msg.CodeID, sdk.AccAddress(msg.Sender), sdk.AccAddress(msg.Admin), msg.Label, createdAt)
	contract := types.NewContract(&contractInfo, contractAddress, tx.Timestamp)

	return db.SaveContract(contract)
}
