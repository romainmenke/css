package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenComment struct {
	Value          []rune
	representation []rune
}

func (t TokenComment) String() string {
	return string(t.Value)
}

func (t TokenComment) Representation() []rune {
	return t.representation
}

func (t TokenComment) Serialize(w io.Writer, options serializer.Options) (int, error) {
	if options.Flag.Has(serializer.Minify) {
		return 0, nil
	}

	return w.Write([]byte("/*" + string(t.Value) + "*/"))
}
