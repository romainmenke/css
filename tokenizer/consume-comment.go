package tokenizer

func consumeComment(t *Tokenizer) Token {
	open, err := t.peekOneRune()
	if err != nil {
		return TokenError{error: err}
	}

	if open != '*' {
		return nil
	}

	_, _, err = t.ReadRune()
	if err != nil {
		return TokenError{error: err}
	}

	for {
		r, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		if r == '*' {
			close, err := t.peekOneRune()
			if err != nil {
				return TokenError{error: err}
			}

			if close == '/' {
				t.ReadRune()

				return TokenComment{
					Value:          append([]rune(nil), t.tracking...),
					representation: append([]rune(nil), t.representation()...),
				}
			}
		}

		t.tracking = append(t.tracking, r)
	}
}
