package tokenizer

type TokenEOF struct{}

func (t TokenEOF) String() string {
	return "EOF"
}
