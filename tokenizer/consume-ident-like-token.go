package tokenizer

import (
	"io"
	"strings"
)

func ConsumeIdentLikeToken(t *Tokenizer) Token {
	name, err := ConsumeName(t)
	if err != nil {
		return TokenError{error: err}
	}

	p1, err := t.PeekOneRune()
	if err == io.EOF {
		return TokenIdent{
			Value:          name,
			representation: t.Representation(),
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

		ConsumeWhiteSpace(t, 2)

		p2, err := t.PeekOneRune()
		if err == io.EOF {
			return TokenIdent{
				Value:          name,
				representation: t.Representation(),
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
				representation: t.Representation(),
			}
		}

		return ConsumeURLToken(t)
	}

	if p1 == '(' {
		_, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		return TokenFunction{
			Value:          name,
			representation: t.Representation(),
		}
	}

	return TokenIdent{
		Value:          name,
		representation: t.Representation(),
	}
}
