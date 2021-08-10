package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenSemicolon struct {
	representation []rune
}

func (t TokenSemicolon) String() string {
	return ";"
}

func (t TokenSemicolon) Representation() []rune {
	return t.representation
}

func (t TokenSemicolon) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(";"))
}
