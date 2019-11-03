package tokenizer

type SquareBracketCloseToken struct{}

func (t SquareBracketCloseToken) String() string {
	return "]"
}
