package tokenizer

import "fmt"

type TokenPercentage struct {
	TokenNumber
}

func (t TokenPercentage) String() string {
	if t.Type == NumberTypeInteger {
		return fmt.Sprintf("%d", int64(t.intValue)) + "%"
	}
	return fmt.Sprintf("%f", t.floatValue) + "%"
}

func (t TokenPercentage) Representation() []rune {
	return t.representation
}

func NewTokenPercentageInt(v int, representation []rune) TokenPercentage {
	return TokenPercentage{
		TokenNumber: TokenNumber{
			floatValue:     float64(v),
			intValue:       int64(v),
			Type:           NumberTypeInteger,
			representation: representation,
		},
	}
}

func NewTokenPercentageInt64(v int64, representation []rune) TokenPercentage {
	return TokenPercentage{
		TokenNumber: TokenNumber{
			floatValue:     float64(v),
			intValue:       v,
			Type:           NumberTypeInteger,
			representation: representation,
		},
	}
}

func NewTokenPercentageFloat64(v float64, representation []rune) TokenPercentage {
	return TokenPercentage{
		TokenNumber: TokenNumber{
			floatValue:     v,
			intValue:       int64(v),
			Type:           NumberTypeInteger,
			representation: representation,
		},
	}
}
