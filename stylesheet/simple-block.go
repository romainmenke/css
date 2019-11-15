package stylesheet

import "github.com/romainmenke/css/tokenizer"

type Block interface{}

// SimpleBlock has an associated token (either a <[-token>, <(-token>, or <{-token>) and a value consisting of a list of component values.
type SimpleBlock struct {
	AssociatedToken tokenizer.Token
	Value           []interface{}
}

func (t SimpleBlock) String() string {
	return ""
}
