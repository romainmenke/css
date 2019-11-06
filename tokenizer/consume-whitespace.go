package tokenizer

import "io"

func ConsumeWhiteSpace(t *Tokenizer) Token {
	for {
		peeked, _, err := t.ReadRune()
		if err == io.EOF {
			return TokenWhitespace{
				represenation: t.representation,
			}
		}
		if err != nil {
			return TokenError{error: err}
		}

		if peeked != '\u000a' && peeked != '\u0009' && peeked != '\u0020' {
			err := t.UnreadRune()
			if err != nil {
				return TokenError{error: err}
			}
			break
		}
	}

	return TokenWhitespace{
		represenation: t.representation,
	}
}