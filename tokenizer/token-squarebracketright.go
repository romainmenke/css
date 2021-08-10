package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenSquareBracketRight struct {
	representation []rune
}

func (t TokenSquareBracketRight) String() string {
	return "]"
}

func (t TokenSquareBracketRight) Representation() []rune {
	return t.representation
}

func (t TokenSquareBracketRight) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenSquareBracketLeft:
		return true
	default:
		return false
	}
}

func (t TokenSquareBracketRight) Mirror() TokenSquareBracketLeft {
	return TokenSquareBracketLeft{}
}

func (t TokenSquareBracketRight) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("]"))
}
