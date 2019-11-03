package tokenizer

type DimensionToken struct {
	Value float64
	Type  NumberType
}

func (t DimensionToken) String() string {
	return ""
}
