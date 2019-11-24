package tokenizer

type TokenHash struct {
	Type           HashTokenType
	Value          []rune
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
	case checkIfFirstCodePointIsInRangeTable(t, NameCodePoint...) || checkIfTwoCodePointsAreAValidEscape(t):
		token := TokenHash{}
		if checkIfThreeCodePointsWouldStartAnIdentifier(t) {
			token.Type = HashTokenTypeID
		}

		name, err := consumeName(t)
		if err != nil {
			return TokenError{error: err}
		}

		token.Value = append([]rune(nil), name...)
		token.representation = append([]rune(nil), t.representation()...)
		return token
	default:
		r, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		return TokenDelim{
			Value:          r,
			representation: append([]rune(nil), t.representation()...),
		}
	}
}
