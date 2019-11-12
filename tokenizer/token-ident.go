package tokenizer

type TokenIdent struct {
	Value          []rune
	representation []rune
}

func (t TokenIdent) String() string {
	return string(t.Value)
}

func (t TokenIdent) Representation() []rune {
	return t.representation
}
