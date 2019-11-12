package tokenizer

import (
	"io"
)

func ConsumeURLToken(t *Tokenizer) Token {
	ConsumeWhiteSpace(t, -1)
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
			return ConsumeBadURLToken(t)

		case '\\':
			unescapedR, err := Unescape(t, r)
			if err != nil {
				return ConsumeBadURLToken(t)
			}

			t.tracking = append(t.tracking, unescapedR)

		case ')': // Right Parenthesis
			return TokenUrl{
				Value:          t.tracking,
				representation: t.Representation(),
			}

		case '\u000a', '\u0009', '\u0020': // Whitespace
			ConsumeWhiteSpace(t, -1)

		default:
			t.tracking = append(t.tracking, r)
		}
	}
}
