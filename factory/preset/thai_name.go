package preset

import (
	"github.com/abhabongse/fuzzymatch-go/candidate/nametitle"
	"github.com/abhabongse/fuzzymatch-go/editdist"
	editdistThai "github.com/abhabongse/fuzzymatch-go/editdist/thai"
	"github.com/abhabongse/fuzzymatch-go/factory"
	sanitaryThai "github.com/abhabongse/fuzzymatch-go/transform/thai"
)

// ThaiNameSimilarityScore computes the similarity score between two input strings
// with the following functionalities:
// 1.  Each input string will be sanitized via sanitaryThai.Sanitize function
//     (e.g. removing diacritics from latin scripts, removing repeated Thai tonal marks, etc.)
// 2.  Each input string will be used to generate bare names
//     (i.e. attempting to remove English and Thai titles such as Mrs. or dek-chai)
// 3.  For optimal alignment distance metric over string space,
//     the specialized substitution/transposition penalty functions are used instead.
var ThaiNameSimilarityScore = factory.PrependStringSanitizerForSimilarityScore(
	sanitaryThai.Sanitize,
	factory.MaxFromCandidatesProduct(
		nametitle.GenerateNamesWithoutTitles,
		editdist.MakeStringSimilarityFunction(
			editdist.MakeOptimalAlignmentDistFunction(editdistThai.SubstPenalty, editdistThai.TransPenalty),
		),
	),
)
