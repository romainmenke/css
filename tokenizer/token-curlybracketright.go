package tokenizer

type TokenCurlyBracketRight struct {
	representation []rune
}

func (t TokenCurlyBracketRight) String() string {
	return "}"
}

func (t TokenCurlyBracketRight) Representation() []rune {
	return t.representation
}

func (t TokenCurlyBracketRight) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenCurlyBracketLeft:
		return true
	default:
		return false
	}
}
