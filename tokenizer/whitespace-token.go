package tokenizer

import (
	"io"
)

type TokenWhitespace struct{}

func (t TokenWhitespace) String() string {
	return ""
}

func TokenizeWhitespace(t *Tokenizer) Token {
	peeked, _, err := t.b.ReadRune()
	if err == io.EOF {
		return TokenWhitespace{}
	}
	if err != nil {
		return TokenError{error: err}
	}

	if peeked != '\n' {
		err := t.b.UnreadRune()
		if err != nil {
			return TokenError{error: err}
		}
	}

	return TokenWhitespace{}
}
