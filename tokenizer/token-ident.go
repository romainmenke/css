package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenIdent struct {
	Value          []rune
	representation []rune
}

func (t TokenIdent) String() string {
	return string(t.Value)
}

func (t TokenIdent) Representation() []rune {
	return t.representation
}

func (t TokenIdent) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(string(t.Value)))
}
