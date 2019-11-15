package tokenizer

type TokenCurlyBracketLeft struct {
	representation []rune
}

func (t TokenCurlyBracketLeft) String() string {
	return "{"
}

func (t TokenCurlyBracketLeft) Representation() []rune {
	return t.representation
}

func (t TokenCurlyBracketLeft) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenCurlyBracketRight:
		return true
	default:
		return false
	}
}
