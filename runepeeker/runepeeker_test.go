package runepeeker

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPreProcessor(t *testing.T) {
	source := "baz"

	peeker := New(bufio.NewReader(bytes.NewBufferString(source)))

	peeked, _, err := peeker.PeekRunes(3)
	if err != nil {
		t.Fatal(err)
	}

	if len(peeked) != 3 {
		t.Fatal("did not peek 3 runes")
	}

	if peeked[0] != 'b' {
		t.Fatalf("unepexted rune '%U'\n", peeked[0])
	}

	if peeked[1] != 'a' {
		t.Fatalf("unepexted rune '%U'\n", peeked[1])
	}

	if peeked[2] != 'z' {
		t.Fatalf("unepexted rune '%U'\n", peeked[2])
	}

	peeked, _, err = peeker.PeekRunes(1)
	if err != nil {
		t.Fatal(err)
	}

	if peeked[0] != 'b' {
		t.Fatalf("unepexted rune '%U'\n", peeked[0])
	}

	read, _, err := peeker.ReadRune()
	if err != nil {
		t.Fatal(err)
	}

	if read != 'b' {
		t.Fatalf("unepexted rune '%U'\n", read)
	}

	read, _, err = peeker.ReadRune()
	if err != nil {
		t.Fatal(err)
	}

	if read != 'a' {
		t.Fatalf("unepexted rune '%U'\n", read)
	}

	read, _, err = peeker.ReadRune()
	if err != nil {
		t.Fatal(err)
	}

	if read != 'z' {
		t.Fatalf("unepexted rune '%U'\n", read)
	}
}

func TestPreProcessor_CR(t *testing.T) {
	source := "b" + string('\u000d') + string('\u000a') + "z"

	peeker := New(bufio.NewReader(bytes.NewBufferString(source)))

	peeked, _, err := peeker.PeekRunes(3)
	if err != nil {
		t.Fatal(err)
	}

	if len(peeked) != 3 {
		t.Fatal("did not peek 3 runes")
	}

	if peeked[0] != 'b' {
		t.Fatalf("unepexted rune '%U'\n", peeked[0])
	}

	if peeked[1] != '\u000a' {
		t.Fatalf("unepexted rune '%U'\n", peeked[1])
	}

	if peeked[2] != 'z' {
		t.Fatalf("unepexted rune '%U'\n", peeked[2])
	}

	read, _, err := peeker.ReadRune()
	if err != nil {
		t.Fatal(err)
	}

	if read != 'b' {
		t.Fatalf("unepexted rune '%U'\n", read)
	}

	read, _, err = peeker.ReadRune()
	if err != nil {
		t.Fatal(err)
	}

	if read != '\u000a' {
		t.Fatalf("unepexted rune '%U'\n", read)
	}

	read, _, err = peeker.ReadRune()
	if err != nil {
		t.Fatal(err)
	}

	if read != 'z' {
		t.Fatalf("unepexted rune '%U'\n", read)
	}
}
