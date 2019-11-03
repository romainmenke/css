package tokenizer

type BracketCloseToken struct{}

func (t BracketCloseToken) String() string {
	return ")"
}
