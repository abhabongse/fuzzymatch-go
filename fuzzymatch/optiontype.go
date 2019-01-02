package fuzzymatch

import (
	"fmt"
	"math"
)

/*
options struct stores the configuration information regarding how two input string
would be used to compute their similarity score: in particular, how strings are
canonicalized, how candidates are generated, what rune distance metrics are used,
and what linear combination weights are used to combine the Optimal Alignment
sub-score and the Dice Similarity sub-score.
*/
type options struct {
	// String canonicalization function which is applied to each input string
	stringCanonicalizationFunc func(string) string
	// Candidate generation function which is applied to each input string
	candidateGenerationFunc func(string) []string
	// Substitution/Transposition distance penalty functions which is used during
	// the computation of the Optimal Alignment by the Levenshtein's algorithm
	substitutionPenaltyFunc  func(rune, rune) float64
	transpositionPenaltyFunc func(rune, rune) float64
	// Linear combination weights to be assigned to each similarity sub-score:
	// the Optimal Alignment sub-score and the Dice Similarity sub-score
	optimalAlignmentWeight float64
	diceSimilarityWeight   float64
}

/*
OptionSetter type maintains the configuration information regarding how
two strings are compared in order to compute the overall similarity score.
*/
type OptionSetter func(*options)

/*
StringCanonicalization specifies the function to canonicalize (i.e. clean up)
each input string before they are compared. The function must receive a string
and output the string already cleaned up.
*/
func StringCanonicalization(stringCanonicalizationFunc func(string) string) OptionSetter {
	return func(config *options) {
		config.stringCanonicalizationFunc = stringCanonicalizationFunc
	}
}

/*
CandidateGeneration specifies the function to generate all candidates based on
the already-canonicalized input string. The function must receive a string and
output a slice of strings each indicating an individual candidate.
*/
func CandidateGeneration(candidateGenerationFunc func(string) []string) OptionSetter {
	return func(config *options) {
		config.candidateGenerationFunc = candidateGenerationFunc
	}
}

/*
RuneDistancePenalties specifies the substitution/transposition rune distance metric
penalty functions which would be used in the computation of Optimal Alignment distance
between two strings.
*/
func RuneDistancePenalties(substitutionPenalty, transpositionPenalty func(rune, rune) float64) OptionSetter {
	return func(config *options) {
		config.substitutionPenaltyFunc = substitutionPenalty
		config.transpositionPenaltyFunc = transpositionPenalty
	}
}

/*
CombinationWeights specifies the linear combination weights for the Optimal Alignment
distance sub-score and the Dice Similarity coefficients sub-score, respectively.

TODO: panic instead of returning errors
TODO: tests for errors
*/
func CombinationWeights(optimalAlignmentWeight, diceSimilarityWeight float64) OptionSetter {
	return func(config *options) {
		if math.IsNaN(optimalAlignmentWeight) || optimalAlignmentWeight < 0.0 {
			panic(fmt.Sprintf("optimalAlignmentWeight should be non-negative: given %v", optimalAlignmentWeight))
		}
		if math.IsNaN(diceSimilarityWeight) || diceSimilarityWeight < 0.0 {
			panic(fmt.Sprintf("diceSimilarityWeight should be non-negative: given %v", diceSimilarityWeight))
		}
		if optimalAlignmentWeight+diceSimilarityWeight <= 0.0 {
			panic(fmt.Sprintf("optimalAlignmentWeight + diceSimilarityWeight should be positive: given 0s"))
		}
		config.optimalAlignmentWeight = optimalAlignmentWeight
		config.diceSimilarityWeight = diceSimilarityWeight
	}
}
