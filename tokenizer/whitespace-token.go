package tokenizer

import "io"

type WhitespaceToken struct{}

func (t WhitespaceToken) String() string {
	return ""
}

func TokenizeWhitespace(t *Tokenizer) Token {
	peeked, err := t.b.Peek(1)
	if err == io.EOF {
		return WhitespaceToken{}
	}
	if err != nil {
		return ErrorToken{error: err}
	}

	if peeked[0] == '\n' {
		_, err := t.b.ReadByte()
		if err != nil {
			panic(err) // already succesfully peeked, no error should happen
		}
	}

	return WhitespaceToken{}
}
