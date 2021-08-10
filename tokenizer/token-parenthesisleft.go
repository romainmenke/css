package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenParenthesisLeft struct {
	representation []rune
}

func (t TokenParenthesisLeft) String() string {
	return "("
}

func (t TokenParenthesisLeft) Representation() []rune {
	return t.representation
}

func (t TokenParenthesisLeft) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenParenthesisRight:
		return true
	default:
		return false
	}
}

func (t TokenParenthesisLeft) Mirror() TokenParenthesisRight {
	return TokenParenthesisRight{}
}

func (t TokenParenthesisLeft) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("("))
}
