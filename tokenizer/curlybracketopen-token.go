package tokenizer

type CurlyBracketOpenToken struct{}

func (t CurlyBracketOpenToken) String() string {
	return "{"
}
