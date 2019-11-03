package tokenizer

import (
	"bufio"
	"errors"
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

		// String
		case '\'', '"':
			quoteKind := SingleQuote
			if b == '"' {
				quoteKind = DoubleQuote
			}

			for {
				b2, err := t.b.ReadByte()
				if err != nil {
					return nil, err
				}

				switch b2 {
				case b:
					return StringToken{
						Value: t.tracking,
						Quote: quoteKind,
					}, nil

				case '\n', '\r', '\f':
					return nil, errors.New("unexpected newline")

				case '\\':

					peeked, err := t.b.Peek(1)
					if err == io.EOF {
						return StringToken{
							Value: t.tracking,
							Quote: quoteKind,
						}, nil
					}
					if err != nil {
						return nil, err
					}

					t.tracking = append(t.tracking, b2)

					if peeked[0] == b {
						b3, err := t.b.ReadByte()
						if err != nil {
							panic(err) // already succesfully peeked, no error should happen
						}

						t.tracking = append(t.tracking, b3)
					}

				default:
					t.tracking = append(t.tracking, b2)
				}
			}

		// Whitespace
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
