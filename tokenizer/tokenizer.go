package tokenizer

import (
	"bufio"
	"io"
	"unicode"

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
		r, size, err := t.ReadRune()
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
				representation: t.Representation(),
			}

		case ')': // Right Parenthesis
			return TokenParenthesisRight{
				representation: t.Representation(),
			}

		case '[': // Left Square Bracket
			return TokenSquareBracketLeft{
				representation: t.Representation(),
			}

		case ']': // Right Square Bracket
			return TokenSquareBracketRight{
				representation: t.Representation(),
			}

		case '{': // Left Curly Bracket
			return TokenCurlyBracketLeft{
				representation: t.Representation(),
			}

		case '}': // Right Curly Bracket
			return TokenCurlyBracketRight{
				representation: t.Representation(),
			}

		case ',': // Comma
			return TokenComma{
				representation: t.Representation(),
			}

		case ':': // Colon
			return TokenColon{
				representation: t.Representation(),
			}

		case ';': // Semicolon
			return TokenSemicolon{
				representation: t.Representation(),
			}

		case '+': // Plus
			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			if CheckIfThreeCodePointsWouldStartANumber(t) {
				return ConsumeNumeric(t, r)
			}

			return TokenDelim{
				Value:          r,
				representation: t.Representation(),
			}

		case '-': // Minus
			p1, p2, _ := t.PeekTwoRunes()
			if p1 == '-' && p2 == '>' {
				t.ReadRune()
				t.ReadRune()

				return TokenCDC{
					representation: t.Representation(),
				}
			}

			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			if CheckIfThreeCodePointsWouldStartANumber(t) {
				return ConsumeNumeric(t, r)
			}

			if CheckIfThreeCodePointsWouldStartAnIdentifier(t) {
				return ConsumeIdentLikeToken(t)
			}

			return TokenDelim{
				Value:          r,
				representation: t.Representation(),
			}

		case '\'', '"': // String
			return ConsumeString(t, r)

		case '\u000a', '\u0009', '\u0020': // Whitespace
			return ConsumeWhiteSpace(t, -1)

		case '/': // Comment
			token := ConsumeComment(t)
			if token != nil {
				// Should return nothing
				// https://drafts.csswg.org/css-syntax-3/#consume-comment
				// But comments can be interesting in our context so return a TokenComment
				return token
			}

		case '@': // Comment
			if CheckIfThreeCodePointsWouldStartAnIdentifier(t) {
				name, err := ConsumeName(t)
				if err != nil {
					return TokenError{error: err}
				}

				return TokenAtKeyword{
					Value:          name,
					representation: t.Representation(),
				}
			}

			return TokenDelim{
				Value:          r,
				representation: t.Representation(),
			}
		case '#': // Number Sign
			return TokenizeHashFromNumberSign(t)
		}

		if unicode.In(r, unicode.Digit) {
			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			if CheckIfThreeCodePointsWouldStartANumber(t) {
				return ConsumeNumeric(t, r)
			}
		}

		if unicode.In(r, NameStartCodePoint...) {
			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			return ConsumeIdentLikeToken(t)
		}

		return TokenDelim{
			Value:          r,
			representation: t.Representation(),
		}
	}
}
