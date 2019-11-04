package tokenizer

type TokenError struct {
	error error
}

func (t TokenError) String() string {
	return t.error.Error()
}

func (t TokenError) Error() string {
	return t.error.Error()
}

func (t TokenError) Err() error {
	return t.error
}
