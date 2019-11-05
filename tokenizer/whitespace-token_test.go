package tokenizer

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestTokenWhitespace_OnlySelf(t *testing.T) {
	sources := []string{
		string('\n'),
		string('\n') + string('\r'),
		string('\r') + string('\n'),
		string('\r'),
		string('\f'),
		string(' '),
		string('\t'),
	}

	for _, source := range sources {
		t.Run(strconv.Quote(source), func(t *testing.T) {
			tokenizer := New(bytes.NewBufferString(source))
			sawToken := false

			for {
				token := tokenizer.Next()
				if _, ok := token.(TokenEOF); ok && sawToken {
					break
				}

				if errToken, ok := token.(TokenError); ok {
					t.Fatal(errToken)
				}

				if wToken, ok := token.(TokenWhitespace); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else if string(wToken.Representation()) != source {
					t.Fatal(fmt.Sprintf("unexpected token reresentation : '%s'", string(wToken.Representation())))
				} else {
					sawToken = true
				}
			}
		})
	}
}
