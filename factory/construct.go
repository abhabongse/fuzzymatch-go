package factory

import (
	"github.com/abhabongse/fuzzymatch-go/candidate"
	"github.com/abhabongse/fuzzymatch-go/dicecoeff"
	"github.com/abhabongse/fuzzymatch-go/editdist"
	"github.com/abhabongse/fuzzymatch-go/sanitary"
	"math"
)

/*
NewSimilarityScoreFunc constructs a new function to compute the string similarity
scores from two input strings based on the provided option setters.

This function follow Go's pattern of functional Options. Since this function is
computationally intensive, one possible optimization is to create a string similarity
score once by putting them at the global level.
*/
func NewSimilarityScoreFunc(setters ...OptionSetter) func(string, string) float64 {

	// Let us start from the default configuration
	config := &Options{
		sanitizeStringFunc:     sanitary.Noop,
		generateCandidatesFunc: candidate.GenerateDefault,
		editDistanceFunc:       editdist.SimpleAlignmentNormDist,
		combineScoresFunc:      EditDistSubScoreOnly,
	}

	// For each addition option setters, apply them to the configuration structure
	for _, setter := range setters {
		setter(config)
	}

	// We linearly combine the Optimal Alignment distance sub-scoring with the Dice Similarity
	// sub-scoring function into a single definite string similarity scoring function.
	combinedSimilarityScore := func(fst, snd string) float64 {
		editDistanceSubScore := 1.0 - config.editDistanceFunc(fst, snd)
		diceCoefficientSubScore := dicecoeff.DiceSimilarityCoefficient(fst, snd)

		combinedScore := config.combineScoresFunc(editDistanceSubScore, diceCoefficientSubScore)
		return clipNumberToBound(combinedScore, 0.0, 1.0)
	}

	// Finally, we use the combined distance scoring function (constructed above) to compute
	// the distances between all possible pairs of candidates from both input string.
	// The final score would be those yielding the maximum combined distance score.
	bestCandidatesSimilarityScore := func(fst, snd string) float64 {
		fst = config.sanitizeStringFunc(fst)
		snd = config.sanitizeStringFunc(snd)

		// Breaking ties to save memory space (see implementation of Levenshtein algorithm)
		if len(fst) < len(snd) {
			fst, snd = snd, fst
		}

		// Find the highest score based on all pairs of generated candidates
		bestScore := 0.0
		for _, candidateOne := range config.generateCandidatesFunc(fst) {
			for _, candidateTwo := range config.generateCandidatesFunc(snd) {
				score := combinedSimilarityScore(candidateOne, candidateTwo)
				bestScore = math.Max(bestScore, score)
			}
		}

		return bestScore
	}

	return bestCandidatesSimilarityScore
}

/*
EditDistSubScoreOnly utilizes only the edit distance sub-score and completely
ignore the Dice coefficient sub-score.
*/
func EditDistSubScoreOnly(editDistSubScore, diceCoeffSubScore float64) float64 {
	return 1.0*editDistSubScore + 0.0*diceCoeffSubScore
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
