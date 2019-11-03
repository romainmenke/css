package tokenizer

import (
	"bufio"
	"errors"
	"io"
)

type Tokenizer struct {
	b        *bufio.Reader
	tracking []byte
}

func New(r io.Reader) *Tokenizer {
	return &Tokenizer{
		b:        bufio.NewReader(r),
		tracking: make([]byte, 0, 1000),
	}
}

func (t *Tokenizer) Next() (Token, error) {
	return nil, errors.New("unimplemented")
}
