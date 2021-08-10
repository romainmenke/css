package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenCurlyBracketLeft struct {
	representation []rune
}

func (t TokenCurlyBracketLeft) String() string {
	return "{"
}

func (t TokenCurlyBracketLeft) Representation() []rune {
	return t.representation
}

func (t TokenCurlyBracketLeft) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenCurlyBracketRight:
		return true
	default:
		return false
	}
}

func (t TokenCurlyBracketLeft) Mirror() TokenCurlyBracketRight {
	return TokenCurlyBracketRight{}
}

func (t TokenCurlyBracketLeft) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("{"))
}
