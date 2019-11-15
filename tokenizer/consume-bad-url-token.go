package tokenizer

import (
	"io"
)

func consumeBadURLToken(t *Tokenizer) Token {
	for {
		r, _, err := t.ReadRune()
		if err == io.EOF {
			return TokenBadUrl{}
		}
		if err != nil {
			return TokenError{error: err}
		}

		// Tokenize
		switch r {

		case ')':
			return TokenBadUrl{}

		case '\\':
			_, err := Unescape(t, r)
			if err != nil {
				// do nothin
			}

		default:
			// do nothin
		}
	}
}
