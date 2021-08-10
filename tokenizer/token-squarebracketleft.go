package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenSquareBracketLeft struct {
	representation []rune
}

func (t TokenSquareBracketLeft) String() string {
	return "["
}

func (t TokenSquareBracketLeft) Representation() []rune {
	return t.representation
}

func (t TokenSquareBracketLeft) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenSquareBracketRight:
		return true
	default:
		return false
	}
}

func (t TokenSquareBracketLeft) Mirror() TokenSquareBracketRight {
	return TokenSquareBracketRight{}
}

func (t TokenSquareBracketLeft) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("["))
}
