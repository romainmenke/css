package tokenizer

type TokenBadUrl struct {
	representation []rune
}

func (t TokenBadUrl) String() string {
	return string(t.representation)
}

func (t TokenBadUrl) Representation() []rune {
	return t.representation
}
