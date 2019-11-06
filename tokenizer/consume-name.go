package tokenizer

import "unicode"

func ConsumeName(t *Tokenizer) ([]rune, error) {
	name := make([]rune, 0, 1000)

	for {
		r, _, err := t.ReadRune()
		if err != nil {
			return nil, err
		}

		switch {
		case CheckIfTwoCodePointsAreAValidEscape(t, r):
			unescaped, err := Unescape(t, r)
			if err != nil {
				return nil, err
			}

			name = append(name, unescaped)

		case unicode.In(r, NameCodePoint...):
			name = append(name, r)
		default:
			err := t.UnreadRune()
			if err != nil {
				return nil, err
			}

			return name, nil
		}
	}
}
