package tokenizer

type TokenDimension struct {
	Value float64
	Type  NumberType
	Unit  []rune
}

func (t TokenDimension) String() string {
	return ""
}
