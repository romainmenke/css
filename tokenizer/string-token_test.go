package tokenizer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTokenString_OnlySelf(t *testing.T) {
	sources := []string{
		`'foo'`,
		`'fo\o'`,
		`'foo\''`,
		`'foo'`,
		`"foo"`,
		`"foo\""`,
	}

	for _, source := range sources {
		t.Run(source, func(t *testing.T) {
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

				if sToken, ok := token.(TokenString); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else if sToken.String() != source {
					t.Fatal(fmt.Sprintf("unexpected token string : %s", sToken.String()))
				} else {
					sawToken = true
				}
			}
		})
	}
}

func TestTokenString_OnlySelf_NoNewLine(t *testing.T) {
	sources := []string{
		`'fo` + string('\r') + `o'`,
		`'fo` + string('\n') + `o'`,
		`'fo` + string('\r') + string('\n') + `o'`,
		`'fo` + string('\f') + `o'`,
	}

	for _, source := range sources {
		t.Run(source, func(t *testing.T) {
			tokenizer := New(bytes.NewBufferString(source))

			for {
				token := tokenizer.Next()
				if errToken, ok := token.(TokenError); ok {
					if errToken.Error() == "unexpected newline" {
						break
					}
				}

				if token != nil {
					t.Log(fmt.Sprintf("unexpected token of type : %T", token))
				}

				t.Fatal("expected 'unexpected newline' error")
			}
		})
	}
}
