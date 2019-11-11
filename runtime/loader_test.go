package runtime

import (
	"bytes"
	"context"
	"testing"
)

func TestLoad(t *testing.T) {
	ctx := context.Background()
	runtime := New()
	defer runtime.Close()

	plugin, err := runtime.CompileGo(ctx, "serve-http-test", bytes.NewBufferString(testHTTPPlugin))
	if err != nil {
		t.Fatal(err)
	}

	_, err = runtime.Load(ctx, plugin, SymbolHTTP)
	if err != nil {
		t.Fatal(err)
	}
}

var testHTTPPlugin = `package main

import "net/http"

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
`
