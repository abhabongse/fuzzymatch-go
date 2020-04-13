package thai

import (
	"github.com/abhabongse/fuzzymatch-go/candidate/diacritics"
	"github.com/abhabongse/fuzzymatch-go/factory"
	fuzzyTransform "github.com/abhabongse/fuzzymatch-go/transform"
	"golang.org/x/text/cases"
	"golang.org/x/text/runes"
	"golang.org/x/text/secure/precis"
	"golang.org/x/text/transform"
)

// Sanitize extends on the LatinExtendedSanitize
// by additionally sanitize an input string containing Thai scripts.
var Sanitize = factory.MakeStringTransformFunction(
	PrecisProfile.NewTransformer(),
)

// PrecisProfile is a Unicode PRECIS profile which prepare strings for a more secure comparison.
var PrecisProfile = precis.NewFreeform(
	precis.FoldWidth,
	precis.AdditionalMapping(func() transform.Transformer {
		return transform.Chain(AdditionalMapping...)
	}),
	precis.FoldCase(cases.HandleFinalSigma(true)),
)

// AdditionalMapping contains a slice of all string transformers
// chained together which is used as additional mapping for PRECIS profile.
var AdditionalMapping = []transform.Transformer{
	runes.ReplaceIllFormed(),
	// Remove non-printing rune characters
	fuzzyTransform.StripNonPrintTransformer,
	// Replace all white-spaces to normal space
	fuzzyTransform.ToNormalSpaceTransformer,
	// Replace all dashes and hyphens to normal hyphen
	fuzzyTransform.ToNormalHyphenTransformer,
	// Remove diacritical marks above latin characters
	diacritics.AsciiFoldTransformer,
	diacritics.StripDiacriticalMarksTransformer,
	// Special rule: combine characters for sara-ae and sara-am
	BigramRecombineTransformer,
	// Special rule: remove accidentally repeated non-spacing marks such as
	// tonal marks, ascending vowels, descending vowels, etc.
	RemoveRepeatedMarksTransformer,
	// Respacing the entire string by stripping out leading and trailing spaces,
	// and then replacing inter-word spaces with a single normal space
	fuzzyTransform.RespaceTransformer,
}
