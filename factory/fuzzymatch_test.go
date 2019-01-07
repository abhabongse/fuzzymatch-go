package factory

import (
	"fmt"
	"github.com/abhabongse/fuzzymatch-go/candidacy"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
	"github.com/abhabongse/fuzzymatch-go/editdistance/addons"
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
		CandidateGeneration(candidacy.DefaultExtractBareNames),
		OptimalAlignmentEditDistance(addons.ThaiSubstitutionErrorDist, addons.ThaiTranspositionErrorDist),
		CustomCombinedScore(func(editDistanceSubScore, diceCoefficientSubScore float64) float64 {
			// Multiply blend-mode
			return 0.5 * editDistanceSubScore * diceCoefficientSubScore +
				0.25 * editDistanceSubScore +
				0.25 * diceCoefficientSubScore
		}),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("นางสาวสยาม", "สาวสวยสาม")
	fmt.Println(score)
}
