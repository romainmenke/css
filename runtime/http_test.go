package runtime

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTP(t *testing.T) {
	ctx := context.Background()
	runtime := New()
	defer runtime.Close()

	plugin, err := runtime.CompileGo(ctx, "serve-http-test", bytes.NewBufferString(testHTTPPlugin))
	if err != nil {
		t.Fatal(err)
	}

	handler, err := runtime.LoadHTTP(ctx, plugin)
	if err != nil {
		t.Fatal(err)
	}

	testServer := httptest.NewServer(handler)
	resp, err := http.Get(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != "hello world!" {
		t.Fatal("unexpected response body")
	}
}

func BenchmarkHTTP(b *testing.B) {
	ctx := context.Background()
	runtime := New()
	defer runtime.Close()

	_, err := runtime.CompileGo(ctx, "serve-http-test", bytes.NewBufferString(testHTTPPlugin))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		plugin, err := runtime.CompileGo(ctx, "serve-http-test", bytes.NewBufferString(testHTTPPlugin))
		if err != nil {
			b.Fatal(err)
		}

		handler, err := runtime.LoadHTTP(ctx, plugin)
		if err != nil {
			b.Fatal(err)
		}

		testServer := httptest.NewServer(handler)
		resp, err := http.Get(testServer.URL)
		if err != nil {
			b.Fatal(err)
		}

		defer resp.Body.Close()
		respB, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			b.Fatal(err)
		}

		if string(respB) != "hello world!" {
			b.Fatal("unexpected response body")
		}

		testServer.Close()
	}
}
