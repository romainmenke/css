package tokenizer

import (
	"io"
	"unicode"
)

func ConsumeName(t *Tokenizer) ([]rune, error) {
	name := make([]rune, 0, 1000)

	for {
		if CheckIfTwoCodePointsAreAValidEscape(t) {
			r, _, err := t.ReadRune()
			if err == io.EOF {
				return name, nil
			}
			if err != nil {
				return nil, err
			}

			unescaped, err := Unescape(t, r)
			if err != nil {
				return nil, err
			}

			name = append(name, unescaped)
			continue
		}

		r, _, err := t.ReadRune()
		if err == io.EOF {
			return name, nil
		}
		if err != nil {
			return nil, err
		}

		if unicode.In(r, NameCodePoint...) {
			name = append(name, r)
			continue
		}

		err = t.UnreadRune()
		if err != nil {
			return nil, err
		}

		return name, nil
	}
}
