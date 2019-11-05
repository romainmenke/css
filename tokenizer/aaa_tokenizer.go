package tokenizer

import (
	"bufio"
	"io"
	"unicode"
)

type Tokenizer struct {
	b        *bufio.Reader
	tracking []rune
}

func New(r io.Reader) *Tokenizer {
	return &Tokenizer{
		b:        bufio.NewReader(r),
		tracking: make([]rune, 0, 1000),
	}
}

func (t *Tokenizer) Next() Token {
	for {
		r, _, err := t.b.ReadRune()
		if err == io.EOF {
			return TokenEOF{}
		}
		if err != nil {
			return TokenError{error: err}
		}

		// substitutions
		switch r {
		//Replace any U+000D CARRIAGE RETURN (CR) code points or pairs of U+000D CARRIAGE RETURN (CR) followed by U+000A LINE FEED (LF), by a single U+000A LINE FEED (LF) code point.
		case '\u000d':
			r = '\u000a'

			peeked, _, err := t.b.ReadRune()
			if err != io.EOF {
				if err != nil {
					return TokenError{error: err}
				}

				if peeked != '\u000a' {
					err := t.b.UnreadRune()
					if err != nil {
						return TokenError{error: err}
					}
				}
			}

			//Replace any U+000C FORM FEED (FF) code points by a single U+000A LINE FEED (LF) code point.
		case '\u000c':
			r = '\u000a'
			// Replace any U+0000 NULL or surrogate code points with U+FFFD REPLACEMENT CHARACTER (�).
		case '\u0000':
			r = '\ufffd'
		default:
			// Replace any U+0000 NULL or surrogate code points with U+FFFD REPLACEMENT CHARACTER (�).
			if unicode.In(r, unicode.Cs) {
				r = '\ufffd'
			}
		}

		switch r {

		// String
		case '\'', '"':
			return TokenizeString(t, r)

		// Whitespace
		case '\n', '\f', ' ', '\t':
			return TokenWhitespace{}

		case '\r':
			return TokenizeWhitespace(t)

		default:
			continue
		}
	}
}
