package tokenizer

type TokenUrl struct {
	Value []rune
}

func (t TokenUrl) String() string {
	return string(t.Value)
}
