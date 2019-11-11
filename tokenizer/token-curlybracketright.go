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
