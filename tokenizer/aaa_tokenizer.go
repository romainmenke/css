package tokenizer

import (
	"bufio"
	"io"
	"unicode"
)

type RuneReader interface {
	ReadRune() (rune, int, error)
	UnreadRune() error
}

type Tokenizer struct {
	b              *bufio.Reader
	tracking       []rune
	representation []rune
}

func New(r io.Reader) *Tokenizer {
	return &Tokenizer{
		b:              bufio.NewReader(r),
		tracking:       make([]rune, 0, 1000),
		representation: make([]rune, 0, 1000),
	}
}

func (t *Tokenizer) ReadRune() (rune, int, error) {
	r, size, err := t.b.ReadRune()
	if size == 0 {
		return r, size, err
	}

	t.representation = append(t.representation, r)

	// SUBSTITUTIONS
	switch r {
	case '\u000d': //Replace any U+000D CARRIAGE RETURN (CR) code points or pairs of U+000D CARRIAGE RETURN (CR) followed by U+000A LINE FEED (LF), by a single U+000A LINE FEED (LF) code point.
		r = '\u000a'

		peeked, _, err := t.b.ReadRune()
		if err != io.EOF {
			if err != nil {
				return r, size, err
			}

			if peeked != '\u000a' {
				err := t.b.UnreadRune()
				if err != nil {
					return r, size, err
				}
			}

			t.representation = append(t.representation, peeked)
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

	return r, size, err
}

func (t *Tokenizer) UnreadRune() error {
	err := t.b.UnreadRune()
	if err != nil {
		return err
	}

	t.representation = t.representation[:len(t.representation)-1]

	return nil
}

func (t *Tokenizer) Next() Token {
	t.tracking = t.tracking[:0]
	t.representation = t.representation[:0]

	for {
		r, _, err := t.ReadRune()
		if err == io.EOF {
			return TokenEOF{}
		}
		if err != nil {
			return TokenError{error: err}
		}

		// Tokenize
		switch r {

		case '\'', '"': // String
			return TokenizeString(t, r)

		// TODO : collapse continous whitespace into 1 token
		case '\u000a', '\u0009', '\u0020': // Whitespace
			return TokenizeWhitespace(t)

		default:
			continue
		}
	}
}
