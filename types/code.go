package types

// Code type
type Code struct {
	CodeID      string
	Source      string
	Builder     string
	Creator     string
	CreatedTime string
	Height      int64
}

// NewCode instance
func NewCode(
	codeID string,
	source string,
	builder string,
	creator string,
	createdTime string,
	height int64,
) Code {
	return Code{codeID, source, builder, creator, createdTime, height}
}
