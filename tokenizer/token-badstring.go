package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenBadString struct{}

func (t TokenBadString) String() string {
	return ""
}

func (t TokenBadString) Representation() []rune {
	return nil
}

func (t TokenBadString) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return 0, nil
}
