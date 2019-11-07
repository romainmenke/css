package tokenizer

import "io"

func ConsumeWhiteSpace(t *Tokenizer) Token {
	for {
		peeked, err := t.PeekOneRune()
		if err == io.EOF {
			return TokenWhitespace{
				represenation: t.Representation(),
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
		represenation: t.Representation(),
	}
}
