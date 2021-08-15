package runtime

import (
	"bytes"
	"context"
	"testing"
)

func TestCompileGo(t *testing.T) {
	ctx := context.Background()
	runtime := New()
	defer runtime.Close()

	_, err := runtime.CompileGo(ctx, "serve-http-test", bytes.NewBufferString(testHTTPPlugin))
	if err != nil {
		t.Fatal(err)
	}
}
