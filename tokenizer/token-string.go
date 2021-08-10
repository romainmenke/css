package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenString struct {
	Value          []rune
	representation []rune
	Quote          QuoteKind
}

func (t TokenString) String() string {
	return string(t.Value)
}

func (t TokenString) Representation() []rune {
	return t.representation
}

func (t TokenString) Serialize(w io.Writer, options serializer.Options) (int, error) {
	if t.Quote == SingleQuote {
		return w.Write([]byte("'" + string(t.Value) + "'"))
	}

	return w.Write([]byte(`"` + string(t.Value) + `"`))
}

type QuoteKind int

const SingleQuote QuoteKind = 0
const DoubleQuote QuoteKind = 1
