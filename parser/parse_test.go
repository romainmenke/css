package parser

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/romainmenke/css/tokenizer"
)

func TestParseStylesheet(t *testing.T) {
	src := `body {
	background-color: var(--color-background);
	box-sizing: border-box;
	color: var(--color-foreground);
	min-height: 100vh;
}

@media (screen) {
	body {
		min-height: 80vh;
	}
}

main:not([class]) {
	display: block;
	width: 100%;
}

textarea {
	resize: none;
}

hr {
	background-color: currentColor;
	border: none;
	height: 1px;
}

img,
video {
	display: block;
	max-width: 100%;
}

table,
td {
	border: 1px solid var(--color-foreground-light);
}

`

	tz := tokenizer.New(bytes.NewBufferString(src))
	for {
		t := tz.Next()
		if _, ok := t.(tokenizer.TokenEOF); ok {
			break
		}
	}

	parser := New(bytes.NewBufferString(src))
	s, err := parser.ParseStylesheet()
	if err != nil {
		t.Fatal(err)
	}

	for _, rule := range s.Rules {
		t.Log(fmt.Sprintf("rule of type : %T", rule))
	}
}
