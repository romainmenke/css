package stylesheet

type Rule interface {
}

// AtRule has a name, a prelude consisting of a list of component values, and an optional block consisting of a simple {} block.
type AtRule struct {
	Block   Block
	Name    string
	Prelude []ComponentValue
}

func (t AtRule) String() string {
	return ""
}

func (t AtRule) Representation() []rune {
	return []rune{}
}

// QualifiedRule has a prelude consisting of a list of component values, and a block consisting of a simple {} block.
type QualifiedRule struct {
	Block   Block
	Prelude []ComponentValue
}

func (t QualifiedRule) String() string {
	return ""
}

func (t QualifiedRule) Representation() []rune {
	return []rune{}
}
