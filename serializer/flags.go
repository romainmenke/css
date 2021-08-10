package serializer

type Flag int

const (
	Noop Flag = 1 << iota
	Minify
)

func Flags(flags ...Flag) Flag {
	if len(flags) == 0 {
		return Noop
	}

	out := Noop
	for _, flag := range flags {
		out = out.Set(flag)
	}

	return out
}

func (f Flag) Set(flag Flag) Flag {
	return f | flag
}

func (f Flag) Clear(flag Flag) Flag  { return f &^ flag }
func (f Flag) Toggle(flag Flag) Flag { return f ^ flag }
func (f Flag) Has(flag Flag) bool    { return f&flag != 0 }
