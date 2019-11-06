package tokenizer

type TokenWhitespace struct {
	represenation []rune
}

func (t TokenWhitespace) String() string {
	return " " // collapsed
}

func (t TokenWhitespace) Representation() []rune {
	return t.represenation
}
