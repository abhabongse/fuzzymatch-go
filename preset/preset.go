/*
Package preset is a collection of pre-combined approximate string matching
algorithms which determines a similarity score between two strings.

In most cases, you would be interested in the function 'SimilarityScore'.
*/
package preset

import (
	"github.com/abhabongse/fuzzymatch-go/canonical"
	"github.com/abhabongse/fuzzymatch-go/dicecoefficient"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
	"golang.org/x/text/runes"
	"math"
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

	optDistRatio := optimalAlignmentDistanceRatio(canonicalFst, canonicalSnd)
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

/*
optimalAlignmentDistanceRatio computes the unit-normalized optimal alignment
distance metrics between two input strings. This unit-normalization is conducted
to make sure that the returned score is between 0 and 1.
*/
func optimalAlignmentDistanceRatio(fst, snd string) float64 {
	// TODO: replace SimpleAlignmentDistance with the customized distance
	//       metric version of the OptimalAlignmentDistance
	dist := editdistance.SimpleAlignmentDistance(fst, snd)
	fstLength := editdistance.SimpleAlignmentDistance(fst, "")
	sndLength := editdistance.SimpleAlignmentDistance("", snd)
	score := 1.0 - (dist / math.Max(fstLength, sndLength))
	if math.IsNaN(score) { // both are empty strings
		score = 1.0
	}
	return score
}

/*
clipNumberToBound readjust the provided values in between the given range as
defined by arguments upper and lower. If the given value is smaller than lower,
lower is returned; if the given value is larger than upper, upper is returned;
otherwise, the value itself is returned.
*/
func clipNumberToBound(value, lower, upper float64) float64 {
	if value < lower {
		return lower
	}
	if value > upper {
		return upper
	}
	return value
}
