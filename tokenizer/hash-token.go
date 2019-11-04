package tokenizer

type TokenHash struct {
	Type HashTokenType
}

func (t TokenHash) String() string {
	return ""
}

// Default is "unrestricted"
type HashTokenType int

const HashTokenTypeUnrestricted HashTokenType = 0
const HashTokenTypeID HashTokenType = 1
