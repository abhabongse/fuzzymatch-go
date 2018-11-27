package fuzzymatch

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
)

/*
Remove leading and trailing white-spaces, then reduce all inter-word white-spaces
into a single normal space.
*/
func NormalizeWhiteSpaces(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

/*
Remove accents from the string as much as possible.
*/
func RemoveAccents(str string) string {
	isNonSpacingMark := runes.Remove(runes.In(AllCombiningDiacriticalMarks))
	transformer := transform.Chain(norm.NFKD, isNonSpacingMark, norm.NFKC)

	result, _, err := transform.String(transformer, str)
	if err == nil {
		return result
	}
	return str
}
