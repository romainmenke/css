package parser

import (
	"bytes"
	"testing"

	"github.com/romainmenke/css/serializer"
)

func Test_Serialize(t *testing.T) {
	parser := New(bytes.NewBufferString(src))
	s, err := parser.ParseStylesheet()
	if err != nil {
		t.Fatal(err)
	}

	serializedBuff := bytes.NewBuffer(nil)

	_, err = serializer.Serialize(serializedBuff, s.Rules, serializer.Options{Flag: 0, Depth: 0})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(serializedBuff.String())
}

func Test_Serialize_Minify(t *testing.T) {
	parser := New(bytes.NewBufferString(src))
	s, err := parser.ParseStylesheet()
	if err != nil {
		t.Fatal(err)
	}

	serializedBuff := bytes.NewBuffer(nil)

	_, err = serializer.Serialize(serializedBuff, s.Rules, serializer.Options{Flag: serializer.Minify, Depth: 0})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(serializedBuff.String())
}

const src = `
@media (min-width: 1024px) {

	/* styles for body */
	body {
		width: 100%;
	}
}
`
