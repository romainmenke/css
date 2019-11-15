package tokenizer

type TokenDelim struct {
	Value          rune
	representation []rune
}

func (t TokenDelim) String() string {
	return string(t.Value)
}

func (t TokenDelim) Representation() []rune {
	return t.representation
}
