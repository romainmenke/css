package tokenizer

import (
	"encoding/hex"
	"errors"
	"io"
	"unicode"
	"unicode/utf8"
)

func Unescape(reader RuneReader, r rune) (rune, error) {
	isHex := false
	capturedHex := []rune{}

ESCAPE_HEX_PEEK:
	peeked, err := reader.peekOneRune()
	if err == io.EOF {
		if len(capturedHex) > 0 {
			return decodeHex(capturedHex)
		}

		return r, nil
	}
	if err != nil {
		return 0, err
	}

	if isHex &&
		((len(capturedHex) == 6) ||
			(peeked == '\n' ||
				peeked == '\f' ||
				peeked == ' ' ||
				peeked == '\t')) {
		_, _, err := reader.ReadRune()
		if err != nil {
			return 0, err
		}

		return decodeHex(capturedHex)
	}

	if unicode.In(peeked, unicode.Hex_Digit) { // Is Hex
		isHex = true

		rr, _, err := reader.ReadRune()
		if err != nil {
			return 0, err
		}

		capturedHex = append(capturedHex, rr)

		goto ESCAPE_HEX_PEEK

	} else if !isHex { // Not newline or hex digit
		switch peeked {
		case '\n', '\r', 'f': // Is newline
			return r, nil // ???
		default: // Is not newline or hex digit
			// unescaped thing
			_, _, err := reader.ReadRune()
			if err != nil {
				return 0, err
			}

			r = peeked
		}
	} else {
		return decodeHex(capturedHex)
	}

	return r, nil
}

func decodeHex(captured []rune) (rune, error) {
	b, err := hex.DecodeString(string(captured))
	if err != nil {
		return 0, err
	}

	unescaped, n := utf8.DecodeRune(b)
	if n == 0 {
		return 0, errors.New("invalid escape sequence")
	}

	return unescaped, err
}
