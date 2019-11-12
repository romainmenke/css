package tokenizer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTokenAtKeyword_OnlySelf(t *testing.T) {
	sources := map[string]string{
		`@media`:  `media`,
		`@screen`: `screen`,
	}

	for source, expected := range sources {
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

				if sToken, ok := token.(TokenAtKeyword); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else if sToken.String() != expected {
					t.Fatal(fmt.Sprintf("unexpected token string : %s", sToken.String()))
				} else if string(sToken.Representation()) != source {
					t.Fatal(fmt.Sprintf("unexpected token representation : %s", string(sToken.Representation())))
				} else {
					sawToken = true
				}
			}
		})
	}
}
