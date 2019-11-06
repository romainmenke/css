package tokenizer

type TokenHash struct {
	Type          HashTokenType
	Value         []rune
	represenation []rune
}

func (t TokenHash) String() string {
	return string(t.Value)
}

func (t TokenHash) Representation() []rune {
	return t.represenation
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
		token.represenation = t.representation
		return token
	default:
		r, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		return TokenDelim{
			Value:         r,
			represenation: t.representation,
		}
	}
}
