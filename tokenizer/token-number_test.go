package tokenizer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTokenNumber_Int_OnlySelf(t *testing.T) {
	sources := map[string]int64{
		`+6`:    6,
		`+1000`: 1000,
		`-6`:    -6,
		`-1000`: -1000,
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

				if sToken, ok := token.(TokenNumber); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else if sToken.IntValue() != expected {
					t.Fatal(fmt.Sprintf("unexpected token string : %d", sToken.IntValue()))
				} else if string(sToken.Representation()) != source {
					t.Fatal(fmt.Sprintf("unexpected token representation : %s", string(sToken.Representation())))
				} else {
					sawToken = true
				}
			}
		})
	}
}

func TestTokenNumber_Float_OnlySelf(t *testing.T) {
	sources := map[string]float64{
		`+6.123`:   6.123,
		`+230.123`: 230.123,
		`-6.123`:   -6.123,
		`-230.123`: -230.123,
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

				if sToken, ok := token.(TokenNumber); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else if sToken.FloatValue() != expected {
					t.Fatal(fmt.Sprintf("unexpected token string : %f", sToken.FloatValue()))
				} else if string(sToken.Representation()) != source {
					t.Fatal(fmt.Sprintf("unexpected token representation : %s", string(sToken.Representation())))
				} else {
					sawToken = true
				}
			}
		})
	}
}
