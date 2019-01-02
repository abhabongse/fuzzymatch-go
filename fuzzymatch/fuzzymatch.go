/*
Package fuzzymatch contains the customizable factory which generates string similarity
scoring functions based on various configurations such as how strings are canonicalized
or which rune distance metrics are used to measure distance between strings.
*/
package fuzzymatch

import (
	"github.com/abhabongse/fuzzymatch-go/canonical"
	"github.com/abhabongse/fuzzymatch-go/dicecoefficient"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
	"golang.org/x/text/runes"
	"math"
)

/*
NewSimilarityScoreFunction constructs a new function to compute the string similarity
scores from two input strings based on the provided option setters.

This function follow Go's pattern of functional options. Since this function is
computationally intensive, one possible optimization is to create a string similarity
score once by putting them at the global level.

TODO: add usage examples
*/
func NewSimilarityScoreFunction(setters ...OptionSetter) (func(string, string) float64, error) {

	// Let us start from the default configuration
	config := &options{
		stringCanonicalizationFunc: DefaultCanonicalizeString,
		candidateGenerationFunc:    DefaultGenerateCandidates,
		substitutionPenaltyFunc:    editdistance.UnitDist,
		transpositionPenaltyFunc:   editdistance.UnitDist,
		optimalAlignmentWeight:     1.0,
		diceSimilarityWeight:       0.0,
	}

	// For each addition option setters, apply them to the configuration structure
	for _, setter := range setters {
		if err := setter(config); err != nil {
			return nil, err // short-circuit if there is an error
		}
	}

	// Based on the configured substitution and transposition penalty functions, we build
	// a customized version of the normalized Optimal Alignment score.
	customOptimalAlignmentDistance := editdistance.MakeNormalized(editdistance.MakeAlignmentDistanceFunction(
		config.substitutionPenaltyFunc, config.transpositionPenaltyFunc,
	))

	// We linearly combine the Optimal Alignment distance sub-scoring with the Dice Similarity
	// sub-scoring function into a single definite string similarity scoring function.
	combinedDistanceScoreFunction := func(fst, snd string) float64 {
		optimalAlignmentSubScore := 1.0 - customOptimalAlignmentDistance(fst, snd)
		diceCoefficientSubScore := dicecoefficient.DiceSimilarityCoefficient(fst, snd)

		numerator := config.optimalAlignmentWeight*optimalAlignmentSubScore + config.diceSimilarityWeight*diceCoefficientSubScore
		denominator := config.optimalAlignmentWeight + config.diceSimilarityWeight
		combinedScore := numerator / denominator

		return clipNumberToBound(combinedScore, 0.0, 1.0)
	}

	// Finally, we use the combined distance scoring function (constructed above) to compute
	// the distances between all possible pairs of candidates from both input string.
	// The final score would be those yielding the maximum combined distance score.
	bestPairDistanceScoreFunction := func(fst, snd string) float64 {
		fst = config.stringCanonicalizationFunc(fst)
		snd = config.stringCanonicalizationFunc(snd)

		// Breaking ties to save memory space (see implementation of Levenshtein algorithm)
		if len(fst) < len(snd) {
			fst, snd = snd, fst
		}

		// Find the highest score based on all pairs of generated candidates
		bestScore := 0.0
		for _, candidateOne := range config.candidateGenerationFunc(fst) {
			for _, candidateTwo := range config.candidateGenerationFunc(snd) {
				score := combinedDistanceScoreFunction(candidateOne, candidateTwo)
				bestScore = math.Max(bestScore, score)
			}
		}

		return bestScore
	}

	return bestPairDistanceScoreFunction, nil
}

/*
DefaultCanonicalizeString normalizes an input string via various string canonicalization
methods specialized for string containing Thai scripts.
*/
func DefaultCanonicalizeString(str string) string {

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
DefaultGenerateCandidates uses the input string itself as the only candidate.
*/
func DefaultGenerateCandidates(str string) []string {
	candidates := []string{str}
	return candidates
}

/*
clipNumberToBound readjust the provided values in between the given range as
defined by arguments upper and lower. If the given value is smaller than lower,
lower is returned; if the given value is larger than upper, upper is returned;
otherwise, the value itself is returned.

Warning: if the value for lower provided is greater than the value for upper,
then this function has undefined behavior and its output has no guarantee.
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
