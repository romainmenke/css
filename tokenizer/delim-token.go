package tokenizer

type TokenDelim struct {
	Value rune
}

func (t TokenDelim) String() string {
	return string(t.Value)
}
