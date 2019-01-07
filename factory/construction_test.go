package factory

import (
	"fmt"
	"github.com/abhabongse/fuzzymatch-go/candidacy"
	"github.com/abhabongse/fuzzymatch-go/canonical"
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
		StringCanonicalization(canonical.CanonicalizeThaiString),
		CandidateGeneration(candidacy.DefaultExtractBareNames),
		CustomEditDistance(addons.NormalizedThaiOptimalAlignmentDistance),
		CustomCombinedScore(func(editDistanceSubScore, diceCoefficientSubScore float64) float64 {
			return 0.5*editDistanceSubScore + 0.5*diceCoefficientSubScore
		}),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("นางสาวสยาม", "สาวสวยสาม")
	fmt.Println(score)
}

func ExampleCompactedCustomizedSimilarityScore() {
	// Similar to ExampleCustomizedSimilarityScore but with compacted option shortcuts
	SimilarityScore := NewSimilarityScoreFunction(
		StringCanonicalization(canonical.CanonicalizeThaiString),
		CandidateGeneration(candidacy.DefaultExtractBareNames),
		OptimalAlignmentEditDistance(addons.ThaiSubstitutionErrorDist, addons.ThaiTranspositionErrorDist),
		LinearCombinedScore(1.0, 1.0),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("นางสาวสยาม", "สาวสวยสาม")
	fmt.Println(score)
}