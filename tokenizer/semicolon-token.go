package tokenizer

type TokenSemicolon struct {
	represenation []rune
}

func (t TokenSemicolon) String() string {
	return ";"
}

func (t TokenSemicolon) Representation() []rune {
	return t.represenation
}
