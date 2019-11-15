package stylesheet

type Rule interface {
}

// AtRule has a name, a prelude consisting of a list of component values, and an optional block consisting of a simple {} block.
type AtRule struct {
	Block   Block
	Name    string
	Prelude []interface{}
}

func (t AtRule) String() string {
	return ""
}

// QualifiedRule has a prelude consisting of a list of component values, and a block consisting of a simple {} block.
type QualifiedRule struct {
	Block   Block
	Prelude []interface{}
}

func (t QualifiedRule) String() string {
	return ""
}

type RuleList []interface{}
