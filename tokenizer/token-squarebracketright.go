package tokenizer

type TokenSquareBracketRight struct {
	representation []rune
}

func (t TokenSquareBracketRight) String() string {
	return "]"
}

func (t TokenSquareBracketRight) Representation() []rune {
	return t.representation
}
