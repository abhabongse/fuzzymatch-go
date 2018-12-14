package normalization

import (
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/runedata"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

/*
ApplyTransformers applies each string transformer in the given sequence of
transformers to the given input string. If any transformer produces an error,
it will be silently ignored and intermediate string will not be affected.
*/
func ApplyTransformers(str string, ts ...transform.Transformer) string {
	for _, t := range ts {
		modifiedStr, _, err := transform.String(t, str)
		if err == nil {
			str = modifiedStr
		}
	}
	return str
}

/*
StripNonPrintingTransform is a Unicode stream transformer object which removes
all occurrences of non-printing and non-spacing rune characters from a string.
*/
var StripNonPrintTransformer = runes.Remove(runes.NotIn(runedata.PrintsAndWhiteSpaces))

/*
ToNormalSpaceTransformer is a Unicode stream transformer object which replaces
all white space rune characters into a normal space.
*/
var ToNormalSpaceTransformer = runes.If(
	runes.In(unicode.White_Space),
	runes.Map(func(r rune) rune { return ' ' }),
	nil,
)

/*
RemoveAccentsTransformer is a Unicode stream transformer object which tries to
removes as many combining diacritical marks from the input string as possible.
It handles various combinations of the same Unicode characters whenever possible
(such as 'ö' as a single codepoint vs. 'o' + '¨' = 'ö' which has 2 codepoints).

The removal process is preceded by Unicode decomposition, and the result is
then re-combined to get final output.
*/
var RemoveAccentsTransformer = transform.Chain(
	norm.NFKD,
	runes.Remove(runes.In(runedata.AllCombiningDiacriticalMarks)),
	norm.NFKC,
)

/*
ToLowerTransformer is a Unicode stream transformer object which transforms all
unicode characters into its lowercase forms as defined by Unicode property.
*/
var ToLowerTransformer = runes.Map(unicode.ToLower)
