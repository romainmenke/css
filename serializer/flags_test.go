package serializer

import "testing"

func Test_FlagsMinify(t *testing.T) {
	if Minify.Has(Minify) {
		return
	}

	t.Fatal("expected true")
}
