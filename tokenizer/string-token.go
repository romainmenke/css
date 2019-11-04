package tokenizer

import (
	"errors"
	"io"
)

type StringToken struct {
	Value []byte
	Quote QuoteKind
}

func (t StringToken) String() string {
	if t.Quote == SingleQuote {
		return `'` + string(t.Value) + `'`
	}

	return `"` + string(t.Value) + `"`
}

type QuoteKind int

const SingleQuote QuoteKind = 0
const DoubleQuote QuoteKind = 1

func TokenizeString(t *Tokenizer, currentQuoteToken byte) Token {
	quoteKind := SingleQuote
	if currentQuoteToken == '"' {
		quoteKind = DoubleQuote
	}

	for {
		b2, err := t.b.ReadByte()
		if err != nil {
			return ErrorToken{error: err}
		}

		switch b2 {
		case currentQuoteToken:
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

			if peeked[0] == currentQuoteToken {
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
}
