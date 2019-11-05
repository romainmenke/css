package tokenizer

import (
	"errors"
)

type TokenString struct {
	Value []rune
	Quote QuoteKind
}

func (t TokenString) String() string {
	return string(t.Value)
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
			unescapedR, err := Unescape(t.b, r)
			if err != nil {
				return TokenError{error: err}
			}

			t.tracking = append(t.tracking, unescapedR)

		default:
			t.tracking = append(t.tracking, r)
		}
	}
}
