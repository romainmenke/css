package streampreprocessor

import (
	"bufio"
	"io"
	"unicode"
)

type PreProcessor struct {
	reader *bufio.Reader
}

func New(r *bufio.Reader) *PreProcessor {
	return &PreProcessor{
		reader: r,
	}
}

func (p *PreProcessor) ReadRune() (rune, int, error) {
	r, size, err := p.reader.ReadRune()
	if size == 0 {
		return r, size, err
	}

	r, err = p.substitute(r)
	if err != nil {
		return 0, 0, err
	}

	return r, size, err
}

func (p *PreProcessor) substitute(r rune) (rune, error) {
	// SUBSTITUTIONS
	switch r {
	case '\u000d': //Replace any U+000D CARRIAGE RETURN (CR) code points or pairs of U+000D CARRIAGE RETURN (CR) followed by U+000A LINE FEED (LF), by a single U+000A LINE FEED (LF) code point.
		r = '\u000a'

		peeked, _, err := p.reader.ReadRune()
		if err == io.EOF {
			return r, nil
		}
		if err != nil {
			return 0, err
		}

		if peeked == '\u000a' {
			return r, err
		}

		p.reader.UnreadRune()

	case '\u000c': //Replace any U+000C FORM FEED (FF) code points by a single U+000A LINE FEED (LF) code point.
		r = '\u000a'

	case '\u0000': // Replace any U+0000 NULL or surrogate code points with U+FFFD REPLACEMENT CHARACTER (�).
		r = '\ufffd'
	default:
		if unicode.In(r, unicode.Cs) { // Replace any U+0000 NULL or surrogate code points with U+FFFD REPLACEMENT CHARACTER (�).
			r = '\ufffd'
		}
	}

	return r, nil
}
