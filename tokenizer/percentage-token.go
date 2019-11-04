package tokenizer

import "fmt"

type TokenPercentage struct {
	Value float64
}

func (t TokenPercentage) String() string {
	return fmt.Sprintf("%f", t.Value) + "%"
}
