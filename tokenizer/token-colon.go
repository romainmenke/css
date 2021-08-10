package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenColon struct {
	representation []rune
}

func (t TokenColon) String() string {
	return ":"
}

func (t TokenColon) Representation() []rune {
	return t.representation
}

func (t TokenColon) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(":"))
}
