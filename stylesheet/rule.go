package stylesheet

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type Rule interface {
}

// AtRule has a name, a prelude consisting of a list of component values, and an optional block consisting of a simple {} block.
type AtRule struct {
	Block   Block
	Name    string
	Prelude []interface{}
}

func (r AtRule) String() string {
	return ""
}

func (r AtRule) Serialize(w io.Writer, options serializer.Options) (int, error) {
	var (
		n int
	)

	nn, err := w.Write([]byte("@" + r.Name))
	n += nn
	if err != nil {
		return n, err
	}

	if len(r.Prelude) > 0 {
		nn, err = serializer.Serialize(w, r.Prelude, options)
		n += nn
		if err != nil {
			return n, err
		}
	}

	if r.Block != nil {
		nn, err = r.Block.Serialize(w, options)
		n += nn
		if err != nil {
			return n, err
		}
	}

	return n, nil
}

// QualifiedRule has a prelude consisting of a list of component values, and a block consisting of a simple {} block.
type QualifiedRule struct {
	Block   Block
	Prelude []interface{}
}

func (r QualifiedRule) String() string {
	return ""
}

func (r QualifiedRule) Serialize(w io.Writer, options serializer.Options) (int, error) {
	var (
		n int
	)

	nn, err := serializer.Serialize(w, r.Prelude, options)
	n += nn
	if err != nil {
		return n, err
	}

	if r.Block != nil {
		nn, err = r.Block.Serialize(w, options)
		n += nn
		if err != nil {
			return n, err
		}
	}

	return n, nil
}

type RuleList []interface{}
