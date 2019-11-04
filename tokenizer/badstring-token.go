package tokenizer

type TokenBadString struct{}

func (t TokenBadString) String() string {
	return ""
}
