package tokenizer

import (
	"io"
	"unicode"
)

func CheckIfTwoCodePointsAreAValidEscape(reader RuneReader, first rune) bool {
	if first != '\u005c' { // "\"
		return false
	}

	second, _, err := reader.ReadRune()
	if err != nil {
		return false
	}

	defer reader.UnreadRune()

	if second == '\u000a' {
		return false
	}

	return true
}

func CheckIfThreeCodePointsWouldStartAnIdentifier(reader RuneReader, first rune) bool {
	switch {
	case first == '\u002d': // "-"
		second, _, err := reader.ReadRune()
		if err != nil {
			return false
		}

		defer reader.UnreadRune()

		switch {
		case second == '\u002d': // "-"
			return true
		case unicode.In(second, NameStartCodePoint...):
			return true
		case CheckIfTwoCodePointsAreAValidEscape(reader, second):
			return true
		default:
			return false
		}

	case first == '\\': // "\"
		return CheckIfTwoCodePointsAreAValidEscape(reader, first)
	default:
		return false
	}
}

func CheckIfThreeCodePointsWouldStartANumber(reader RuneReader, first rune) bool {
	switch {
	case first == '\u002b' || first == '\u002d': // "+" or "-"
		second, _, err := reader.ReadRune()
		if err != nil {
			return false
		}

		defer reader.UnreadRune()

		switch {
		case unicode.In(second, unicode.Digit):
			return true
		case second == '\u002e': // "."
			return CheckIfThreeCodePointsWouldStartANumber(reader, second)
		default:
			return false
		}
	case first == '\u002e': // "."
		second, _, err := reader.ReadRune()
		if err != nil {
			return false
		}

		defer reader.UnreadRune()

		if unicode.In(second, unicode.Digit) {
			return true
		}

		return false
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
