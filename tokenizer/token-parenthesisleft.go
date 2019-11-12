package tokenizer

type TokenParenthesisLeft struct {
	representation []rune
}

func (t TokenParenthesisLeft) String() string {
	return "("
}

func (t TokenParenthesisLeft) Representation() []rune {
	return t.representation
}
