/*
Package normalization contains a collection of string normalization functions; they may
be used to pre-process input strings.
*/
package normalization

import (
	"golang.org/x/text/transform"
	"strings"
)

/*
ReSpace removes leading and trailing white-spaces, then it reduces all inter-word
white-spaces into a single normal space.
*/
func ReSpace(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

/*
StripNonPrint removes all occurrences of non-printing and non-spacing rune characters
from a string.
*/
func StripNonPrint(str string) string {
	return ApplyTransformer(StripNonPrintTransformer, str)
}

/*
ToNormalSpace replaces all white space rune characters into a normal space.
*/
func ToNormalSpace(str string) string {
	return ApplyTransformer(ToNormalSpaceTransformer, str)
}

/*
RemoveAccents tries to remove as many combining diacritical marks from the input
string as possible.
*/
func RemoveAccents(str string) string {
	return ApplyTransformer(RemoveAccentsTransformer, str)
}

/*
_ToLower transforms all unicode characters into its lowercase forms as defined
by Unicode property. Avoid this function and use string.Lower instead; this function
exists solely for testing of Unicode transformers.
*/
func _ToLower(str string) string {
	return ApplyTransformer(ToLowerTransformer, str)
}

/*
ApplyTransformer is a helper function which applies the unicode transformer to
an input string; whenever errors occur, the original input string will be returned.
*/
func ApplyTransformer(t transform.Transformer, str string) string {
	result, _, err := transform.String(t, str)
	if err == nil {
		return result
	}
	return str
}
