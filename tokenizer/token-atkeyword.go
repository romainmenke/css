package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenAtKeyword struct {
	Value          []rune
	representation []rune
}

func (t TokenAtKeyword) String() string {
	return string(t.Value)
}

func (t TokenAtKeyword) Representation() []rune {
	return t.representation
}

func (t TokenAtKeyword) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(string(t.Value)))
}
