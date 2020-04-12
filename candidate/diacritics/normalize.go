package diacritics

import (
	"github.com/abhabongse/fuzzymatch-go/runedata"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// StripDiacriticalMarksTransformer is a Unicode stream transformer object
// which tries to remove as many combining diacritical marks
// from the input string as possible.
// It handles various combinations of the same Unicode characters whenever possible
// (such as 'ö' as a single codepoint vs. 'o' + '¨' = 'ö' which has 2 codepoints).
//
// The removal process is preceded by Unicode decomposition,
// and the result is then re-combined to get final output.
var StripDiacriticalMarksTransformer = transform.Chain(
	norm.NFKD,
	// this cannot be unicode.Mn since it will also remove Thai non-space characters
	runes.Remove(runes.In(runedata.CombiningDiacriticalMarks)),
	norm.NFKC,
)
