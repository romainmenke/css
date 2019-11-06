package tokenizer

import (
	"io"
	"unicode"
)

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
	first, _, err := t.ReadRune()
	if err == io.EOF {
		return TokenEOF{}
	}
	if err != nil {
		return TokenError{error: err}
	}

	switch {
	case CheckIfTwoCodePointsAreAValidEscape(t, first) || unicode.In(first, NameCodePoint...):
		token := TokenHash{}
		if CheckIfThreeCodePointsWouldStartAnIdentifier(t, first) {
			token.Type = HashTokenTypeID
		}

		err = t.UnreadRune()
		if err != nil {
			return TokenError{error: err}
		}

		name, err := ConsumeName(t)
		if err != nil {
			return TokenError{error: err}
		}

		token.Value = name
		token.represenation = t.representation
		return token
	default:
		return TokenDelim{
			Value:         first,
			represenation: t.representation,
		}
	}
}
