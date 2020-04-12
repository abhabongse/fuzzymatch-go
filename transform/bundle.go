package transform

import (
	"github.com/abhabongse/fuzzymatch-go/candidate/diacritics"
	"golang.org/x/text/cases"
	"golang.org/x/text/runes"
	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
)

// LatinExtendedSanitize sanitizes an input string
// via various string sanitization methods related to Extended Latin scripts.
func LatinExtendedSanitize(input string) string {
	output, _, _ := transform.String(precisLatinExtendedSanitizeTransformer, input)
	return output
}

var precisLatinExtendedSanitizeTransformer = precisLatinExtendedSanitizeProfile.NewTransformer()
var precisLatinExtendedSanitizeProfile = precis.NewFreeform(
	precis.FoldWidth,
	precis.AdditionalMapping(chainedTransformer),
	precis.FoldCase(cases.HandleFinalSigma(true)),
)

func chainedTransformer() transform.Transformer {
	return transform.Chain(
		runes.ReplaceIllFormed(),
		// Remove non-printing rune characters
		StripNonPrintTransformer,
		// Replace all white-spaces to normal space
		ToNormalSpaceTransformer,
		// Replace all dashes and hyphens to normal hyphen
		ToNormalHyphenTransformer,
		// Convert western characters into their lowercase forms
		CaseFoldingTransformer,
		// Remove diacritical marks above latin characters
		diacritics.AsciiFoldTransformer,
		diacritics.StripDiacriticalMarksTransformer,
		// Respacing the entire string by stripping out leading and trailing spaces,
		// and then replacing inter-word spaces with a single normal space
		RespaceTransformer,
	)
}
