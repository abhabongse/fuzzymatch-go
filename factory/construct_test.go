package factory

import (
	"fmt"
	"github.com/abhabongse/fuzzymatch-go/candidate"
	"github.com/abhabongse/fuzzymatch-go/legacy_editdist"
	"github.com/abhabongse/fuzzymatch-go/legacy_editdist/extra"
	"github.com/abhabongse/fuzzymatch-go/sanitary"
	extralang2 "github.com/abhabongse/fuzzymatch-go/sanitary/extra"
)

func ExampleSimilarityScore() {
	// Create a default string similarity score function
	SimilarityScore := NewSimilarityScoreFunc()

	// Alternatively, Options can be supplemented to configure the function.
	// These are the default Options which yield identical behavior as above.
	SimilarityScore = NewSimilarityScoreFunc(
		StringSanitization(sanitary.Noop),
		CandidatesGeneration(candidate.GenerateDefault),
		OptimalAlignmentEditDistance(legacy_editdist.UnitPenalty, legacy_editdist.UnitPenalty),
		LinearCombinedScore(1.0, 0.0),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("saturday", "sunday")
	fmt.Println(score)
}

func ExampleCustomizedSimilarityScore() {
	// Create a string similarity score function that can handle discrepancies in the existence
	// of salutation titles and that enforces non-unit distance penalties.
	SimilarityScore := NewSimilarityScoreFunc(
		StringSanitization(extralang2.ThaiSanitize),
		CandidatesGeneration(candidate.NamesWithoutTitles),
		CustomEditDistance(extra.ThaiOptimalAlignmentNormDist),
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
	SimilarityScore := NewSimilarityScoreFunc(
		StringSanitization(extralang2.ThaiSanitize),
		CandidatesGeneration(candidate.NamesWithoutTitles),
		OptimalAlignmentEditDistance(extra.ThaiSubstPenalty, extra.ThaiTransPenalty),
		LinearCombinedScore(1.0, 1.0),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("นางสาวสยาม", "สาวสวยสาม")
	fmt.Println(score)
}
