package tokenizer

import (
	"io"
)

func ConsumeNumeric(t *Tokenizer, initial rune) Token {
	v, err := ConsumeNumber(t, initial)
	if err == io.EOF {
		return TokenEOF{}
	}
	if err != nil {
		return TokenError{error: err}
	}

	if CheckIfThreeCodePointsWouldStartAnIdentifier(t) {
		unit, err := ConsumeName(t)
		if err != nil {
			return TokenError{error: err}
		}

		switch vv := v.(type) {
		case float64:
			return NewTokenDimensionFloat64(vv, unit, t.Representation())
		case int64:
			return NewTokenDimensionInt64(vv, unit, t.Representation())
		case int:
			return NewTokenDimensionInt(vv, unit, t.Representation())
		}
	}

	peeked, err := t.PeekOneRune()
	if err != nil && err != io.EOF {
		return TokenError{error: err}
	}

	if peeked == '%' {
		t.ReadRune()

		switch vv := v.(type) {
		case float64:
			return NewTokenPercentageFloat64(vv, t.Representation())
		case int64:
			return NewTokenPercentageInt64(vv, t.Representation())
		case int:
			return NewTokenPercentageInt(vv, t.Representation())
		}
	}

	switch vv := v.(type) {
	case float64:
		return NewTokenNumberFloat64(vv, t.Representation())
	case int64:
		return NewTokenNumberInt64(vv, t.Representation())
	case int:
		return NewTokenNumberInt(vv, t.Representation())
	}

	return NewTokenNumberInt(0, t.Representation())
}
