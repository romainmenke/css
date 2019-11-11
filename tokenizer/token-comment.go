package tokenizer

type TokenComment struct {
	Value         []rune
	representation []rune
}

func (t TokenComment) String() string {
	return string(t.Value)
}

func (t TokenComment) Representation() []rune {
	return t.representation
}
