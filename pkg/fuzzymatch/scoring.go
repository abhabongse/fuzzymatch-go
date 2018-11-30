/*
Package fuzzymatch implements an approximate string matching algorithm to determine
the similarity score between two given strings. In most cases, you would be
interested in the function SimilarityScore.
*/
package fuzzymatch

import (
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/bigrams"
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/editdistance"
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/normalization"
	"math"
	"strings"
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
	diceCoefficient := bigrams.DiceSimilarityCoefficient(normalizedFst, normalizedSnd)

	combinedScore := (optDistRatio + 2.0*diceCoefficient) / 3.0
	return clipNumber(combinedScore, 0.0, 1.0)
}

/*
normalizeString normalizes an input string via various normalization methods.
*/
func normalizeString(str string) string {
	// TODO: introduce multiple string normalization functions
	str = normalization.NormalizeWhiteSpaces(str)
	str = normalization.RemoveNonPrinting(str)
	str = normalization.RemoveAccents(str)
	str = strings.ToLower(str)
	// ...

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
clipNumber readjust the provided values in between the given range as defined
by upper and lower. If the given value is smaller than lower, lower is returned;
if the given value is larger than upper, upper is returned; otherwise, the value
itself is returned.
*/
func clipNumber(value, lower, upper float64) float64 {
	if value > upper {
		return upper
	}
	if value < lower {
		return lower
	}
	return value
}
