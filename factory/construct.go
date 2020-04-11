package factory

import (
	"github.com/abhabongse/fuzzymatch-go/editdist"
	"math"
)

// MakeSimilarityScoreFunction constructs a new function
// which computes the string similarity score from two input strings.
// By default,
// - Input strings will not be sanitized
// - Extra candidates will not be generated
// - String similarity will be based on optimal alignment with unit penalty
func MakeSimilarityScoreFunction(setters ...OptionSetter) func(string, string) float64 {

	// Let us start from the default configuration
	config := &Options{
		sanitizeString:     sanitizeNoop,
		generateCandidates: generateInputOnlyCandidate,
		computeSimilarity:  simpleOptimalAlignmentStringSimilarity,
	}

	// For each addition option setters, apply them to the configuration structure
	for _, setter := range setters {
		setter(config)
	}

	// Finally, we use the combined distance scoring function (constructed above) to compute
	// the distances between all possible pairs of candidates from both input string.
	// The final score would be those yielding the maximum combined distance score.
	bestCandidatesSimilarityScore := func(fst, snd string) float64 {

		// Sanitize input strings
		fst = config.sanitizeString(fst)
		snd = config.sanitizeString(snd)

		// Breaking ties to save memory space (see implementation of Levenshtein algorithm)
		if len(fst) < len(snd) {
			fst, snd = snd, fst
		}

		// Generate candidates from two string inputs
		candidateGroupA := config.generateCandidates(fst)
		candidateGroupB := config.generateCandidates(snd)

		// Find the highest score based on all pairs of generated candidates
		bestScore := 0.0
		for _, candidateA := range candidateGroupA {
			for _, candidateB := range candidateGroupB {
				score := config.computeSimilarity(candidateA, candidateB)
				bestScore = math.Max(bestScore, score)
			}
		}
		return bestScore
	}
	return bestCandidatesSimilarityScore
}

// sanitizeNoop does not do anything to the input string and return the string as-is.
func sanitizeNoop(str string) string {
	return str
}

// generateInputOnlyCandidate uses the input string itself as the only candidate.
func generateInputOnlyCandidate(str string) []string {
	candidates := []string{str}
	return candidates
}

// simpleOptimalAlignmentDist computes the optimal alignment distance between two strings
// using unit penalty for both character substitutions and transpositions
var simpleOptimalAlignmentDist = editdist.MakeOptimalAlignmentDistFunction(editdist.UnitPenalty, editdist.UnitPenalty)

// simpleOptionalAlignmentStringSimilarity computes the similarity score of two strings
// based on simpleOptimalAlignmentDist function
var simpleOptimalAlignmentStringSimilarity = editdist.MakeStringSimilarityFunction(simpleOptimalAlignmentDist)
