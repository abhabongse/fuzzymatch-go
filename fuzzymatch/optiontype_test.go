package fuzzymatch

import (
	"fmt"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
)

func ExampleSimilarityScore() {
	var SimilarityScore func(string, string) float64
	// Construct a basic string similarity score function (these lines are equivalent)
	SimilarityScore = NewSimilarityScoreFunction()
	SimilarityScore = NewSimilarityScoreFunction(
		StringCanonicalization(DefaultCanonicalizeString),
		CandidateGeneration(DefaultGenerateCandidates),
		RuneDistancePenalties(editdistance.UnitDist, editdistance.UnitDist),
		CombinationWeights(1.0, 0.0),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("saturday", "sunday")
	fmt.Println(score)
}
