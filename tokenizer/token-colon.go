package tokenizer

type TokenColon struct {
	representation []rune
}

func (t TokenColon) String() string {
	return ":"
}

func (t TokenColon) Representation() []rune {
	return t.representation
}
