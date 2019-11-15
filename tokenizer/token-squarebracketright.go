package tokenizer

type TokenSquareBracketRight struct {
	representation []rune
}

func (t TokenSquareBracketRight) String() string {
	return "]"
}

func (t TokenSquareBracketRight) Representation() []rune {
	return t.representation
}

func (t TokenSquareBracketRight) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenSquareBracketLeft:
		return true
	default:
		return false
	}
}
