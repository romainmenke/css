package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenComma struct {
	representation []rune
}

func (t TokenComma) String() string {
	return ","
}

func (t TokenComma) Representation() []rune {
	return t.representation
}

func (t TokenComma) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(","))
}
