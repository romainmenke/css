package tokenizer

import (
	"fmt"
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenNumber struct {
	intValue       int64
	floatValue     float64
	Type           NumberType
	representation []rune
}

func (t TokenNumber) IntValue() int64 {
	return t.intValue
}

func (t TokenNumber) FloatValue() float64 {
	return t.floatValue
}

func (t TokenNumber) String() string {
	if t.Type == NumberTypeInteger {
		return fmt.Sprintf("%d", int64(t.intValue))
	}

	// See : https://www.w3.org/TR/cssom-1/#serializing-css-values
	// rounding the value if necessary to not produce more than 6 decimals
	return fmt.Sprintf("%.6f", t.floatValue)
}

func (t TokenNumber) Representation() []rune {
	return t.representation
}

func (t TokenNumber) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(t.String()))
}

func NewTokenNumberInt(v int, representation []rune) TokenNumber {
	return TokenNumber{
		floatValue:     float64(v),
		intValue:       int64(v),
		Type:           NumberTypeInteger,
		representation: append([]rune(nil), representation...),
	}
}

func NewTokenNumberInt64(v int64, representation []rune) TokenNumber {
	return TokenNumber{
		floatValue:     float64(v),
		intValue:       v,
		Type:           NumberTypeInteger,
		representation: append([]rune(nil), representation...),
	}
}

func NewTokenNumberFloat64(v float64, representation []rune) TokenNumber {
	return TokenNumber{
		floatValue:     v,
		intValue:       int64(v),
		Type:           NumberTypeInteger,
		representation: append([]rune(nil), representation...),
	}
}
