package tokenizer

import (
	"bufio"
	"io"

	"github.com/romainmenke/css/tokenizer/runepeeker"
)

type Tokenizer struct {
	reader   *runepeeker.Peeker
	tracking []rune
}

func New(r io.Reader) *Tokenizer {
	return &Tokenizer{
		reader:   runepeeker.New(bufio.NewReader(r)),
		tracking: make([]rune, 0, 1000),
	}
}

func (t *Tokenizer) Next() Token {
	t.tracking = t.tracking[:0]
	t.reader.ResetRepresentation()

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
				represenation: t.Representation(),
			}

		case ')': // Right Parenthesis
			return TokenParenthesisRight{
				represenation: t.Representation(),
			}

		case '[': // Left Square Bracket
			return TokenSquareBracketLeft{
				represenation: t.Representation(),
			}

		case ']': // Right Square Bracket
			return TokenSquareBracketRight{
				represenation: t.Representation(),
			}

		case '{': // Left Curly Bracket
			return TokenCurlyBracketLeft{
				represenation: t.Representation(),
			}

		case '}': // Right Curly Bracket
			return TokenCurlyBracketRight{
				represenation: t.Representation(),
			}

		case ',': // Comma
			return TokenComma{
				represenation: t.Representation(),
			}

		case ':': // Colon
			return TokenColon{
				represenation: t.Representation(),
			}

		case ';': // Semicolon
			return TokenSemicolon{
				represenation: t.Representation(),
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
			represenation: t.Representation(),
		}
	}
}
