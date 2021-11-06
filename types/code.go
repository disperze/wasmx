package types

// Code type
type Code struct {
	CodeID      string
	Creator     string
	CreatedTime string
	Height      int64
}

// NewCode instance
func NewCode(
	codeID string,
	creator string,
	createdTime string,
	height int64,
) Code {
	return Code{codeID, creator, createdTime, height}
}
