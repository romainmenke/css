package tokenizer

type TokenWhitespace struct {
	representation []rune
}

func (t TokenWhitespace) String() string {
	return " " // collapsed
}

func (t TokenWhitespace) Representation() []rune {
	return t.representation
}
