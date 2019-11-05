package tokenizer

import (
	"errors"
	"io"
)

type TokenString struct {
	Value []rune
	Quote QuoteKind
}

func (t TokenString) String() string {
	if t.Quote == SingleQuote {
		return `'` + string(t.Value) + `'`
	}

	return `"` + string(t.Value) + `"`
}

type QuoteKind int

const SingleQuote QuoteKind = 0
const DoubleQuote QuoteKind = 1

func TokenizeString(t *Tokenizer, currentQuoteToken rune) Token {
	quoteKind := SingleQuote
	if currentQuoteToken == '"' {
		quoteKind = DoubleQuote
	}

	for {
		r, _, err := t.b.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		switch r {
		case currentQuoteToken:
			return TokenString{
				Value: t.tracking,
				Quote: quoteKind,
			}

		case '\n', '\r', '\f':
			return TokenError{error: errors.New("unexpected newline")}

		case '\\':
			t.tracking = append(t.tracking, r)

			peeked, _, err := t.b.ReadRune()
			if err == io.EOF {
				return TokenString{
					Value: t.tracking,
					Quote: quoteKind,
				}
			}
			if err != nil {
				return TokenError{error: err}
			}

			if peeked == currentQuoteToken {
				t.tracking = append(t.tracking, peeked)
			} else {
				err := t.b.UnreadRune()
				if err != nil {
					return TokenError{error: err}
				}
			}

		default:
			t.tracking = append(t.tracking, r)
		}
	}
}
