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

func (t *Tokenizer) Next() (Token, error) {
	for {
		b, err := t.b.ReadByte()
		if err != nil {
			return nil, err
		}

		switch b {

		case '\n', '\f', ' ', '\t':
			return WhitespaceToken{}, nil

		case '\r':
			peeked, err := t.b.Peek(1)
			if err == io.EOF {
				return WhitespaceToken{}, nil
			}
			if err != nil {
				return nil, err
			}

			if peeked[0] == '\n' {
				_, err := t.b.ReadByte()
				if err != nil {
					panic(err) // already succesfully peeked, no error should happen
				}
			}

			return WhitespaceToken{}, nil

		default:
			continue
		}
	}
}
