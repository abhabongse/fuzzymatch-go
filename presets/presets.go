/*
Package presets is a collection of pre-combined approximate string matching
algorithms which determines a similarity score between two strings.

In most cases, you would be interested in the function 'SimilarityScore'.
*/
package presets

import (
	"github.com/abhabongse/fuzzymatch-go/canonical"
	"github.com/abhabongse/fuzzymatch-go/dicecoefficient"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
	"golang.org/x/text/runes"
)

/*
SimilarityScore is the definitive function which computes the similarity score
between two input strings. The returned score value is a floating point between
0 (meaning two strings are very distinct) and 1 (meaning two strings are identical
when canonicalized).
*/
func SimilarityScore(fst, snd string) float64 {
	canonicalFst := canonicalize(fst)
	canonicalSnd := canonicalize(snd)

	// Breaking ties to save memory in the long run
	if len(canonicalFst) < len(canonicalSnd) {
		canonicalFst, canonicalSnd = canonicalSnd, canonicalFst
	}

	optDistRatio := 1.0 - optimalAlignmentDistanceRatio(canonicalFst, canonicalSnd)
	diceCoefficient := dicecoefficient.DiceSimilarityCoefficient(canonicalFst, canonicalSnd)

	combinedScore := (optDistRatio + 2.0*diceCoefficient) / 3.0
	return clipNumberToBound(combinedScore, 0.0, 1.0)
}

/*
canonicalize normalizes an input string via various canonicalization methods.
*/
func canonicalize(str string) string {

	str = canonical.ApplyTransformers(
		str,
		// Sanitize for errors in decoding of Unicode string
		runes.ReplaceIllFormed(),
		// Remove non-printing rune characters
		canonical.StripNonPrintTransformer,
		// Replace all white-spaces to normal space
		canonical.ToNormalSpaceTransformer,
		// Remove diacritical marks above latin characters
		canonical.RemoveAccentsTransformer,
		// Convert western characters into their lowercase forms
		canonical.ToLowerTransformer,
	)

	// Re-spacing the entire string by stripping out leading+trailing spaces,
	// and then transforming multiple consecutive spaces with a single space
	str = canonical.ReSpace(str)

	// Special rule: combine characters for sara-ae and sara-am
	str = canonical.RecombineThaiGrams(str)
	// Special rule: remove accidentally repeated non-spacing marks such as
	// tonal marks, ascending vowels, descending vowels, etc.
	str = canonical.RemoveThaiRepeatedAccidents(str)

	return str
}

// Normalized version of the Optimal Alignment distance scoring function with unit-distance penalty
var optimalAlignmentDistanceRatio = editdistance.MakeNormalized(editdistance.SimpleAlignmentDistance)
