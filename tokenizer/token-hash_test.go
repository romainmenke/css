package tokenizer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTokenHash_OnlySelf(t *testing.T) {
	sources := map[string]string{
		`#foo`:    `foo`,
		`#f\26 o`: `f&o`,
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

				if hToken, ok := token.(TokenHash); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else if hToken.String() != expected {
					t.Fatal(fmt.Sprintf("unexpected token string : %s", hToken.String()))
				} else if string(hToken.Representation()) != source {
					t.Fatal(fmt.Sprintf("unexpected token representation : %s", string(hToken.Representation())))
				} else if hToken.Type != HashTokenTypeID {
					t.Fatal(fmt.Sprintf("unexpected token type : %v", hToken.Type))
				} else {
					sawToken = true
				}
			}
		})
	}
}

func TestTokenHash_WithTypeID_OnlySelf(t *testing.T) {
	sources := map[string]string{
		`#-foo`:  `-foo`,
		`#--foo`: `--foo`,
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

				if hToken, ok := token.(TokenHash); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else if hToken.String() != expected {
					t.Fatal(fmt.Sprintf("unexpected token string : %s", hToken.String()))
				} else if string(hToken.Representation()) != source {
					t.Fatal(fmt.Sprintf("unexpected token representation : %s", string(hToken.Representation())))
				} else if hToken.Type != HashTokenTypeID {
					t.Fatal(fmt.Sprintf("unexpected token type : %v", hToken.Type))
				} else {
					sawToken = true
				}
			}
		})
	}
}
