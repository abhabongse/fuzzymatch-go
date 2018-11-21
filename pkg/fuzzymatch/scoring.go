/*
Package fuzzymatch implements an approximate string matching algorithm to determine the similarity
score between two given strings. In most cases, you would be interested in the function SimilarityScore.
*/
package fuzzymatch

import (
	"math"
)

/*
A definitive function which computes the similarity score between two input strings.
The returned score value is a floating point between 0 (strings are very distinct)
and 1 (strings are identical).
*/
func SimilarityScore(fst, snd string) float64 {
	fstRunes := normalizeString(fst)
	sndRunes := normalizeString(snd)

	// Breaking ties to save memory in the long run
	if len(fstRunes) < len(sndRunes) {
		fstRunes, sndRunes = sndRunes, fstRunes
	}

	normOptDistScore := normalizedOptimalAlignmentDistance(fstRunes, sndRunes)
	return normOptDistScore
}

/*
Normalize an input string using various methods and return as a slice of runes.
*/
func normalizeString(str string) []rune {
	// TODO: introduce multiple string normalization functions
	str = NormalizeWhiteSpaces(str)
	// ...

	runesData := []rune(str)
	return runesData
}

/*
Compute the "normalized" optimal alignment distance metrics between two given
slices of rune characters. This "normalization" is conducted to make sure that
the returned score is between 0 and 1, and is not to be confused with the
normalization of input strings.
*/
func normalizedOptimalAlignmentDistance(fstRunes, sndRunes []rune) float64 {
	// TODO: replace SimpleAlignmentDistance with the customized distance metric version of the OptimalAlignmentDistance
	dist := SimpleAlignmentDistance(fstRunes, sndRunes)
	fstLength := SimpleAlignmentDistance(fstRunes, []rune{})
	sndLength := SimpleAlignmentDistance(sndRunes, []rune{})
	score := 1.0 - (dist / math.Max(fstLength, sndLength))
	if math.IsNaN(score) { // both are empty strings
		score = 1.0
	}
	return score
}
