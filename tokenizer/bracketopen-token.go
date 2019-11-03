package tokenizer

type BracketOpenToken struct{}

func (t BracketOpenToken) String() string {
	return "("
}
