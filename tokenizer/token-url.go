package tokenizer

type TokenUrl struct {
	Value          []rune
	representation []rune
}

func (t TokenUrl) String() string {
	return string(t.Value)
}

func (t TokenUrl) Representation() []rune {
	return t.representation
}
