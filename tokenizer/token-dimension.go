package tokenizer

import (
	"fmt"
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenDimension struct {
	TokenNumber

	Unit []rune
}

func (t TokenDimension) String() string {
	if t.Type == NumberTypeInteger {
		return fmt.Sprintf("%d", int64(t.intValue)) + string(t.Unit)
	}

	// See : https://www.w3.org/TR/cssom-1/#serializing-css-values
	// rounding the value if necessary to not produce more than 6 decimals
	return fmt.Sprintf("%.6f", t.floatValue) + string(t.Unit)
}

func (t TokenDimension) Representation() []rune {
	return t.representation
}

func (t TokenDimension) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(t.String()))
}

func NewTokenDimensionInt(v int, unit []rune, representation []rune) TokenDimension {
	return TokenDimension{
		TokenNumber: TokenNumber{
			floatValue:     float64(v),
			intValue:       int64(v),
			Type:           NumberTypeInteger,
			representation: append([]rune(nil), representation...),
		},
		Unit: append([]rune(nil), unit...),
	}
}

func NewTokenDimensionInt64(v int64, unit []rune, representation []rune) TokenDimension {
	return TokenDimension{
		TokenNumber: TokenNumber{
			floatValue:     float64(v),
			intValue:       v,
			Type:           NumberTypeInteger,
			representation: append([]rune(nil), representation...),
		},
		Unit: append([]rune(nil), unit...),
	}
}

func NewTokenDimensionFloat64(v float64, unit []rune, representation []rune) TokenDimension {
	return TokenDimension{
		TokenNumber: TokenNumber{
			floatValue:     v,
			intValue:       int64(v),
			Type:           NumberTypeInteger,
			representation: append([]rune(nil), representation...),
		},
		Unit: append([]rune(nil), unit...),
	}
}
