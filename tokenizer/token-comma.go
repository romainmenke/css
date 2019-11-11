package tokenizer

type TokenComma struct {
	representation []rune
}

func (t TokenComma) String() string {
	return ","
}

func (t TokenComma) Representation() []rune {
	return t.representation
}
