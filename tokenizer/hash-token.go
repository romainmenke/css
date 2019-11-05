package tokenizer

type TokenHash struct {
	Type  HashTokenType
	Value []rune
}

func (t TokenHash) String() string {
	return string(t.Value)
}

// Default is "unrestricted"
type HashTokenType int

const HashTokenTypeUnrestricted HashTokenType = 0
const HashTokenTypeID HashTokenType = 1
