package tokenizer

import "bufio"

type BufioRuneReader struct {
	*bufio.Reader
}

func (r *BufioRuneReader) PeekOneRune() (rune, error) {
	first, _, err := r.ReadRune()
	if err != nil {
		return first, err
	}

	err = r.UnreadRune()
	if err != nil {
		return first, err
	}

	return first, nil
}

func (r *BufioRuneReader) PeekTwoRunes() (rune, rune, error) {
	first, _, err := r.ReadRune()
	if err != nil {
		return first, 0, err
	}

	err = r.UnreadRune()
	if err != nil {
		return first, 0, err
	}

	second, _, err := r.ReadRune()
	if err != nil {
		return first, second, err
	}

	err = r.UnreadRune()
	if err != nil {
		return first, second, err
	}

	return first, second, nil
}

func (r *BufioRuneReader) PeekThreeRunes() (rune, rune, rune, error) {
	first, _, err := r.ReadRune()
	if err != nil {
		return first, 0, 0, err
	}

	err = r.UnreadRune()
	if err != nil {
		return first, 0, 0, err
	}

	second, _, err := r.ReadRune()
	if err != nil {
		return first, second, 0, err
	}

	err = r.UnreadRune()
	if err != nil {
		return first, second, 0, err
	}

	third, _, err := r.ReadRune()
	if err != nil {
		return first, second, third, err
	}

	err = r.UnreadRune()
	if err != nil {
		return first, second, third, err
	}

	return first, second, third, nil
}
