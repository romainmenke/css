package tokenizer

type NumberToken struct {
	Value float64
	Type  NumberType
}

func (t NumberToken) String() string {
	return ""
}
