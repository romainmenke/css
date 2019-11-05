package tokenizer

import (
	"bufio"
	"encoding/hex"
	"errors"
	"io"
	"unicode"
	"unicode/utf8"
)

func Unescape(b *bufio.Reader, r rune) (rune, error) {
	isHex := false
	capturedHex := []rune{}

ESCAPE_HEX_PEEK:
	peeked, _, err := b.ReadRune()
	if err == io.EOF {
		return decodeHex(capturedHex)
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
		return decodeHex(capturedHex)
	}

	if unicode.In(peeked, unicode.Hex_Digit) { // Is Hex
		isHex = true
		capturedHex = append(capturedHex, peeked)

		goto ESCAPE_HEX_PEEK

	} else if !isHex { // Not newline or hex digit
		switch peeked {
		case '\n', '\r', 'f': // Is newline
			err := b.UnreadByte()
			if err != nil {
				return 0, err
			}

			return r, nil // ???
		default: // Is not newline or hex digit
			// unescaped thing
			r = peeked
		}
	} else {
		return 0, errors.New("invalid escape sequence")
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
