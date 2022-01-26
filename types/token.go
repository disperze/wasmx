package types

// Token type
type Token struct {
	Contract     string
	Name         string
	Symbol       string
	Decimals     uint8
	InitalSupply string
}

// NewToken instance
func NewToken(
	contract string,
	name string,
	symbol string,
	decimals uint8,
	initalSupply string,
) Token {
	return Token{contract, name, symbol, decimals, initalSupply}
}
