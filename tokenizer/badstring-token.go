package tokenizer

type BadStringToken struct{}

func (t BadStringToken) String() string {
	return ""
}
