package factory

import (
	"fmt"
	"github.com/abhabongse/fuzzymatch-go/candidate/nametitle"
	"github.com/abhabongse/fuzzymatch-go/editdist"
	editdistThai "github.com/abhabongse/fuzzymatch-go/editdist/thai"
	sanitaryThai "github.com/abhabongse/fuzzymatch-go/sanitary/thai"
)

func ExampleCustomizedSimilarityScore() {
	// Create a default string similarity score function
	SimilarityScore := MakeSimilarityScoreFunction()

	// Alternatively, Options can be supplemented to configure the function.
	thaiOptimalAlignmentDist := editdist.MakeOptimalAlignmentDistFunction(
		editdistThai.SubstPenalty, editdistThai.TransPenalty,
	)
	thaiStringSimilarity := editdist.MakeStringSimilarityFunction(thaiOptimalAlignmentDist)
	SimilarityScore = MakeSimilarityScoreFunction(
		StringSanitizer(sanitaryThai.Sanitize),
		CandidatesGenerator(nametitle.GenerateNamesWithoutTitles),
		SimilarityComputer(thaiStringSimilarity),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("saturday", "sunday")
	fmt.Println(score)
}
