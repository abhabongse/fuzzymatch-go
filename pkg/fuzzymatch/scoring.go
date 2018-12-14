/*
Package fuzzymatch implements an approximate string matching algorithm to determine
the similarity score between two given strings. In most cases, you would be
interested in the function SimilarityScore.
*/
package fuzzymatch

import (
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/dicecoeff"
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/editdistance"
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/normalization"
	"golang.org/x/text/runes"
	"math"
)

/*
SimilarityScore is the definitive function which computes the similarity score
between two input strings. The returned score value is a floating point between
0 (meaning two strings are very distinct) and 1 (meaning two strings are identical
when normalized).
*/
func SimilarityScore(fst, snd string) float64 {
	normalizedFst := normalizeString(fst)
	normalizedSnd := normalizeString(snd)

	// Breaking ties to save memory in the long run
	if len(normalizedFst) < len(normalizedSnd) {
		normalizedFst, normalizedSnd = normalizedSnd, normalizedFst
	}

	optDistRatio := optimalAlignmentDistanceRatio(normalizedFst, normalizedSnd)
	diceCoefficient := dicecoeff.DiceSimilarityCoefficient(normalizedFst, normalizedSnd)

	combinedScore := (optDistRatio + 2.0*diceCoefficient) / 3.0
	return clipNumberToBound(combinedScore, 0.0, 1.0)
}

/*
normalizeString normalizes an input string via various normalization methods.
*/
func normalizeString(str string) string {

	str = normalization.ApplyTransformers(
		str,
		// Sanitize for errors in decoding of Unicode string
		runes.ReplaceIllFormed(),
		// Remove non-printing rune characters
		normalization.StripNonPrintTransformer,
		// Replace all white-spaces to normal space
		normalization.ToNormalSpaceTransformer,
	)

	// Re-spacing the entire string by stripping out leading+trailing spaces,
	// and then transforming multiple consecutive spaces with a single space
	str = normalization.ReSpace(str)

	str = normalization.ApplyTransformers(
		str,
		// Remove diacritical marks above latin characters
		normalization.RemoveAccentsTransformer,
		// Convert western characters into their lowercase forms
		normalization.ToLowerTransformer,
	)

	// Special rule: combine characters for sara-ae and sara-am
	str = normalization.RecombineThaiGrams(str)

	// TODO: introduce multiple string normalization functions

	return str
}

/*
optimalAlignmentDistanceRatio computes the unit-normalized optimal alignment
distance metrics between two input strings. This unit-normalization is conducted
to make sure that the returned score is between 0 and 1. The term "normalization"
used here is not to be confused with the normalization of strings.
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
