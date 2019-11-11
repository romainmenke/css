package runtime

import (
	"context"
	"errors"
	"net/http"
	"plugin"
)

func (r *Runtime) Load(ctx context.Context, name string, symbol Symbol) (interface{}, error) {
	plug, err := plugin.Open(name)
	if err != nil {
		return nil, err
	}

	sym, err := plug.Lookup(symbol.Name())
	if err != nil {
		return nil, err
	}

	if symbol.MatchesInterface(sym) {
		return sym, nil
	}

	return nil, errors.New("incorrect symbol interface")
}

func (r *Runtime) LoadHTTP(ctx context.Context, name string) (http.HandlerFunc, error) {
	plug, err := plugin.Open(name)
	if err != nil {
		return nil, err
	}

	sym, err := plug.Lookup(SymbolHTTP.Name())
	if err != nil {
		return nil, err
	}

	if handler, ok := sym.(func(http.ResponseWriter, *http.Request)); ok && handler != nil {
		return http.HandlerFunc(handler), nil
	}

	return nil, errors.New("incorrect symbol interface")
}
