package tokenizer

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"testing"
)

func TestWhitespaceToken_OnlySelf(t *testing.T) {
	sources := []string{
		string('\n'),
		string('\n') + string('\r'),
		string('\r'),
		string('\f'),
		string(' '),
		string('\t'),
	}

	for _, source := range sources {
		t.Run(strconv.Quote(source), func(t *testing.T) {
			tokenizer := New(bytes.NewBufferString(source))
			sawWhiteSpaceToken := false

			for {
				token, err := tokenizer.Next()
				if err == io.EOF && sawWhiteSpaceToken {
					break
				}
				if err != nil {
					t.Fatal(err)
				}

				if _, ok := token.(WhitespaceToken); !ok {
					t.Fatal(fmt.Sprintf("unexpected token of type : %T", token))
				} else {
					sawWhiteSpaceToken = true
				}
			}
		})
	}
}
