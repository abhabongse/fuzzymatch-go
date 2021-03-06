// This source file contains additional definitions of ranges of Unicode characters.
// Characters within the same range would share some particular properties or traits.

package runedata

import (
	"golang.org/x/text/unicode/rangetable"
	"unicode"
)

// CombiningDiacriticalMarks is a collection of all combining critical marks
// defined in Unicode standard.
var CombiningDiacriticalMarks = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0300, 0x036f, 1}, // Combining Diacritical Marks
		{0x1ab0, 0x1aff, 1}, // Combining Diacritical Marks Extended
		{0x1dc0, 0x1dff, 1}, // Combining Diacritical Marks Supplement
		{0x20d0, 0x20ff, 1}, // Combining Diacritical Marks for Symbols
	},
}

// PrintsAndWhiteSpaces is a collection of all printable characters
// (i.e. those with Unicode category L, M, N, P, S)
// as well as all kinds of white spaces.
var PrintsAndWhiteSpaces = rangetable.Merge(
	unicode.White_Space,
	unicode.Letter,
	unicode.Mark,
	unicode.Number,
	unicode.Punct,
	unicode.Symbol,
)

// HyphensAndDashes is a collection of symbols looking alike horizontal bar.
var HyphensAndDashes = rangetable.Merge(
	unicode.Hyphen,
	unicode.Dash,
)
