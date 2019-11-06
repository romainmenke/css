package tokenizer

type TokenNumber struct {
	Value float64
	Type  NumberType
}

func (t TokenNumber) String() string {
	return ""
}
