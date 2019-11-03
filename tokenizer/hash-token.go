package tokenizer

type HashToken struct {
	Type HashTokenType
}

func (t HashToken) String() string {
	return ""
}

// Default is "unrestricted"
type HashTokenType int

const HashTokenTypeUnrestricted HashTokenType = 0
const HashTokenTypeID HashTokenType = 1
