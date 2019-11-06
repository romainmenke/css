package tokenizer

type TokenColon struct {
	represenation []rune
}

func (t TokenColon) String() string {
	return ":"
}

func (t TokenColon) Representation() []rune {
	return t.represenation
}
