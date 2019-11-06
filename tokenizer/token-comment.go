package tokenizer

type TokenComment struct {
	Value         []rune
	represenation []rune
}

func (t TokenComment) String() string {
	return string(t.Value)
}

func (t TokenComment) Representation() []rune {
	return t.represenation
}
