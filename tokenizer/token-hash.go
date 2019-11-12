package tokenizer

type TokenHash struct {
	Type          HashTokenType
	Value         []rune
	representation []rune
}

func (t TokenHash) String() string {
	return string(t.Value)
}

func (t TokenHash) Representation() []rune {
	return t.representation
}

// Default is "unrestricted"
type HashTokenType int

const HashTokenTypeUnrestricted HashTokenType = 0
const HashTokenTypeID HashTokenType = 1

func TokenizeHashFromNumberSign(t *Tokenizer) Token {
	switch {
	case CheckIfFirstCodePointIsInRangeTable(t, NameCodePoint...) || CheckIfTwoCodePointsAreAValidEscape(t):
		token := TokenHash{}
		if CheckIfThreeCodePointsWouldStartAnIdentifier(t) {
			token.Type = HashTokenTypeID
		}

		name, err := ConsumeName(t)
		if err != nil {
			return TokenError{error: err}
		}

		token.Value = name
		token.representation = t.Representation()
		return token
	default:
		r, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		return TokenDelim{
			Value:         r,
			representation: t.Representation(),
		}
	}
}
