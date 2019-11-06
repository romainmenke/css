package tokenizer

type TokenComma struct {
	represenation []rune
}

func (t TokenComma) String() string {
	return ","
}

func (t TokenComma) Representation() []rune {
	return t.represenation
}
