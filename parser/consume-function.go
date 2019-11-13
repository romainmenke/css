package parser

import (
	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func consumeFunction(p *Parser, inputToken tokenizer.Token) stylesheet.Function {
	function := stylesheet.Function{
		Name: inputToken.String(),
	}

	for {
		t := p.tz.Next()
		if t == nil {
			return function
		}

		switch token := t.(type) {
		case tokenizer.TokenParenthesisRight:
			return function
		case tokenizer.TokenEOF:
			return function
		default:
			function.Value = append(function.Value, consumeComponentValue(p, token))
		}
	}
}
