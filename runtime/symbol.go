package runtime

import (
	"net/http"
)

type Symbol int

const (
	SymbolHTTP Symbol = iota
	SymbolCLI
)

func (s Symbol) Name() string {
	switch s {
	case SymbolHTTP:
		return "ServeHTTP"
	}

	return ""
}

func (s Symbol) MatchesInterface(i interface{}) bool {
	switch s {
	case SymbolHTTP:
		if v, ok := i.(func(http.ResponseWriter, *http.Request)); ok && v != nil {
			return true
		}
	}

	return false
}
