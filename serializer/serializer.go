package serializer

import (
	"io"
)

type Serializer interface {
	Serialize(w io.Writer, options Options) (int, error)
}

type Options struct {
	Flag  Flag
	Depth int
}

func Serialize(w io.Writer, list []interface{}, options Options) (int, error) {
	n := 0

	for _, el := range list {
		if s, ok := el.(Serializer); ok {
			nn, err := s.Serialize(w, options)
			n += nn
			if err != nil {
				return n, err
			}
		}
	}

	return n, nil
}
