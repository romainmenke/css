package tokenizer

type TokenString struct {
	Value         []rune
	represenation []rune
	Quote         QuoteKind
}

func (t TokenString) String() string {
	return string(t.Value)
}

func (t TokenString) Representation() []rune {
	return t.represenation
}

type QuoteKind int

const SingleQuote QuoteKind = 0
const DoubleQuote QuoteKind = 1
