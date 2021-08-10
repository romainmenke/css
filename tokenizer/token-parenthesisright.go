package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenParenthesisRight struct {
	representation []rune
}

func (t TokenParenthesisRight) String() string {
	return ")"
}

func (t TokenParenthesisRight) Representation() []rune {
	return t.representation
}

func (t TokenParenthesisRight) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenParenthesisLeft:
		return true
	default:
		return false
	}
}

func (t TokenParenthesisRight) Mirror() TokenParenthesisLeft {
	return TokenParenthesisLeft{}
}

func (t TokenParenthesisRight) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(")"))
}
