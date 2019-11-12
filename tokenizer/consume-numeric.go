package tokenizer

import (
	"io"
)

func consumeNumeric(t *Tokenizer, initial rune) Token {
	v, err := consumeNumber(t, initial)
	if err == io.EOF {
		return TokenEOF{}
	}
	if err != nil {
		return TokenError{error: err}
	}

	if checkIfThreeCodePointsWouldStartAnIdentifier(t) {
		unit, err := consumeName(t)
		if err != nil {
			return TokenError{error: err}
		}

		switch vv := v.(type) {
		case float64:
			return NewTokenDimensionFloat64(vv, unit, t.representation())
		case int64:
			return NewTokenDimensionInt64(vv, unit, t.representation())
		case int:
			return NewTokenDimensionInt(vv, unit, t.representation())
		}
	}

	peeked, err := t.peekOneRune()
	if err != nil && err != io.EOF {
		return TokenError{error: err}
	}

	if peeked == '%' {
		t.ReadRune()

		switch vv := v.(type) {
		case float64:
			return NewTokenPercentageFloat64(vv, t.representation())
		case int64:
			return NewTokenPercentageInt64(vv, t.representation())
		case int:
			return NewTokenPercentageInt(vv, t.representation())
		}
	}

	switch vv := v.(type) {
	case float64:
		return NewTokenNumberFloat64(vv, t.representation())
	case int64:
		return NewTokenNumberInt64(vv, t.representation())
	case int:
		return NewTokenNumberInt(vv, t.representation())
	}

	return NewTokenNumberInt(0, t.representation())
}
