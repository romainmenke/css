package stylesheet

import (
	"io"
	"strings"

	"github.com/romainmenke/css/serializer"
	"github.com/romainmenke/css/tokenizer"
)

// Declaration has a name, a value consisting of a list of component values, and an important flag which is initially unset.
// Declarations are further categorized as "properties" or "descriptors", with the former typically appearing in qualified rules and the latter appearing in at-rules.
// (This categorization does not occur at the Syntax level; instead, it is a product of where the declaration appears, and is defined by the respective specifications defining the given rule.)
type Declaration struct {
	Name      string
	Value     []interface{}
	Important bool
}

func (d Declaration) String() string {
	return ""
}

func (d *Declaration) Serialize(w io.Writer, options serializer.Options) (int, error) {
	var (
		n int
	)

	nn, err := w.Write([]byte(d.Name))
	n += nn
	if err != nil {
		return n, err
	}

	nn, err = w.Write([]byte(": "))
	n += nn
	if err != nil {
		return n, err
	}

	if !options.Flag.Has(serializer.Minify) {
		nn, err = w.Write([]byte(" "))
		n += nn
		if err != nil {
			return n, err
		}
	}

	for i, v := range d.Value {
		switch vv := v.(type) {
		case serializer.Serializer:
			nn, err = vv.Serialize(w, options)
			n += nn
			if err != nil {
				return n, err
			}

			if (i + 1) != len(d.Value) {
				nn, err = w.Write([]byte(" "))
				n += nn
				if err != nil {
					return n, err
				}
			}
		}
	}

	if d.Important {
		nn, err = w.Write([]byte(" !important"))
		n += nn
		if err != nil {
			return n, err
		}
	}

	nn, err = w.Write([]byte(";"))
	n += nn
	if err != nil {
		return n, err
	}

	return n, nil
}

func (d *Declaration) SetImportant() {
	if d == nil {
		return
	}

	if len(d.Value) < 2 {
		return
	}

	expectWhitespace := 2
	expectImportant := true
	expectExclamation := false
	removeLen := 0
	isImportant := true

	for counter := 0; counter < len(d.Value); counter++ {
		i := len(d.Value) - (1 + counter)
		token := d.Value[i]

		switch tt := token.(type) {
		case tokenizer.TokenWhitespace:
			if expectWhitespace > 0 {
				expectWhitespace--
				removeLen++
				continue
			} else {
				break
			}

		case tokenizer.TokenIdent:
			if expectImportant && strings.ToLower(string(tt.Value)) == "important" {
				expectWhitespace = 0
				expectImportant = false
				expectExclamation = true
				removeLen++
				continue
			} else {
				break
			}

		case tokenizer.TokenDelim:
			if expectExclamation && tt.Value == '!' {
				removeLen++
				isImportant = true
				break
			}

			break

		default:
			break
		}
	}

	if isImportant {
		d.Important = true
		d.Value = d.Value[:len(d.Value)-removeLen]
	}

	return
}

func (d *Declaration) RemoveTrailingWhitespace() {
	if d == nil {
		return
	}

	if len(d.Value) == 0 {
		return
	}

	for counter := 0; counter < len(d.Value); counter++ {
		i := len(d.Value) - (1 + counter)
		token := d.Value[i]

		switch token.(type) {
		case tokenizer.TokenWhitespace:
			d.Value = d.Value[:len(d.Value)-1]
			continue

		default:
			break
		}
	}

	return
}

type DeclarationList []interface{}
