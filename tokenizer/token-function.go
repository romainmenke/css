package tokenizer

type TokenFunction struct {
	Value          []rune
	representation []rune
}

func (t TokenFunction) String() string {
	return string(t.Value)
}

func (t TokenFunction) Representation() []rune {
	return t.representation
}
