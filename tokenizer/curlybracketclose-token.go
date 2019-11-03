package tokenizer

type CurlyBracketCloseToken struct{}

func (t CurlyBracketCloseToken) String() string {
	return "}"
}
