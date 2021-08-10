package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenBadUrl struct {
	representation []rune
}

func (t TokenBadUrl) String() string {
	return string(t.representation)
}

func (t TokenBadUrl) Representation() []rune {
	return t.representation
}

func (t TokenBadUrl) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return 0, nil
}
