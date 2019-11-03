package tokenizer

import "fmt"

type PercentageToken struct {
	Value float64
}

func (t PercentageToken) String() string {
	return fmt.Sprintf("%f", t.Value) + "%"
}
