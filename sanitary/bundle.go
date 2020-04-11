package sanitary

import (
	"golang.org/x/text/runes"
)

// LatinExtendedSanitize sanitizes an input string
// via various string sanitization methods related to Extended Latin scripts.
func LatinExtendedSanitize(str string) string {
	str = ApplyTransformers(
		str,
		// LatinExtendedSanitize for errors in decoding of Unicode string
		runes.ReplaceIllFormed(),
		// Remove non-printing rune characters
		StripNonPrintTransformer,
		// Replace all white-spaces to normal space
		ToNormalSpaceTransformer,
		// Remove diacritical marks above latin characters
		StripAccentTransformer,
		// Convert western characters into their lowercase forms
		ToLowerTransformer,
	)
	// Re-spacing the entire string by stripping out leading+trailing spaces,
	// and then transforming multiple consecutive spaces with a single space.
	str = ReSpace(str)
	return str
}
