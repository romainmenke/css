package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenUrl struct {
	Value          []rune
	representation []rune
}

func (t TokenUrl) String() string {
	return string(t.Value)
}

func (t TokenUrl) Representation() []rune {
	return t.representation
}

func (t TokenUrl) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("url(" + string(t.Value) + ")"))
}
