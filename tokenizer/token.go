package tokenizer

type Token interface {
	String() string
	Representation() []rune
}

type MirroreableToken interface {
	Token
	IsMirror(Token) bool
}
