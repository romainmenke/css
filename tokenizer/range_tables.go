package tokenizer

import (
	"math"
	"unicode"
)

// https://drafts.csswg.org/css-syntax-3/#tokenizer-definitions
// Upper unicode.Upper
// Lower  unicode.Lower
// Hex Digit unicode.Hex_Digit
// Letter unicode.Letter

// r > unicode.MaxASCII
var NonASCII = &unicode.RangeTable{
	R16: []unicode.Range16{
		{Lo: unicode.MaxASCII + 1, Hi: math.MaxUint16, Stride: 1},
	},
	R32: []unicode.Range32{
		{Lo: 0x10000, Hi: unicode.MaxRune, Stride: 1},
	},
}

var NameStartCodePoint = []*unicode.RangeTable{
	unicode.Letter,
	NonASCII,
	&unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: '\u005f', Hi: '\u005f', Stride: 1},
		},
	},
}

var NameCodePoint = append(
	NameStartCodePoint,
	unicode.Digit,
	&unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: '\u002d', Hi: '\u002d', Stride: 1},
		},
	},
)
