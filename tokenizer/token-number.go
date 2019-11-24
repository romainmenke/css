package tokenizer

import "fmt"

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
	return fmt.Sprintf("%f", t.floatValue)
}

func (t TokenNumber) Representation() []rune {
	return t.representation
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
