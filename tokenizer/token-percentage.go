package tokenizer

import (
	"fmt"
	"io"

	"github.com/romainmenke/css/serializer"
)

type TokenPercentage struct {
	TokenNumber
}

func (t TokenPercentage) String() string {
	if t.Type == NumberTypeInteger {
		return fmt.Sprintf("%d", int64(t.intValue)) + "%"
	}

	// See : https://www.w3.org/TR/cssom-1/#serializing-css-values
	// rounding the value if necessary to not produce more than 6 decimals
	return fmt.Sprintf("%.6f", t.floatValue) + "%"
}

func (t TokenPercentage) Representation() []rune {
	return t.representation
}

func (t TokenPercentage) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return w.Write([]byte(t.String()))
}

func NewTokenPercentageInt(v int, representation []rune) TokenPercentage {
	return TokenPercentage{
		TokenNumber: TokenNumber{
			floatValue:     float64(v),
			intValue:       int64(v),
			Type:           NumberTypeInteger,
			representation: append([]rune(nil), representation...),
		},
	}
}

func NewTokenPercentageInt64(v int64, representation []rune) TokenPercentage {
	return TokenPercentage{
		TokenNumber: TokenNumber{
			floatValue:     float64(v),
			intValue:       v,
			Type:           NumberTypeInteger,
			representation: append([]rune(nil), representation...),
		},
	}
}

func NewTokenPercentageFloat64(v float64, representation []rune) TokenPercentage {
	return TokenPercentage{
		TokenNumber: TokenNumber{
			floatValue:     v,
			intValue:       int64(v),
			Type:           NumberTypeInteger,
			representation: append([]rune(nil), representation...),
		},
	}
}
