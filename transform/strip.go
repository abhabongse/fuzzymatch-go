package transform

import (
	"github.com/abhabongse/fuzzymatch-go/runedata"
	"golang.org/x/text/runes"
	"strings"
)

// StripNonPrintingTransform is a Unicode stream transformer object
// which removes all occurrences of non-printing and non-spacing rune characters
// from a string.
var StripNonPrintTransformer = runes.Remove(runes.NotIn(runedata.PrintsAndWhiteSpaces))

// ReSpace removes leading and trailing white-spaces,
// then it reduces all inter-word white-spaces into a single normal space.
//
// TODO: implement this function in terms of a string transformer
func ReSpace(str string) string {
	return strings.Join(strings.Fields(str), " ")
}
