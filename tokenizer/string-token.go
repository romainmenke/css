package tokenizer

type StringToken struct {
	Value []byte
	Quote QuoteKind
}

func (t StringToken) String() string {
	if t.Quote == SingleQuote {
		return `'` + string(t.Value) + `'`
	}

	return `"` + string(t.Value) + `"`
}

type QuoteKind int

const SingleQuote QuoteKind = 0
const DoubleQuote QuoteKind = 1
