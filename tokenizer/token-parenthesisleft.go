package tokenizer

type TokenParenthesisLeft struct {
	represenation []rune
}

func (t TokenParenthesisLeft) String() string {
	return "("
}

func (t TokenParenthesisLeft) Representation() []rune {
	return t.represenation
}
