package fuzzymatch

import (
	"fmt"
	"math"
)

/*
options type stores information of how two strings are being compared,
particularly, how strings are canonicalized, and what metrics are used
individually, follows Go's functional options patterns.
*/
type options struct {
	// String canonicalization function which is applied to each input string
	stringCanonicalization func(string) string
	// Candidate generation function which is applied to each input string
	candidateGeneration func(string) []string
	// Substitution/Transposition distance penalty functions which is used during
	// the computation of the Optimal Alignment by the Levenshtein's algorithm
	substitutionPenalty func(rune, rune) float64
	transpositionPenalty func(rune, rune) float64
	// Linear combination weights to be assigned to each similarity sub-score:
	// the Optimal Alignment sub-score and the Dice Similarity sub-score
	optimalAlignmentWeight float64
	diceSimilarityWeight float64
}

/*
OptionSetter type maintains the configuration information regarding how
two strings are compared in order to compute the overall similarity score.
*/
type OptionSetter func(*options) error

/*
StringCanonicalization specifies the string canonicalization function used
to clean up each input string.
 */
func StringCanonicalization(stringCanonicalizationFunc func(string) string) OptionSetter {
	return func(config *options) error {
		config.stringCanonicalization = stringCanonicalizationFunc
		return nil
	}
}

/*
CandidateGeneration specifies the function which generates candidates from each
input string.
*/
func CandidateGeneration(candidateGenerationFunc func(string) []string) OptionSetter {
	return func(config *options) error {
		config.candidateGeneration = candidateGenerationFunc
		return nil
	}
}

/*
RuneDistancePenalties specifies the substitution and the transposition distance penalty
functions used in the computation of Optimal Alignment distance of two strings.
 */
func RuneDistancePenalties(substitutionPenalty, tranpositionPenalty func(rune, rune) float64) OptionSetter {
	return func(config *options) error {
		config.substitutionPenalty = substitutionPenalty
		config.transpositionPenalty = tranpositionPenalty
		return nil
	}
}

/*
CombinationWeights specifies the linear combination weights for the Optimal
Alignment distance sub-score and the Dice Similarity coefficients sub-score.
*/
func CombinationWeights(optimalAlignmentWeight, diceSimilarityWeight float64) OptionSetter {
	return func(config *options) error {
		if math.IsNaN(optimalAlignmentWeight) || optimalAlignmentWeight <= 0.0 {
			return fmt.Errorf("optimalAlignmentWeight should be positive: given %v", optimalAlignmentWeight)
		}
		if math.IsNaN(diceSimilarityWeight) || diceSimilarityWeight <= 0.0 {
			return fmt.Errorf("diceSimilarityWeight should be positive: given %v", diceSimilarityWeight)
		}
		config.optimalAlignmentWeight = optimalAlignmentWeight
		config.diceSimilarityWeight = diceSimilarityWeight
		return nil
	}
}
