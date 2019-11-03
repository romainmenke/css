package tokenizer

type SemicolonToken struct{}

func (t SemicolonToken) String() string {
	return ";"
}
