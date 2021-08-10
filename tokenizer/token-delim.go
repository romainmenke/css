package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenDelim struct {
	Value          rune
	representation []rune
}

func (t TokenDelim) String() string {
	return string(t.Value)
}

func (t TokenDelim) Representation() []rune {
	return t.representation
}

func (t TokenDelim) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(string(t.Value)))
}
