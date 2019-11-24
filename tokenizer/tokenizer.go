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
				representation: append([]rune(nil), t.representation()...),
			}

		case ')': // Right Parenthesis
			return TokenParenthesisRight{
				representation: append([]rune(nil), t.representation()...),
			}

		case '[': // Left Square Bracket
			return TokenSquareBracketLeft{
				representation: append([]rune(nil), t.representation()...),
			}

		case ']': // Right Square Bracket
			return TokenSquareBracketRight{
				representation: append([]rune(nil), t.representation()...),
			}

		case '{': // Left Curly Bracket
			return TokenCurlyBracketLeft{
				representation: append([]rune(nil), t.representation()...),
			}

		case '}': // Right Curly Bracket
			return TokenCurlyBracketRight{
				representation: append([]rune(nil), t.representation()...),
			}

		case ',': // Comma
			return TokenComma{
				representation: append([]rune(nil), t.representation()...),
			}

		case ':': // Colon
			return TokenColon{
				representation: append([]rune(nil), t.representation()...),
			}

		case ';': // Semicolon
			return TokenSemicolon{
				representation: append([]rune(nil), t.representation()...),
			}

		case '+': // Plus
			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			if checkIfThreeCodePointsWouldStartANumber(t) {
				return consumeNumeric(t, r)
			}

			t.ReadRune()

			return TokenDelim{
				Value:          r,
				representation: append([]rune(nil), t.representation()...),
			}

		case '-': // Minus
			p1, p2, _ := t.peekTwoRunes()
			if p1 == '-' && p2 == '>' {
				t.ReadRune()
				t.ReadRune()

				return TokenCDC{
					representation: append([]rune(nil), t.representation()...),
				}
			}

			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			if checkIfThreeCodePointsWouldStartANumber(t) {
				return consumeNumeric(t, r)
			}

			if checkIfThreeCodePointsWouldStartAnIdentifier(t) {
				return consumeIdentLikeToken(t)
			}

			return TokenDelim{
				Value:          r,
				representation: append([]rune(nil), t.representation()...),
			}

		case '\'', '"': // String
			return consumeString(t, r)

		case '\u000a', '\u0009', '\u0020': // Whitespace
			return consumeWhiteSpace(t, -1)

		case '/': // Comment
			token := consumeComment(t)
			if token != nil {
				// Should return nothing
				// https://drafts.csswg.org/css-syntax-3/#consume-comment
				// But comments can be interesting in our context so return a TokenComment
				return token
			}

		case '@': // Comment
			if checkIfThreeCodePointsWouldStartAnIdentifier(t) {
				name, err := consumeName(t)
				if err != nil {
					return TokenError{error: err}
				}

				return TokenAtKeyword{
					Value:          append([]rune(nil), name...),
					representation: append([]rune(nil), t.representation()...),
				}
			}

			return TokenDelim{
				Value:          r,
				representation: append([]rune(nil), t.representation()...),
			}
		case '#': // Number Sign
			return TokenizeHashFromNumberSign(t)
		}

		if unicode.In(r, unicode.Digit) {
			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			if checkIfThreeCodePointsWouldStartANumber(t) {
				return consumeNumeric(t, r)
			}
		}

		if unicode.In(r, NameStartCodePoint...) {
			err := t.reader.UnreadRune(r, size)
			if err != nil {
				return TokenError{error: err}
			}

			return consumeIdentLikeToken(t)
		}

		return TokenDelim{
			Value:          r,
			representation: append([]rune(nil), t.representation()...),
		}
	}
}
