package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/romainmenke/css/tokenizer"
)

func TestMain(t *testing.T) {
	src := `@context (http) {
		server {
			port: env(port);
			timeout: 30s;
		}
	}`

	tz := tokenizer.New(bytes.NewBufferString(src))
	for {
		token := tz.Next()

		t.Log(fmt.Sprintf("token of type : %T", token))
		t.Log(token.String())

		if _, ok := token.(tokenizer.TokenEOF); ok {
			break
		}

		if _, ok := token.(tokenizer.TokenError); ok {
			break
		}
	}
}

func TestEndpoint_A(t *testing.T) {
	src := `@context (http) {
		response::body {
			content: 'hello ' db('SELECT name FROM users ORDER BY rand LIMIT 1');
		}

		response::headers {
			content-type: 'text/plain';
		}
	}`

	tz := tokenizer.New(bytes.NewBufferString(src))
	for {
		token := tz.Next()

		t.Log(fmt.Sprintf("token of type : %T", token))
		t.Log(token.String())

		if _, ok := token.(tokenizer.TokenEOF); ok {
			break
		}

		if _, ok := token.(tokenizer.TokenError); ok {
			break
		}
	}
}
