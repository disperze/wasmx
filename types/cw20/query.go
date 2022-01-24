package cw20

// TokenInfo type
type TokenInfo struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    uint8  `json:"decimals"`
	TotalSupply string `json:"total_supply"`
}
