package tokenizer

type TokenSquareBracketLeft struct {
	representation []rune
}

func (t TokenSquareBracketLeft) String() string {
	return "["
}

func (t TokenSquareBracketLeft) Representation() []rune {
	return t.representation
}
