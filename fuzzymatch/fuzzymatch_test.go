package fuzzymatch

import (
	"fmt"
	"github.com/abhabongse/fuzzymatch-go/candidacy"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
)

func ExampleSimilarityScore() {
	// Create a default string similarity score function
	SimilarityScore := NewSimilarityScoreFunction()

	// Alternatively, Options can be supplemented to configure the function.
	// These are the default Options which yield identical behavior as above.
	SimilarityScore = NewSimilarityScoreFunction(
		StringCanonicalization(DefaultCanonicalizeString),
		CandidateGeneration(DefaultGenerateCandidates),
		OptimalAlignmentEditDistance(editdistance.UnitDist, editdistance.UnitDist),
		LinearCombinedScore(1.0, 0.0),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("saturday", "sunday")
	fmt.Println(score)
}

func ExampleCustomizedSimilarityScore() {
	// Create a string similarity score function that can handle discrepancies in the existence
	// of salutation titles and that enforces non-unit distance penalties.
	SimilarityScore := NewSimilarityScoreFunction(
		CandidateGeneration(candidacy.ExtractBareNames),
		OptimalAlignmentEditDistance(editdistance.ThaiSubstitutionErrorDist, editdistance.ThaiTranspositionErrorDist),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("นางสาวสยาม", "สาวสวยสาม")
	fmt.Println(score)
}
