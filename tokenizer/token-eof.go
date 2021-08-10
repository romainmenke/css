package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenEOF struct{}

func (t TokenEOF) String() string {
	return "EOF"
}

func (t TokenEOF) Representation() []rune {
	return nil
}

func (t TokenEOF) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return 0, nil
}
