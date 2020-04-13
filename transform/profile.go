package transform

import (
	"github.com/abhabongse/fuzzymatch-go/candidate/diacritics"
	"github.com/abhabongse/fuzzymatch-go/factory"
	"golang.org/x/text/cases"
	"golang.org/x/text/runes"
	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
)

// LatinExtendedSanitize sanitizes an input string via various
// string sanitization methods related to Extended Latin scripts.
var LatinExtendedSanitize = factory.MakeStringTransformFunction(
	LatinExtendedPrecisProfile.NewTransformer(),
)

// LatinExtendedPrecisProfile is a Unicode PRECIS profile
// which prepare strings for a more secure comparison.
var LatinExtendedPrecisProfile = precis.NewFreeform(
	precis.FoldWidth,
	precis.AdditionalMapping(func() transform.Transformer {
		return transform.Chain(LatinExtendedAdditionalMapping...)
	}),
	precis.FoldCase(cases.HandleFinalSigma(true)),
)

// LatinExtendedAdditionalMapping contains a slice of all string transformers
// chained together which is used as additional mapping for PRECIS profile.
var LatinExtendedAdditionalMapping = []transform.Transformer{
	runes.ReplaceIllFormed(),
	// Remove non-printing rune characters
	StripNonPrintTransformer,
	// Replace all white-spaces to normal space
	ToNormalSpaceTransformer,
	// Replace all dashes and hyphens to normal hyphen
	ToNormalHyphenTransformer,
	// Remove diacritical marks above latin characters
	diacritics.AsciiFoldTransformer,
	diacritics.StripDiacriticalMarksTransformer,
	// Respacing the entire string by stripping out leading and trailing spaces,
	// and then replacing inter-word spaces with a single normal space
	RespaceTransformer,
}
