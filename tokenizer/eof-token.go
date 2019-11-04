package tokenizer

type EOFToken struct{}

func (t EOFToken) String() string {
	return "EOF"
}
