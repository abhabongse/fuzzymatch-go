package fuzzymatch

import "strings"

/*
Remove leading and trailing white-spaces, then reduce all inter-word white-spaces
into a single normal space.
*/
func NormalizeWhiteSpaces(str string) string {
	return strings.Join(strings.Fields(str), " ")
}
