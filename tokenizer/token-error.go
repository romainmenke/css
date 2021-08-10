package tokenizer

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

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

func (t TokenError) Representation() []rune {
	return nil
}

func (t TokenError) Serialize(w io.Writer, options serializer.Options) (int, error) {
	return 0, nil
}
