package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenWhitespace struct {
	representation []rune
}

func (t TokenWhitespace) String() string {
	return " " // collapsed
}

func (t TokenWhitespace) Representation() []rune {
	return t.representation
}

func (t TokenWhitespace) Serialize(w io.Writer, options serializer.Options) (int, error) {
	if options.Flag.Has(serializer.Minify) {
		return 0, nil
	}

	return w.Write([]byte(" "))
}
