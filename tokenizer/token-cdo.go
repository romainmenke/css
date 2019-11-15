package tokenizer

type TokenCDO struct {
	representation []rune
}

func (t TokenCDO) String() string {
	return "<!--"
}

func (t TokenCDO) Representation() []rune {
	return t.representation
}
