package tokenizer

import (
	"io"
	"unicode"
)

func CheckIfTwoCodePointsAreAValidEscape(reader RuneReader) bool {
	first, second, err := reader.PeekTwoRunes()
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

func CheckIfThreeCodePointsWouldStartAnIdentifier(reader RuneReader) bool {
	first, second, third, err := reader.PeekThreeRunes()
	if err != nil {
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

	case first == '\u005c': // "\"
		return second != '\u000a'
	default:
		return false
	}
}

func CheckIfThreeCodePointsWouldStartANumber(reader RuneReader) bool {
	first, second, third, err := reader.PeekThreeRunes()
	if err != nil {
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

func CheckIfNextIsEOF(reader RuneReader) bool {
	_, _, err := reader.ReadRune()
	if err == io.EOF {
		return true
	}
	if err != nil {
		return false // should not happen
	}

	reader.UnreadRune()

	return false
}

func CheckIfFirstCodePointIsInRangeTable(reader RuneReader, rt ...*unicode.RangeTable) bool {
	first, err := reader.PeekOneRune()
	if err != nil {
		return false
	}

	return unicode.In(first, rt...)
}
