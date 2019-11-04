package tokenizer

import "io"

type TokenWhitespace struct{}

func (t TokenWhitespace) String() string {
	return ""
}

func TokenizeWhitespace(t *Tokenizer) Token {
	peeked, err := t.b.Peek(1)
	if err == io.EOF {
		return TokenWhitespace{}
	}
	if err != nil {
		return TokenError{error: err}
	}

	if peeked[0] == '\n' {
		_, err := t.b.ReadByte()
		if err != nil {
			panic(err) // already succesfully peeked, no error should happen
		}
	}

	return TokenWhitespace{}
}
