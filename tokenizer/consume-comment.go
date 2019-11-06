package tokenizer

func ConsumeComment(t *Tokenizer) Token {
	open, _, err := t.ReadRune()
	if err != nil {
		return TokenError{error: err}
	}

	if open != '*' {
		err := t.UnreadRune()
		if err != nil {
			return TokenError{error: err}
		}

		return nil
	}

	for {
		r, _, err := t.ReadRune()
		if err != nil {
			return TokenError{error: err}
		}

		if r == '*' {
			close, _, err := t.ReadRune()
			if err != nil {
				return TokenError{error: err}
			}

			if close == '/' {
				return TokenComment{
					Value:         t.tracking,
					represenation: t.representation,
				}
			}

			err = t.UnreadRune()
			if err != nil {
				return TokenError{error: err}
			}
		}

		t.tracking = append(t.tracking, r)
	}
}
