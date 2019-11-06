package tokenizer

type TokenCurlyBracketLeft struct {
	represenation []rune
}

func (t TokenCurlyBracketLeft) String() string {
	return "{"
}

func (t TokenCurlyBracketLeft) Representation() []rune {
	return t.represenation
}
