package thai

import (
	"github.com/abhabongse/fuzzymatch-go/candidate/diacritics"
	"github.com/abhabongse/fuzzymatch-go/transform"
	"golang.org/x/text/cases"
	"golang.org/x/text/runes"
	"golang.org/x/text/secure/precis"
	xTextTransform "golang.org/x/text/transform"
)

// Sanitize extends on the LatinExtendedSanitize
// by additionally sanitize an input string containing Thai scripts.
//
// TODO: revamp this function
func Sanitize(input string) string {

	output, _, _ := xTextTransform.String(precisThaiSanitizeTransformer, input)

	// Special rule: remove accidentally repeated non-spacing marks such as
	// tonal marks, ascending vowels, descending vowels, etc.
	output = RemoveThaiRepeatedAccidents(output)

	return output
}

var precisThaiSanitizeTransformer = precisThaiSanitizeProfile.NewTransformer()
var precisThaiSanitizeProfile = precis.NewFreeform(
	precis.FoldWidth,
	precis.AdditionalMapping(chainedTransformer),
	precis.FoldCase(cases.HandleFinalSigma(true)),
)

func chainedTransformer() xTextTransform.Transformer {
	return xTextTransform.Chain(
		runes.ReplaceIllFormed(),
		// Remove non-printing rune characters
		transform.StripNonPrintTransformer,
		// Replace all white-spaces to normal space
		transform.ToNormalSpaceTransformer,
		// Replace all dashes and hyphens to normal hyphen
		transform.ToNormalHyphenTransformer,
		// Convert western characters into their lowercase forms
		transform.CaseFoldingTransformer,
		// Remove diacritical marks above latin characters
		diacritics.AsciiFoldTransformer,
		diacritics.StripDiacriticalMarksTransformer,
		// Special rule: combine characters for sara-ae and sara-am
		BigramRecombineTransformer,
		// Respacing the entire string by stripping out leading and trailing spaces,
		// and then replacing inter-word spaces with a single normal space
		transform.RespaceTransformer,
	)
}
