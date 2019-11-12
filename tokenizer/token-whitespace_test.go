package tokenizer

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestTokenWhitespace_OnlySelf(t *testing.T) {
	sources := map[string]string{
		string('\n'):                string('\n'),
		string('\n') + string('\r'): string('\n') + string('\n'),
		string('\r') + string('\n'): string('\n'),
		string('\r'):                string('\n'),
		string('\f'):                string('\n'),
		string(' '):                 string(' '),
		string('\t'):                string('\t'),
	}

	for source, expected := range sources {
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
				} else if string(wToken.Representation()) != expected {
					rep := ""
					for _, r := range wToken.Representation() {
						rep += fmt.Sprintf("%U", r)
					}

					in := ""
					for _, r := range source {
						in += fmt.Sprintf("%U", r)
					}

					t.Fatal(fmt.Sprintf("unexpected token representation : %s, input : %s", rep, in))
				} else {
					sawToken = true
				}
			}
		})
	}
}
