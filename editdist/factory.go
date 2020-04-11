// This source file contains higher-order factory functions
// which creates another string distance computation functions.

package editdist

import "math"

// MakeLevenshteinDistFunction is a higher-order function
// which receives the substitution penalty function
// and returns the modified version of the otherwise generic OptimalAlignmentDist function
// with those specified edit distance metrics.
// The resulting function will receive two input strings and output the distance between them.
func MakeLevenshteinDistFunction(substPenaltyFunction RunePenaltyFunction) StringDistFunction {
	return func(fst, snd string) float64 {
		return LevenshteinDist(fst, snd, substPenaltyFunction)
	}
}

// MakeOptimalAlignmentDistFunction is a higher-order function
// which receives the substitution and the transposition penalty functions
// and returns the modified version of the otherwise generic OptimalAlignmentDist function
// with those specified edit distance metrics.
// The resulting function will receive two input strings and output the distance between them.
func MakeOptimalAlignmentDistFunction(substPenaltyFunction, transPenaltyFunction RunePenaltyFunction) StringDistFunction {
	return func(fst, snd string) float64 {
		return OptimalAlignmentDist(fst, snd, substPenaltyFunction, transPenaltyFunction)
	}
}

// MakeNormalized is a higher-order function
// which converts a typical edit-distance computation function in the string space
// into its normalized version.
// Therefore, the returned score of the new edit-distance function
// should always be between 0 (meaning two strings are very similar)
// and 1  (meaning two strings are very different).
//
// Implementation details: the denominator of the fraction
// is _not_ the length of the longer string.
// The reason for this is that some insertion/deletion errors incur sub-unit penalties.
// Without the size-fitting denominator,
// a malicious user may attack by saturating those insertions/deletions
// in order to decrease the total edit distances.
func MakeNormalized(inputFunction StringDistFunction) StringDistFunction {
	return func(fst, snd string) float64 {
		dist := inputFunction(fst, snd)
		fstCapacity := inputFunction(fst, "")
		sndCapacity := inputFunction("", snd)
		score := dist / math.Max(fstCapacity, sndCapacity)
		if math.IsNaN(score) {
			score = 0.0 // both strings are empty
		}
		return score
	}
}
