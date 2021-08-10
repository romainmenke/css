package stylesheet

import (
	"io"

	"github.com/romainmenke/css/serializer"
	"github.com/romainmenke/css/tokenizer"
)

type Block interface {
	Serialize(w io.Writer, options serializer.Options) (int, error)
}

// SimpleBlock has an associated token (either a <[-token>, <(-token>, or <{-token>) and a value consisting of a list of component values.
type SimpleBlock struct {
	AssociatedToken tokenizer.Token
	Value           []interface{}
}

func (b SimpleBlock) String() string {
	return ""
}

func (b SimpleBlock) Serialize(w io.Writer, options serializer.Options) (int, error) {
	var (
		n int
	)

	nn, err := w.Write([]byte(b.AssociatedToken.String()))
	n += nn
	if err != nil {
		return n, err
	}

	nn, err = serializer.Serialize(w, b.Value, options)
	n += nn
	if err != nil {
		return n, err
	}

	if m, ok := b.AssociatedToken.(tokenizer.MirroreableToken); ok {
		nn, err := w.Write([]byte(m.Mirror().String()))
		n += nn
		if err != nil {
			return n, err
		}

	}

	return n, nil
}
