package tokenizer

type TokenSemicolon struct {
	representation []rune
}

func (t TokenSemicolon) String() string {
	return ";"
}

func (t TokenSemicolon) Representation() []rune {
	return t.representation
}
