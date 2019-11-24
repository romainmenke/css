package tokenizer

import (
	"io"
)

func consumeURLToken(t *Tokenizer) Token {
	consumeWhiteSpace(t, -1)
	t.tracking = t.tracking[:0]

	for {
		r, _, err := t.ReadRune()
		if err == io.EOF {
			return TokenUrl{}
		}
		if err != nil {
			return TokenError{error: err}
		}

		// Tokenize
		switch r {

		case '(', '\'', '"':
			return consumeBadURLToken(t)

		case '\\':
			unescapedR, err := Unescape(t, r)
			if err != nil {
				return consumeBadURLToken(t)
			}

			t.tracking = append(t.tracking, unescapedR)

		case ')': // Right Parenthesis
			return TokenUrl{
				Value:          append([]rune(nil), t.tracking...),
				representation: append([]rune(nil), t.representation()...),
			}

		case '\u000a', '\u0009', '\u0020': // Whitespace
			consumeWhiteSpace(t, -1)

		default:
			t.tracking = append(t.tracking, r)
		}
	}
}
