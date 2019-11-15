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

func (t TokenParenthesisRight) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenParenthesisLeft:
		return true
	default:
		return false
	}
}
