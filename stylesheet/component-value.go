package stylesheet

import (
	"io"

	"github.com/romainmenke/css/serializer"
)

// ComponentValue is one of the preserved tokens, a function, or a simple block.
type ComponentValue interface {
	String() string
	Serialize(w io.Writer, options serializer.Options) (int, error)
}
