package tokenizer

type TokenAtKeyword struct {
	Value []rune
}

func (t TokenAtKeyword) String() string {
	return string(t.Value)
}
