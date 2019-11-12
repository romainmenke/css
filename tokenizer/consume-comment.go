package tokenizer

func ConsumeComment(t *Tokenizer) Token {
	open, err := t.PeekOneRune()
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
			close, err := t.PeekOneRune()
			if err != nil {
				return TokenError{error: err}
			}

			if close == '/' {
				return TokenComment{
					Value:         t.tracking,
					representation: t.Representation(),
				}
			}
		}

		t.tracking = append(t.tracking, r)
	}
}
