package tokenizer

import "io"

func consumeString(t *Tokenizer, currentQuoteToken rune) Token {
	quoteKind := SingleQuote
	if currentQuoteToken == '"' {
		quoteKind = DoubleQuote
	}

	for {
		r, _, err := t.ReadRune()
		if err != nil {
			if err == io.EOF {
				return TokenString{
					Value:          t.tracking,
					representation: t.representation(),
					Quote:          quoteKind,
				}
			}

			return TokenError{error: err}
		}

		switch r {
		case currentQuoteToken:
			return TokenString{
				Value:          t.tracking,
				representation: t.representation(),
				Quote:          quoteKind,
			}

		case '\n', '\r', '\f':
			return TokenBadString{}

		case '\\':
			unescapedR, err := Unescape(t, r)
			if err != nil {
				return TokenError{error: err}
			}

			if checkIfNextIsEOF(t) {
				return TokenString{
					Value:          t.tracking,
					representation: t.representation(),
					Quote:          quoteKind,
				}
			}

			t.tracking = append(t.tracking, unescapedR)

		default:
			t.tracking = append(t.tracking, r)
		}
	}
}
