package tokenizer

import (
	"io"
	"strconv"
	"unicode"
)

func ConsumeNumber(t *Tokenizer, initial rune) (interface{}, error) {

	isNumber := false

	if initial == '+' || initial == '-' {
		t.ReadRune() // reconsume
		t.tracking = append(t.tracking, initial)
	}

	// While the next input code point is a digit, consume it and append it to repr.
DIGITS_1:
	for {
		r, _, err := t.ReadRune()
		if err != nil {
			return nil, err
		}

		switch {
		case unicode.In(r, unicode.Digit):
			t.tracking = append(t.tracking, r)
			err = ConsumeDigits(t)
			if err == io.EOF {
				return parseNumber(t, isNumber)
			}
			if err != nil {
				return nil, err
			}

			break DIGITS_1
		}
	}

	p1, p2, err := t.PeekTwoRunes()
	if err == io.EOF {
		return parseNumber(t, isNumber)
	}
	if err != nil {
		return nil, err
	}

	if p1 == '.' && unicode.In(p2, unicode.Digit) {
		isNumber = true

		{ // dot
			r, _, err := t.ReadRune()
			if err == io.EOF {
				return nil, io.ErrUnexpectedEOF
			}
			if err != nil {
				return nil, err
			}

			t.tracking = append(t.tracking, r)
		}

		{ // digits
			err := ConsumeDigits(t)
			if err == io.EOF {
				return parseNumber(t, isNumber)
			}

			if err != nil {
				return nil, err
			}
		}
	}

	p1, p2, p3, _ := t.PeekThreeRunes()
	p1IsE := p1 == 'e' || p1 == 'E'
	p2IsSign := p2 == '-' || p2 == '+'
	p2IsDigit := unicode.In(p2, unicode.Digit)
	p3IsDigit := unicode.In(p3, unicode.Digit)
	if (p1IsE && p2IsSign && p3IsDigit) || (p1IsE && p2IsDigit) {
		isNumber = true

		{ // e
			r, _, err := t.ReadRune()
			if err == io.EOF {
				return parseNumber(t, isNumber)
			}
			if err != nil {
				return nil, err
			}

			t.tracking = append(t.tracking, r)
		}

		if p2IsSign {
			r, _, err := t.ReadRune()
			if err == io.EOF {
				return parseNumber(t, isNumber)
			}
			if err != nil {
				return nil, err
			}

			t.tracking = append(t.tracking, r)
		}

		{ // digits
			err := ConsumeDigits(t)
			if err == io.EOF {
				return parseNumber(t, isNumber)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	return parseNumber(t, isNumber)
}

func ConsumeDigits(t *Tokenizer) error {
	for {
		peeked, err := t.PeekOneRune()
		if err != nil {
			return err
		}

		if unicode.In(peeked, unicode.Digit) {
			r, _, err := t.ReadRune()
			if err != nil {
				return err
			}

			t.tracking = append(t.tracking, r)
		} else {
			return nil
		}
	}
}

func parseNumber(t *Tokenizer, isNumber bool) (interface{}, error) {
	if isNumber {
		value, err := strconv.ParseFloat(string(t.tracking), 64)
		if err != nil {
			return nil, err
		}

		return value, nil
	}

	value, err := strconv.ParseInt(string(t.tracking), 10, 64)
	if err != nil {
		return nil, err
	}

	return value, nil
}
