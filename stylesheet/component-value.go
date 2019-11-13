package stylesheet

// ComponentValue is one of the preserved tokens, a function, or a simple block.
type ComponentValue interface {
	String() string
	Representation() []rune
}
