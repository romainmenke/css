package stylesheet

// Function has a name and a value consisting of a list of component values.
type Function struct {
	Name  string
	Value []ComponentValue
}

func (t Function) String() string {
	return ""
}

func (t Function) Representation() []rune {
	return []rune{}
}
