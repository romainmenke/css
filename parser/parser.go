package parser

import (
	"io"

	"github.com/romainmenke/css/tokenizer"
)

type Parser struct {
	tz *tokenizer.Tokenizer
}

func New(r io.Reader) *Parser {
	return &Parser{
		tz: tokenizer.New(r),
	}
}
