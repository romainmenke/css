package tokenizer

type TokenComment struct {
	Value         []rune
	represenation []rune
}

func (t TokenComment) String() string {
	return string(t.Value)
}

func (t TokenComment) Representation() []rune {
	return t.represenation
}

func TokenizeComment(t *Tokenizer) Token {
	rOpen, _, err := t.ReadRune()
	if err != nil {
		return TokenError{error: err}
	}

	if rOpen != '*' {
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
			rClose, _, err := t.ReadRune()
			if err != nil {
				return TokenError{error: err}
			}

			if rClose == '/' {
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
