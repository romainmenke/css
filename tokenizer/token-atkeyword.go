package tokenizer

type TokenAtKeyword struct {
	Value          []rune
	representation []rune
}

func (t TokenAtKeyword) String() string {
	return string(t.Value)
}

func (t TokenAtKeyword) Representation() []rune {
	return t.representation
}
