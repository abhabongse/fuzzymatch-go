package fuzzymatch

import (
	"github.com/abhabongse/fuzzymatch-go/canonical"
	"github.com/abhabongse/fuzzymatch-go/dicecoefficient"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
	"golang.org/x/text/runes"
	"math"
)

/*
NewSimilarityScoreComputer creates a new function which computes similarity
scores between two given inputs based on given option setters.
*/
func NewSimilarityScoreFunction(setters ...OptionSetter) (func(string, string) float64, error) {

	// Default configuration
	config := &options{
		stringCanonicalization: DefaultCanonicalizeString,
		candidateGeneration:    DefaultGenerateCandidates,
		substitutionPenalty:    editdistance.UnitDist,
		transpositionPenalty:   editdistance.UnitDist,
		optimalAlignmentWeight: 1.0,
		diceSimilarityWeight:   2.0,
	}

	// Apply option setters
	for _, setter := range setters {
		if err := setter(config); err != nil {
			return nil, err
		}
	}

	// Build a customized version of the normalized Optimal Alignment score based on
	// the configured substitution and transposition penalty functions
	customizedOptimalAlignmentDistance := editdistance.MakeNormalized(
		func(fst, snd string) float64 {
			return editdistance.OptimalAlignmentDistance(fst, snd, config.substitutionPenalty, config.transpositionPenalty)
		})

	// Linearly combine all string similarity sub-scoring functions into one definite function
	combinedDistanceScoreFunction := func(fst, snd string) float64 {
		optimalAlignmentSubScore := 1.0 - customizedOptimalAlignmentDistance(fst, snd)
		diceCoefficientSubScore := dicecoefficient.DiceSimilarityCoefficient(fst, snd)

		combinedScore := (config.optimalAlignmentWeight*optimalAlignmentSubScore + config.diceSimilarityWeight*diceCoefficientSubScore) / (config.optimalAlignmentWeight + config.diceSimilarityWeight)
		return clipNumberToBound(combinedScore, 0.0, 1.0)
	}

	// Introduce the final string similarity scoring function which first canonicalizes
	// input strings, generates possible candidates for each input string, then tries all
	// possible pairs of candidates to see which one yields the maximum combined distance
	// score in the end.
	bestPairDistanceScoreFunction := func(fst, snd string) float64 {
		// Clean up strings
		fst = config.stringCanonicalization(fst)
		snd = config.stringCanonicalization(snd)

		// Breaking ties to save memory space
		if len(fst) < len(snd) {
			fst, snd = snd, fst
		}

		// Try all possible pairs of candidates
		bestScore := 0.0
		for _, candidateOne := range config.candidateGeneration(fst) {
			for _, candidateTwo := range config.candidateGeneration(snd) {
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
