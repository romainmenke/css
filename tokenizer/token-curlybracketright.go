package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenCurlyBracketRight struct {
	representation []rune
}

func (t TokenCurlyBracketRight) String() string {
	return "}"
}

func (t TokenCurlyBracketRight) Representation() []rune {
	return t.representation
}

func (t TokenCurlyBracketRight) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenCurlyBracketLeft:
		return true
	default:
		return false
	}
}

func (t TokenCurlyBracketRight) Mirror() TokenCurlyBracketLeft {
	return TokenCurlyBracketLeft{}
}

func (t TokenCurlyBracketRight) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("}"))
}
