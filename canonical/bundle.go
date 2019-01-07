package canonical

import "golang.org/x/text/runes"

/*
CanonicalizeString normalizes an input string via various string canonicalization methods.
*/
func CanonicalizeString(str string) string {

	str = ApplyTransformers(
		str,
		// Sanitize for errors in decoding of Unicode string
		runes.ReplaceIllFormed(),
		// Remove non-printing rune characters
		StripNonPrintTransformer,
		// Replace all white-spaces to normal space
		ToNormalSpaceTransformer,
		// Remove diacritical marks above latin characters
		RemoveAccentsTransformer,
		// Convert western characters into their lowercase forms
		ToLowerTransformer,
	)

	// Re-spacing the entire string by stripping out leading+trailing spaces,
	// and then transforming multiple consecutive spaces with a single space
	str = ReSpace(str)

	return str
}

/*
CanonicalizeThaiString normalizes an input string via various string canonicalization
methods, some of which are specialized for string containing Thai scripts.
*/
func CanonicalizeThaiString(str string) string {

	// Pre-process the string with the most common string canonicalization functions
	str = CanonicalizeString(str)
	// Special rule: combine characters for sara-ae and sara-am
	str = RecombineThaiGrams(str)
	// Special rule: remove accidentally repeated non-spacing marks such as
	// tonal marks, ascending vowels, descending vowels, etc.
	str = RemoveThaiRepeatedAccidents(str)

	return str
}
