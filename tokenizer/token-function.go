package tokenizer

type TokenFunction struct {
	Value []rune
}

func (t TokenFunction) String() string {
	return string(t.Value)
}
