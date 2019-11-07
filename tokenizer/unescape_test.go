package tokenizer

import (
	"bytes"
	"strconv"
	"testing"
)

func TestUnescape(t *testing.T) {
	sources := map[string]rune{
		`\26`:   '&',
		`\26 B`: '&',
		`\\`:    '\\',
		`\{`:    '{',
	}

	for escaped, expected := range sources {
		t.Run(strconv.Quote(escaped), func(t *testing.T) {
			reader := New(bytes.NewBufferString(escaped))
			r, _, err := reader.ReadRune()
			if err != nil {
				t.Fatal(err)
			}

			unescaped, err := Unescape(reader, r)
			if err != nil {
				t.Fatal(err)
			}

			if unescaped != expected {
				t.Fatal("expected : ", string(expected), "got : ", string(unescaped))
			}
		})
	}
}
