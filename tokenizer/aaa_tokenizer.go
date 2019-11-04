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
			quoteKind := SingleQuote
			if b == '"' {
				quoteKind = DoubleQuote
			}

			for {
				b2, err := t.b.ReadByte()
				if err != nil {
					return ErrorToken{error: err}
				}

				switch b2 {
				case b:
					return StringToken{
						Value: t.tracking,
						Quote: quoteKind,
					}

				case '\n', '\r', '\f':
					return ErrorToken{error: errors.New("unexpected newline")}

				case '\\':

					peeked, err := t.b.Peek(1)
					if err == io.EOF {
						return StringToken{
							Value: t.tracking,
							Quote: quoteKind,
						}
					}
					if err != nil {
						return ErrorToken{error: err}
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
			return WhitespaceToken{}

		case '\r':
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

		default:
			continue
		}
	}
}
