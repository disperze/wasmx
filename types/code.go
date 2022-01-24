package types

// Code type
type Code struct {
	CodeID      uint64
	Creator     string
	Hash        string
	Size        uint64
	CreatedTime string
	Height      int64
}

// NewCode instance
func NewCode(
	codeID uint64,
	creator string,
	hash string,
	size uint64,
	createdTime string,
	height int64,
) Code {
	return Code{codeID, creator, hash, size, createdTime, height}
}
