package tokenizer

type HashToken struct {
	Type HashTokenType
}

// Default is "unrestricted"
type HashTokenType int

const HashTokenTypeUnrestricted HashTokenType = 0
const HashTokenTypeID HashTokenType = 1
