package transform

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

// LatinExtendedSanitize sanitizes an input string
// via various string sanitization methods related to Extended Latin scripts.
//
// TODO: revamp this function
func LatinExtendedSanitize(str string) string {
	str = ApplyTransformers(
		str,
		// LatinExtendedSanitize for errors in decoding of Unicode string
		runes.ReplaceIllFormed(),
		// Remove non-printing rune characters
		StripNonPrintTransformer,
		// Replace all white-spaces to normal space
		ToNormalSpaceTransformer,
		// Replace all dashes and hyphens to normal hyphen
		ToNormalHyphenTransformer,
		// Convert western characters into their lowercase forms
		CaseFoldingTransformer,
		//// Remove diacritical marks above latin characters
		//StripAccentTransformer,
	)
	// Re-spacing the entire string by stripping out leading+trailing spaces,
	// and then transforming multiple consecutive spaces with a single space.
	str = ReSpace(str)
	return str
}

// ApplyTransformers applies each string transformer
// from the given sequence of transformers to the given input string.
// If any transformer produces an error,
// it will be silently ignored and intermediate string will not be affected.
//
// TODO: remove this function
func ApplyTransformers(str string, ts ...transform.Transformer) string {
	for _, t := range ts {
		modifiedStr, _, err := transform.String(t, str)
		if err == nil {
			str = modifiedStr
		}
	}
	return str
}
