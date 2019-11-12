package tokenizer

type TokenEOF struct{}

func (t TokenEOF) String() string {
	return "EOF"
}

func (t TokenEOF) Representation() []rune {
	return nil
}
