package tokenizer

import "fmt"

type TokenDimension struct {
	TokenNumber

	Unit []rune
}

func (t TokenDimension) String() string {
	if t.Type == NumberTypeInteger {
		return fmt.Sprintf("%d", int64(t.intValue)) + string(t.Unit)
	}
	return fmt.Sprintf("%f", t.floatValue) + string(t.Unit)
}

func (t TokenDimension) Representation() []rune {
	return t.representation
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
