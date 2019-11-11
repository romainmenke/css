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
