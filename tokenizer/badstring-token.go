package tokenizer

type TokenBadString struct{}

func (t TokenBadString) String() string {
	return ""
}

func (t TokenBadString) Representation() []rune {
	return nil
}
