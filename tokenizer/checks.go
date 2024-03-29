package tokenizer

import (
	"io"
	"unicode"
)

func checkIfTwoCodePointsAreAValidEscape(reader RuneReader) bool {
	first, second, err := reader.peekTwoRunes()
	if err != nil {
		return false
	}

	if first != '\u005c' { // "\"
		return false
	}

	if second == '\u000a' {
		return false
	}

	return true
}

func checkIfThreeCodePointsWouldStartAnIdentifier(reader RuneReader) bool {
	first, second, third, err := reader.peekThreeRunes()
	if err != nil && err != io.EOF {
		return false
	}

	switch {
	case first == '\u002d': // "-"
		switch {
		case second == '\u002d': // "-"
			return true
		case unicode.In(second, NameStartCodePoint...):
			return true
		case second != '\u005c':
			return third != '\u000a'
		default:
			return false
		}

	case unicode.In(first, NameStartCodePoint...):
		return true

	case first == '\u005c': // "\"
		return second != '\u000a'
	default:
		return false
	}
}

func checkIfThreeCodePointsWouldStartANumber(reader RuneReader) bool {
	first, second, third, err := reader.peekThreeRunes()
	if err != nil && err != io.EOF {
		return false
	}

	switch {
	case first == '\u002b' || first == '\u002d': // "+" or "-"
		switch {
		case unicode.In(second, unicode.Digit):
			return true
		case second == '\u002e': // "."
			return unicode.In(third, unicode.Digit)
		default:
			return false
		}
	case first == '\u002e': // "."
		return unicode.In(second, unicode.Digit)
	case unicode.In(first, unicode.Digit):
		return true
	default:
		return false
	}
}

func checkIfNextIsEOF(reader RuneReader) bool {
	_, err := reader.peekOneRune()
	if err == io.EOF {
		return true
	}
	if err != nil {
		return false // should not happen
	}

	return false
}

func checkIfFirstCodePointIsInRangeTable(reader RuneReader, rt ...*unicode.RangeTable) bool {
	first, err := reader.peekOneRune()
	if err != nil {
		return false
	}

	return unicode.In(first, rt...)
}
