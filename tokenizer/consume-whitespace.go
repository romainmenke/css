package tokenizer

import "io"

func consumeWhiteSpace(t *Tokenizer, max int) Token {
	current := 0

	for {
		if max != -1 && current > 0 && current == max {
			return TokenWhitespace{
				representation: append([]rune(nil), t.representation()...),
			}
		}

		current++

		peeked, err := t.peekOneRune()
		if err == io.EOF {
			return TokenWhitespace{
				representation: append([]rune(nil), t.representation()...),
			}
		}
		if err != nil {
			return TokenError{error: err}
		}

		if peeked != '\u000a' && peeked != '\u0009' && peeked != '\u0020' {
			break
		}

		_, _, err = t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}
	}

	return TokenWhitespace{
		representation: append([]rune(nil), t.representation()...),
	}
}
