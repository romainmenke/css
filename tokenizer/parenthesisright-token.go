package tokenizer

type TokenParenthesisRight struct {
	represenation []rune
}

func (t TokenParenthesisRight) String() string {
	return ")"
}

func (t TokenParenthesisRight) Representation() []rune {
	return t.represenation
}
