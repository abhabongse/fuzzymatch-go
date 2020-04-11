package thai

import "github.com/abhabongse/fuzzymatch-go/sanitary"

// Sanitize extends on the LatinExtendedSanitize
// by additionally sanitize an input string containing Thai scripts.
func Sanitize(str string) string {

	// Pre-process the string with the most common string sanitization functions
	str = sanitary.LatinExtendedSanitize(str)
	// Special rule: combine characters for sara-ae and sara-am
	str = RecombineThaiGrams(str)
	// Special rule: remove accidentally repeated non-spacing marks such as
	// tonal marks, ascending vowels, descending vowels, etc.
	str = RemoveThaiRepeatedAccidents(str)

	return str
}
