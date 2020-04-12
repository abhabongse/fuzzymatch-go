package thai

import (
	"github.com/abhabongse/fuzzymatch-go/transform"
)

// Sanitize extends on the LatinExtendedSanitize
// by additionally sanitize an input string containing Thai scripts.
//
// TODO: revamp this function
func Sanitize(str string) string {

	// Pre-process the string with the most common string sanitization functions
	str = transform.LatinExtendedSanitize(str)
	// Special rule: combine characters for sara-ae and sara-am
	str = RecombineThaiGrams(str)
	// Special rule: remove accidentally repeated non-spacing marks such as
	// tonal marks, ascending vowels, descending vowels, etc.
	str = RemoveThaiRepeatedAccidents(str)

	return str
}
