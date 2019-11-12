package tokenizer

type TokenString struct {
	Value          []rune
	representation []rune
	Quote          QuoteKind
}

func (t TokenString) String() string {
	return string(t.Value)
}

func (t TokenString) Representation() []rune {
	return t.representation
}

type QuoteKind int

const SingleQuote QuoteKind = 0
const DoubleQuote QuoteKind = 1
