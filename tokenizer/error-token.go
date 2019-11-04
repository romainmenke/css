package tokenizer

type ErrorToken struct {
	error error
}

func (t ErrorToken) String() string {
	return t.error.Error()
}

func (t ErrorToken) Error() string {
	return t.error.Error()
}

func (t ErrorToken) Err() error {
	return t.error
}
