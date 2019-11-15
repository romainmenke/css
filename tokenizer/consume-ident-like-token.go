package tokenizer

import (
	"io"
	"strings"
)

func consumeIdentLikeToken(t *Tokenizer) Token {
	name, err := consumeName(t)
	if err != nil {
		return TokenError{error: err}
	}

	p1, err := t.peekOneRune()
	if err == io.EOF {
		return TokenIdent{
			Value:          name,
			representation: t.representation(),
		}
	}
	if err != nil {
		return TokenError{error: err}
	}

	if strings.ToLower(string(name)) == "url" && p1 == '(' {
		_, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		consumeWhiteSpace(t, 2)

		p2, err := t.peekOneRune()
		if err == io.EOF {
			return TokenIdent{
				Value:          name,
				representation: t.representation(),
			}
		}
		if err != nil {
			return TokenError{error: err}
		}

		if p2 == '"' || p2 == '\'' {
			_, _, err := t.ReadRune()
			if err != nil {
				return TokenError{error: err}
			}

			return TokenFunction{
				Value:          name,
				representation: t.representation(),
			}
		}

		return consumeURLToken(t)
	}

	if p1 == '(' {
		_, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		return TokenFunction{
			Value:          name,
			representation: t.representation(),
		}
	}

	return TokenIdent{
		Value:          name,
		representation: t.representation(),
	}
}
