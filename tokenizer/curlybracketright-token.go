package tokenizer

type TokenCurlyBracketRight struct {
	represenation []rune
}

func (t TokenCurlyBracketRight) String() string {
	return "}"
}

func (t TokenCurlyBracketRight) Representation() []rune {
	return t.represenation
}
