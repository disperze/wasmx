package types

// Code type
type Code struct {
	CodeID      uint64
	Creator     string
	CreatedTime string
	Height      int64
}

// NewCode instance
func NewCode(
	codeID uint64,
	creator string,
	createdTime string,
	height int64,
) Code {
	return Code{codeID, creator, createdTime, height}
}
