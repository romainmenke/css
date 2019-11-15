package tokenizer

type TokenSquareBracketLeft struct {
	representation []rune
}

func (t TokenSquareBracketLeft) String() string {
	return "["
}

func (t TokenSquareBracketLeft) Representation() []rune {
	return t.representation
}

func (t TokenSquareBracketLeft) IsMirror(m Token) bool {
	switch m.(type) {
	case TokenSquareBracketRight:
		return true
	default:
		return false
	}
}
