package preset

import (
	"github.com/abhabongse/fuzzymatch-go/editdist"
)

// SimpleSimilarityScore computes the similarity score between two input strings.
// Two input strings will be directly compared under optimal alignment distance metric
// without any pre-processing, and the resulting distance will be re-normalized to
// a similarity score between 0 and 1 (inclusive).
var SimpleSimilarityScore = editdist.MakeStringSimilarityFunction(
	editdist.MakeOptimalAlignmentDistFunction(editdist.UnitPenalty, editdist.UnitPenalty),
)
