package tokenizer

type TokenSquareBracketLeft struct {
	represenation []rune
}

func (t TokenSquareBracketLeft) String() string {
	return "["
}

func (t TokenSquareBracketLeft) Representation() []rune {
	return t.represenation
}
