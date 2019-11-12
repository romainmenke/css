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
