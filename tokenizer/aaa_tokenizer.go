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

		// SUBSTITUTIONS
		switch r {
		case '\u000d': //Replace any U+000D CARRIAGE RETURN (CR) code points or pairs of U+000D CARRIAGE RETURN (CR) followed by U+000A LINE FEED (LF), by a single U+000A LINE FEED (LF) code point.
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

		case '\u000c': //Replace any U+000C FORM FEED (FF) code points by a single U+000A LINE FEED (LF) code point.
			r = '\u000a'

		case '\u0000': // Replace any U+0000 NULL or surrogate code points with U+FFFD REPLACEMENT CHARACTER (�).
			r = '\ufffd'
		default:
			if unicode.In(r, unicode.Cs) { // Replace any U+0000 NULL or surrogate code points with U+FFFD REPLACEMENT CHARACTER (�).
				r = '\ufffd'
			}
		}

		// ESCAPE
		if r == '\\' {
			unescapedR, err := Unescape(t.b, r)
			if err != nil {
				return TokenError{error: err}
			}

			r = unescapedR
		}

		// Tokenize
		switch r {

		case '\'', '"': // String
			return TokenizeString(t, r)

		// TODO : collapse continous whitespace into 1 token
		case '\n', '\f', ' ', '\t': // Whitespace
			return TokenWhitespace{}

		case '\r': // Whitespace
			return TokenizeWhitespace(t)

		default:
			continue
		}
	}
}
