package tokenizer

import (
	"bufio"
	"io"
)

type Tokenizer struct {
	b              *bufio.Reader
	tracking       []rune
	representation []rune
}

func New(r io.Reader) *Tokenizer {
	return &Tokenizer{
		b:              bufio.NewReader(r),
		tracking:       make([]rune, 0, 1000),
		representation: make([]rune, 0, 1000),
	}
}

func (t *Tokenizer) Next() Token {
	t.tracking = t.tracking[:0]
	t.representation = t.representation[:0]

	for {
		r, _, err := t.ReadRune()
		if err == io.EOF || (err != nil && err.Error() == io.EOF.Error()) {
			return TokenEOF{}
		}
		if err != nil {
			return TokenError{error: err}
		}

		// Tokenize
		switch r {

		case '(': // Left Parenthesis
			return TokenParenthesisLeft{
				represenation: t.representation,
			}

		case ')': // Right Parenthesis
			return TokenParenthesisRight{
				represenation: t.representation,
			}

		case '[': // Left Square Bracket
			return TokenSquareBracketLeft{
				represenation: t.representation,
			}

		case ']': // Right Square Bracket
			return TokenSquareBracketRight{
				represenation: t.representation,
			}

		case '{': // Left Curly Bracket
			return TokenCurlyBracketLeft{
				represenation: t.representation,
			}

		case '}': // Right Curly Bracket
			return TokenCurlyBracketRight{
				represenation: t.representation,
			}

		case ',': // Comma
			return TokenComma{
				represenation: t.representation,
			}

		case ':': // Colon
			return TokenColon{
				represenation: t.representation,
			}

		case ';': // Semicolon
			return TokenSemicolon{
				represenation: t.representation,
			}

		case '\'', '"': // String
			return ConsumeString(t, r)

		case '\u000a', '\u0009', '\u0020': // Whitespace
			return ConsumeWhiteSpace(t)

		case '/': // Comment
			token := ConsumeComment(t)
			if token != nil {
				// Should return nothing
				// https://drafts.csswg.org/css-syntax-3/#consume-comment
				// But comments can be interesting in our context so return a TokenComment
				return token
			}

		case '#': // Number Sign
			return TokenizeHashFromNumberSign(t)
		}

		return TokenDelim{
			Value:         r,
			represenation: t.representation,
		}
	}
}
