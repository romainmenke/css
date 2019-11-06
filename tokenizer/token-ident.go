package tokenizer

type TokenIdent struct {
	Value []rune
}

func (t TokenIdent) String() string {
	return string(t.Value)
}
