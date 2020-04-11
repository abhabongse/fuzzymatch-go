package sanitary

import (
	"strings"
)

// ReSpace removes leading and trailing white-spaces,
// then it reduces all inter-word white-spaces into a single normal space.
func ReSpace(str string) string {
	return strings.Join(strings.Fields(str), " ")
}
