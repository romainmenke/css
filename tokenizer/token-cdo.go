package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenCDO struct {
	representation []rune
}

func (t TokenCDO) String() string {
	return "<!--"
}

func (t TokenCDO) Representation() []rune {
	return t.representation
}

func (t TokenCDO) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte("<!--"))
}
