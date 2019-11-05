package tokenizer

import (
	"io"
)

type TokenWhitespace struct {
	represenation []rune
}

func (t TokenWhitespace) String() string {
	return ""
}

func (t TokenWhitespace) Representation() string {
	return string(t.represenation)
}

func TokenizeWhitespace(t *Tokenizer) Token {
	peeked, _, err := t.ReadRune()
	if err == io.EOF {
		return TokenWhitespace{
			represenation: t.representation,
		}
	}
	if err != nil {
		return TokenError{error: err}
	}

	if peeked != '\u000a' {
		err := t.UnreadRune()
		if err != nil {
			return TokenError{error: err}
		}
	}

	return TokenWhitespace{
		represenation: t.representation,
	}
}
