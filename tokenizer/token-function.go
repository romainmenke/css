package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenFunction struct {
	Value          []rune
	representation []rune
}

func (t TokenFunction) String() string {
	return string(t.Value)
}

func (t TokenFunction) Representation() []rune {
	return t.representation
}

func (t TokenFunction) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(string(t.Value)))
}
