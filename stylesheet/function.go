package stylesheet

// Function has a name and a value consisting of a list of component values.
type Function struct {
	Name  string
	Value []interface{}
}

func (t Function) String() string {
	return ""
}
