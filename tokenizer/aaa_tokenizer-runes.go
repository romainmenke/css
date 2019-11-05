package tokenizer

import (
	"io"
	"unicode"
)

type RuneReader interface {
	ReadRune() (rune, int, error)
	UnreadRune() error
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

			t.representation = append(t.representation, peeked)
			if peeked != '\u000a' {
				err := t.UnreadRune()
				if err != nil {
					return r, size, err
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
