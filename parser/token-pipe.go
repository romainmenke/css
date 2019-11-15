package parser

import "github.com/romainmenke/css/tokenizer"

type tokenStream interface {
	Next() tokenizer.Token
}

type tokenStreamFunc func() tokenizer.Token

func (f tokenStreamFunc) Next() tokenizer.Token {
	return f()
}

type streamCondition func(tokenizer.Token) bool

func pipeWithCondition(s tokenStream, condition streamCondition) tokenStreamFunc {
	return tokenStreamFunc(func() tokenizer.Token {
		next := s.Next()
		if condition(next) && next != nil {
			return next
		}

		return tokenizer.TokenEOF{}
	})
}
