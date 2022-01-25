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

// CodeData type
type CodeData struct {
	CodeID  uint64
	Version *string
	IBC     *bool
	CW20    *bool
}

// NewCodeData instance
func NewCodeData(
	codeID uint64,
	version string,
	ibc bool,
	cw20 bool,
) CodeData {
	return CodeData{codeID, &version, &ibc, &cw20}
}
