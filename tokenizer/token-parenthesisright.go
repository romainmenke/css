package tokenizer

type TokenParenthesisRight struct {
	representation []rune
}

func (t TokenParenthesisRight) String() string {
	return ")"
}

func (t TokenParenthesisRight) Representation() []rune {
	return t.representation
}
