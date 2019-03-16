package sanitary

import (
	"strings"
)

/*
Noop does not do anything to the input string and return the string as-is.
*/
func Noop(str string) string {
	return str
}

/*
ReSpace removes leading and trailing white-spaces, then it reduces all inter-word
white-spaces into a single normal space.
*/
func ReSpace(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

