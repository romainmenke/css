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

func (t TokenParenthesisLeft) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenParenthesisRight:
		return true
	default:
		return false
	}
}
