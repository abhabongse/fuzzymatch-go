// This source file contains higher-order factory functions
// which creates another string distance computation functions.

package editdist

import (
	"math"
)

// MakeLevenshteinDistFunction is a higher-order function
// which uses the provided substitution penalty function to construct
// the modified version of the otherwise generic OptimalAlignmentDist function
// with those specified edit distance metrics.
// The resulting function would receive two input strings and output the distance between them.
func MakeLevenshteinDistFunction(substPenaltyFunction RunePenaltyFunction) StringDistFunction {
	return func(fst, snd string) float64 {
		return LevenshteinDist(fst, snd, substPenaltyFunction)
	}
}

// MakeOptimalAlignmentDistFunction is a higher-order function
// which uses the provided substitution and the transposition penalty functions
// to construct the modified version of the otherwise generic OptimalAlignmentDist function
// with those specified edit distance metrics.
// The resulting function will receive two input strings and output the distance between them.
func MakeOptimalAlignmentDistFunction(substPenaltyFunction, transPenaltyFunction RunePenaltyFunction) StringDistFunction {
	return func(fst, snd string) float64 {
		return OptimalAlignmentDist(fst, snd, substPenaltyFunction, transPenaltyFunction)
	}
}

// MakeStringSimilarityFunction is a higher-order function
// which converts the provided typical edit-distance computation function
// (i.e. function which receives two strings and returns edit distance between them)
// into the normalized string similarity function.
// The resulting function would return a score within the range from 0 to 1
// where 1 indicates that both strings are identical
// and 0 indicates that both strings are totally distinct.
//
// Implementation details: the denominator of the fraction is _not_ the length of the longer string.
// The reason for this is that some insertion/deletion errors incur sub-unit penalties.
// Without the size-fitting denominator, a malicious user may attack
// by saturating those insertions/deletions in order to decrease the total edit distances.
func MakeStringSimilarityFunction(inputFunction StringDistFunction) StringDistFunction {
	return func(fst, snd string) float64 {
		dist := inputFunction(fst, snd)
		fstCapacity := inputFunction(fst, "")
		sndCapacity := inputFunction("", snd)
		normDist := dist / math.Max(fstCapacity, sndCapacity)
		if math.IsNaN(normDist) {
			normDist = 0.0 // both strings are empty
		}
		return clipNumberToBound(1-normDist, 0, 1)
	}
}

// clipNumberToBound re-adjust the provided values
// in between the given range as defined by arguments upper and lower.
// If the given value is smaller than lower, lower is returned;
// if the given value is larger than upper, upper is returned;
// otherwise, the value itself is returned.
//
// Warning: if the value for lower provided is greater than the value for upper,
// then this function has undefined behavior and its output has no guarantee.
func clipNumberToBound(value, lower, upper float64) float64 {
	if value < lower {
		return lower
	}
	if value > upper {
		return upper
	}
	return value
}
