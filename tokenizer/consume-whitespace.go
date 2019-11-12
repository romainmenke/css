package tokenizer

import "io"

func ConsumeWhiteSpace(t *Tokenizer, max int) Token {
	current := 0

	for {
		if max != -1 && current > 0 && current == max {
			return TokenWhitespace{
				representation: t.Representation(),
			}
		}

		current++

		peeked, err := t.PeekOneRune()
		if err == io.EOF {
			return TokenWhitespace{
				representation: t.Representation(),
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
		representation: t.Representation(),
	}
}
