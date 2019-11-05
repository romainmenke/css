package tokenizer

type Token interface {
	String() string
	Representation() []rune
}
