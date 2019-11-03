package tokenizer

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestStringToken_OnlySelf(t *testing.T) {
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
				token, err := tokenizer.Next()
				if err == io.EOF && sawToken {
					break
				}
				if err != nil {
					t.Fatal(err)
				}

				if sToken, ok := token.(StringToken); !ok {
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

func TestStringToken_OnlySelf_NoNewLine(t *testing.T) {
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
				token, err := tokenizer.Next()
				if err == nil {
					if token != nil {
						t.Log(fmt.Sprintf("unexpected token of type : %T", token))
					}
					t.Fatal("expected 'unexpected newline' error")
				}

				if err.Error() == "unexpected newline" {
					break
				}

				t.Fatal(err)
			}
		})
	}
}
