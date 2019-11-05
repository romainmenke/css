package tokenizer

type TokenSquareBracketRight struct {
	represenation []rune
}

func (t TokenSquareBracketRight) String() string {
	return "]"
}

func (t TokenSquareBracketRight) Representation() []rune {
	return t.represenation
}
