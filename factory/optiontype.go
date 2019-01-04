package factory

import (
	"fmt"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
	"math"
)

/*
Options is a type struct that stores the configuration information regarding
how to compute string similarity score between two input strings, particularly
(1) how strings are canonicalized,
(2) how variants/candidates are generates,
(3) how edit distances between both strings are computed, and
(4) how edit distance sub-score and the Dice coefficient sub-score will be
combined to compute the final score, etc.
*/
type Options struct {
	// String canonicalization function to be applied to each input string
	stringCanonicalizationFunc func(string) string
	// Candidate generation function to be applied to each input string
	candidateGenerationFunc func(string) []string
	// Edit distance function to be applied to both strings
	editDistanceFunc func(string, string) float64
	// Combined score function to be applied to edit distance sub-score and
	// the Dice coefficient sub-score
	combinedScoreFunc func(float64, float64) float64
}

/*
OptionSetter is a type alias for functions that modify given Options type struct.
Functions of this type can be used to configure how to compute the overall
string similarity scores between any two input strings.
*/
type OptionSetter func(*Options)

/*
StringCanonicalization assigns the function that would be used to clean up each
of the input strings before they are subsequently compared. In specific, this
wrappedFunc will receive a string as the only input argument and it should return
the canonicalized string of the input.
*/
func StringCanonicalization(wrappedFunc func(string) string) OptionSetter {
	return func(config *Options) {
		config.stringCanonicalizationFunc = wrappedFunc
	}
}

/*
CandidateGeneration assigns the function that would be used to generate all
normalization variants of the already-canonicalized input string. In specific,
this function will receive a string as the only input argument and it should return
a slice of strings each indicating a possible variant of the input string.
*/
func CandidateGeneration(wrappedFunc func(string) []string) OptionSetter {
	return func(config *Options) {
		config.candidateGenerationFunc = wrappedFunc
	}
}

/*
OptimalAlignmentEditDistance constructs a customized version of the normalized
Optimal Alignment scoring function from the provided substitution/transposition
penalty rune distance metrics. This resulting function is then assigned to compute
the edit distance between two input strings.
*/
func OptimalAlignmentEditDistance(substitutionPenalty, transpositionPenalty editdistance.RuneDistanceMetric) OptionSetter {
	return func(config *Options) {
		alignmentDistanceFunc := editdistance.MakeAlignmentDistanceFunction(substitutionPenalty, transpositionPenalty)
		config.editDistanceFunc = editdistance.MakeNormalized(alignmentDistanceFunc)
	}
}

/*
LinearCombinedScore constructs a linear combination function which combines the
edit distance sub-score with the Dice coefficient sub-score with the pre-specified
weights. The resulting function is then assigned to compute the combined score.
*/
func LinearCombinedScore(editDistanceWeight, diceCoefficientWeight float64) OptionSetter {
	return func(config *Options) {
		if math.IsNaN(editDistanceWeight) || editDistanceWeight < 0.0 {
			panic(fmt.Sprintf("editDistanceWeight should be non-negative: given %v", editDistanceWeight))
		}
		if math.IsNaN(diceCoefficientWeight) || diceCoefficientWeight < 0.0 {
			panic(fmt.Sprintf("diceCoefficientWeight should be non-negative: given %v", diceCoefficientWeight))
		}
		if editDistanceWeight+diceCoefficientWeight <= 0.0 {
			panic(fmt.Sprintf("editDistanceWeight + diceCoefficientWeight should be positive: given 0s"))
		}

		config.combinedScoreFunc = func(editDistanceSubScore, diceCoefficientSubScore float64) float64 {
			numerator := editDistanceWeight*editDistanceSubScore + diceCoefficientWeight*diceCoefficientSubScore
			denominator := editDistanceWeight + diceCoefficientWeight
			return numerator / denominator
		}
	}
}
