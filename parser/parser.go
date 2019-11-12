package parser

import (
	"fmt"
	"io"

	"github.com/romainmenke/css/stylesheet"
	"github.com/romainmenke/css/tokenizer"
)

func Parse(r io.Reader) (stylesheet.Stylesheet, error) {
	tz := tokenizer.New(r)

	for {
		t := tz.Next()
		if _, ok := t.(tokenizer.TokenEOF); ok || t == nil {
			return stylesheet.Stylesheet{}, nil
		}

		switch token := t.(type) {
		case tokenizer.TokenWhitespace:
			continue
		case tokenizer.TokenAtKeyword:
			continue
		default:
			fmt.Println(token)
		}

		break
	}

	return stylesheet.Stylesheet{}, nil
}
