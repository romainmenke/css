package stylesheet

// Declaration has a name, a value consisting of a list of component values, and an important flag which is initially unset.
// Declarations are further categorized as "properties" or "descriptors", with the former typically appearing in qualified rules and the latter appearing in at-rules.
// (This categorization does not occur at the Syntax level; instead, it is a product of where the declaration appears, and is defined by the respective specifications defining the given rule.)
type Declaration struct {
	Name      string
	Value     []ComponentValue
	Important bool
}

func (t Declaration) String() string {
	return ""
}

func (t Declaration) Representation() []rune {
	return []rune{}
}
