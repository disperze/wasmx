package wasm

import (
	"context"
	"fmt"

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
	contracts, err := GetAllContracts(tx, index, wasmtypes.EventTypeInstantiate)
	if err != nil {
		return err
	}

	if len(contracts) == 0 {
		return fmt.Errorf("No contract address found")
	}

	createdAt := &wasmtypes.AbsoluteTxPosition{
		BlockHeight: uint64(tx.Height),
		TxIndex:     uint64(index),
	}
	ctx := context.Background()
	for _, contractAddress := range contracts {
		response, err := m.client.ContractInfo(ctx, &wasmtypes.QueryContractInfoRequest{
			Address: contractAddress,
		})
		if err != nil {
			return err
		}

		creator, _ := sdk.AccAddressFromBech32(response.Creator)
		var admin sdk.AccAddress
		if response.Admin != "" {
			admin, _ = sdk.AccAddressFromBech32(response.Admin)
		}

		contractInfo := wasmtypes.NewContractInfo(response.CodeID, creator, admin, response.Label, createdAt)
		contract := types.NewContract(&contractInfo, contractAddress, tx.Timestamp)

		if err = m.db.SaveContract(contract); err != nil {
			return err
		}
	}

	return nil
}

func GetAllContracts(tx *juno.Tx, index int, eventType string) ([]string, error) {
	contracts := []string{}
	event, err := tx.FindEventByType(index, eventType)
	if err != nil {
		return contracts, err
	}

	for _, attr := range event.Attributes {
		if attr.Key == wasmtypes.AttributeKeyContractAddr {
			contracts = append(contracts, attr.Value)
		}
	}

	return contracts, nil
}
