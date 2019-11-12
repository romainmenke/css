package tokenizer

type TokenCDC struct {
	representation []rune
}

func (t TokenCDC) String() string {
	return "-->"
}

func (t TokenCDC) Representation() []rune {
	return t.representation
}
