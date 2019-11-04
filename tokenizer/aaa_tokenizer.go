package tokenizer

import (
	"bufio"
	"io"
)

type Tokenizer struct {
	b        *bufio.Reader
	tracking []byte
}

func New(r io.Reader) *Tokenizer {
	return &Tokenizer{
		b:        bufio.NewReader(r),
		tracking: make([]byte, 0, 1000),
	}
}

func (t *Tokenizer) Next() Token {
	for {
		b, err := t.b.ReadByte()
		if err == io.EOF {
			return EOFToken{}
		}
		if err != nil {
			return ErrorToken{error: err}
		}

		switch b {

		// String
		case '\'', '"':
			return TokenizeString(t, b)

		// Whitespace
		case '\n', '\f', ' ', '\t':
			return WhitespaceToken{}

		case '\r':
			return TokenizeWhitespace(t)

		default:
			continue
		}
	}
}
