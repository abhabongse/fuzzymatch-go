package candidate

import (
	"github.com/abhabongse/fuzzymatch-go/factory"
	"math"
)

// Chain is a higher-order function
// which multiplies together candidate generator functions
// so that the resulting candidate generator function will
// produce all combinations of candidates from the provided generators.
func Chain(
	generators ...factory.CandidatesGeneratorFunction,
) factory.CandidatesGeneratorFunction {
	return func(input string) []string {
		return recursiveExpandCandidates([]string{input}, generators)
	}
}

func recursiveExpandCandidates(
	previousCandidates []string,
	generators []factory.CandidatesGeneratorFunction,
) []string {
	if len(generators) == 0 {
		return previousCandidates
	}
	firstGenerator := generators[0]
	restGenerators := generators[1:]
	candidates := make([]string, 0)
	for _, previousCandidate := range previousCandidates {
		candidates = append(candidates, firstGenerator(previousCandidate)...)
	}
	return recursiveExpandCandidates(candidates, restGenerators)
}

// PrependStringSanitizer is a higher-order function
// which modifies the provided generateCandidates function
// so that the input string to the function will be sanitized first
// via a call to sanitize function.
func PrependStringSanitizer(
	sanitize factory.StringTransformerFunction,
	generateCandidates factory.CandidatesGeneratorFunction,
) factory.CandidatesGeneratorFunction {
	return func(input string) []string {
		return generateCandidates(sanitize(input))
	}
}

// MaxFromCandidates is a higher-order function which combines
// the generateCandidates function and similarityScore function together.
// Specifically, the newly created string similarity score function
// will generate possible candidates from each input string,
// then the similarity score of all possible pair of candidates
// from both input strings will be computed.
// The returned score will the maximum scores among scores
// between all possible pairs of candidates.
func MaxFromCandidatesProduct(
	generateCandidates factory.CandidatesGeneratorFunction,
	similarityScore factory.SimilarityScoreFunction,
) factory.SimilarityScoreFunction {
	return func(fst, snd string) float64 {
		fstCandidateSet := generateCandidates(fst)
		sndCandidateSet := generateCandidates(snd)
		bestScore := math.Inf(-1)
		for _, candidateA := range fstCandidateSet {
			for _, candidateB := range sndCandidateSet {
				score := similarityScore(candidateA, candidateB)
				bestScore = math.Max(bestScore, score)
			}
		}
		return bestScore
	}
}
