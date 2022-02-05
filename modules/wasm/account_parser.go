package wasm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/forbole/juno/v2/modules/messages"
)

func WasmMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *wasmtypes.MsgStoreCode:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgInstantiateContract:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgExecuteContract:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgMigrateContract:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgUpdateAdmin:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgClearAdmin:
		return []string{msg.Sender}, nil
	}

	return nil, messages.MessageNotSupported(cosmosMsg)
}
