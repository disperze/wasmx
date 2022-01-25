package wasm

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/disperze/wasmx/types"
	"github.com/disperze/wasmx/types/cw20"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v2/types"
)

// HandleMsg allows to handle the different utils related to the gov module
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *wasmtypes.MsgStoreCode:
		return m.handleMsgStoreCode(tx, index, cosmosMsg)
	case *wasmtypes.MsgInstantiateContract:
		return m.handleMsgInstantiateContract(tx, index, cosmosMsg)
	case *wasmtypes.MsgExecuteContract:
		return m.handleMsgExecuteContract(tx, index, cosmosMsg)
	}

	return nil
}

func (m *Module) handleMsgStoreCode(tx *juno.Tx, index int, msg *wasmtypes.MsgStoreCode) error {
	event, err := tx.FindEventByType(index, wasmtypes.EventTypeStoreCode)
	if err != nil {
		return err
	}

	codeIDVal, err := tx.FindAttributeByKey(event, wasmtypes.AttributeKeyCodeID)
	if err != nil {
		return err
	}

	codeID, _ := strconv.Atoi(codeIDVal)
	response, err := m.client.Code(context.Background(), &wasmtypes.QueryCodeRequest{
		CodeId: uint64(codeID),
	})
	if err != nil {
		return err
	}

	hash := response.CodeInfoResponse.DataHash.String()
	size := len(response.Data)
	code := types.NewCode(uint64(codeID), msg.Sender, hash, uint64(size), tx.Timestamp, tx.Height)

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

	fee := tx.GetFee()
	feeAmount := int64(0)
	if fee.Len() == 1 {
		feeAmount = fee[0].Amount.Int64()
	}

	ctx := context.Background()
	for i, contractAddress := range contracts {
		response, err := m.client.ContractInfo(ctx, &wasmtypes.QueryContractInfoRequest{
			Address: contractAddress,
		})
		if err != nil {
			return err
		}

		creator, _ := sdk.AccAddressFromBech32(response.Creator)
		admin, _ := sdk.AccAddressFromBech32(response.Admin)
		contractInfo := wasmtypes.NewContractInfo(response.CodeID, creator, admin, response.Label, createdAt)
		contract := types.NewContract(&contractInfo, contractAddress, tx.Timestamp)

		if i == 0 {
			err = m.db.SaveContract(contract, tx.GasUsed, feeAmount)
		} else {
			err = m.db.SaveContract(contract, 0, 0)
		}

		if err != nil {
			return err
		}

		// Store code data
		data, err := m.db.GetCodeData(response.CodeID)
		if err != nil {
			return err
		}

		// Check if cw20 token
		var tokenInfo *cw20.TokenInfo
		if data.Version == nil || *data.CW20 {
			cw20Response, err := m.client.SmartContractState(ctx, &wasmtypes.QuerySmartContractStateRequest{
				Address:   contractAddress,
				QueryData: []byte(`{"token_info":{}}`),
			})
			if err == nil {
				var token cw20.TokenInfo
				err = json.Unmarshal(cw20Response.Data, &token)
				if err == nil {
					tokenInfo = &token
				}
			}
		}

		if data.Version == nil {
			res, err := m.client.RawContractState(ctx, &wasmtypes.QueryRawContractStateRequest{
				Address:   contractAddress,
				QueryData: []byte("contract_info"),
			})

			version := "none"
			if err == nil && res.Data != nil {
				version = string(res.Data)
			}

			isIBC := response.IBCPortID != ""
			isCW20 := tokenInfo != nil
			newData := types.NewCodeData(response.CodeID, &version, &isIBC, &isCW20)
			err = m.db.SetCodeData(newData)
			if err != nil {
				return err
			}
		}

		// Save cw20 token
		if tokenInfo != nil {
			token := types.NewToken(contractAddress, tokenInfo.Name, tokenInfo.Symbol, tokenInfo.Decimals, tokenInfo.TotalSupply)
			err = m.db.SaveToken(token)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Module) handleMsgExecuteContract(tx *juno.Tx, index int, msg *wasmtypes.MsgExecuteContract) error {
	contracts, err := GetAllContracts(tx, index, wasmtypes.EventTypeExecute)
	if err != nil {
		return err
	}

	if len(contracts) == 0 {
		return fmt.Errorf("No contract address found")
	}

	fee := tx.GetFee()
	feeAmount := int64(0)
	if fee.Len() == 1 {
		feeAmount = fee[0].Amount.Int64()
	}

	for i, contract := range contracts {
		if i == 0 {
			err = m.db.UpdateContractStats(contract, 1, tx.GasUsed, feeAmount)
		} else {
			err = m.db.UpdateContractStats(contract, 1, 0, 0)
		}

		if err != nil {
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
