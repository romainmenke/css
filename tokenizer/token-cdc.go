package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenCDC struct {
	representation []rune
}

func (t TokenCDC) String() string {
	return "-->"
}

func (t TokenCDC) Representation() []rune {
	return t.representation
}

func (t TokenCDC) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("-->"))
}
