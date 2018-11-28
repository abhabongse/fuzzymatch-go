package fuzzymatch

import "unicode"

/*
A collection of combining diacritical marks in Unicode.
*/
var AllCombiningDiacriticalMarks = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0300, 0x036f, 1}, // Combining Diacritical Marks
		{0x1ab0, 0x1aff, 1}, // Combining Diacritical Marks Extended
		{0x1dc0, 0x1dff, 1}, // Combining Diacritical Marks Supplement
		{0x20d0, 0x20ff, 1}, // Combining Diacritical Marks for Symbols
		{0xfe20, 0xfe2f, 1}, // Combining Half Marks
	},
}
