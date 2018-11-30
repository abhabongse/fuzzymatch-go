/*
Package normalization contains a collection of string normalization functions; they may
be used to pre-process input strings.
*/
package normalization

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

/*
NormalizeWhiteSpaces removes leading and trailing white-spaces, then it reduces all
inter-word white-spaces into a single normal space.
*/
func NormalizeWhiteSpaces(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

/*
RemoveAccents tries to remove as many combining diacritical marks from the input string
as possible. It handles various combinations of the same Unicode characters whenever
possible (such as 'ö' as a single codepoint vs. 'o' + '¨' = 'ö' which has 2 codepoints).
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

/*
RemoveNonPrinting will remove all non-printing characters as well as all white-space
characters that are not normal spaces (U+0020).
*/
func RemoveNonPrinting(str string) string {
	isNonPrinting := runes.Remove(runes.Predicate(func(r rune) bool { return !unicode.IsPrint(r)}))
	result, _, err := transform.String(isNonPrinting, str)
	if err == nil {
		return result
	}
	return str
}
