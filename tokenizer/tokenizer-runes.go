package tokenizer

import (
	"io"
)

type RuneReader interface {
	ReadRune() (rune, int, error)
	peekOneRune() (rune, error)
	peekTwoRunes() (rune, rune, error)
	peekThreeRunes() (rune, rune, rune, error)
}

func (t *Tokenizer) ReadRune() (rune, int, error) {
	return t.reader.ReadRune()
}

func (t *Tokenizer) representation() []rune {
	return t.reader.Representation()
}

func (t *Tokenizer) peekOneRune() (rune, error) {
	runes, _, err := t.reader.PeekRunes(1)
	if err != nil && err != io.EOF {
		return 0, err
	}

	if len(runes) == 0 {
		return 0, io.EOF
	}

	return runes[0], nil
}

func (t *Tokenizer) peekTwoRunes() (rune, rune, error) {
	runes, _, err := t.reader.PeekRunes(2)
	if err != nil && err != io.EOF {
		return 0, 0, err
	}

	switch len(runes) {
	case 0:
		return 0, 0, io.EOF
	case 1:
		return runes[0], 0, io.EOF
	default:
		return runes[0], runes[1], nil
	}
}

func (t *Tokenizer) peekThreeRunes() (rune, rune, rune, error) {
	runes, _, err := t.reader.PeekRunes(3)
	if err != nil && err != io.EOF {
		return 0, 0, 0, err
	}

	switch len(runes) {
	case 0:
		return 0, 0, 0, io.EOF
	case 1:
		return runes[0], 0, 0, io.EOF
	case 2:
		return runes[0], runes[1], 0, io.EOF
	default:
		return runes[0], runes[1], runes[2], nil
	}
}
