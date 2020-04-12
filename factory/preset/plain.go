package preset

import (
	"github.com/abhabongse/fuzzymatch-go/transform"
)

// PlainSimilarityScore computes the similarity score between two input strings
// but each input string will be sanitized before they are compared to each other.
var PlainSimilarityScore = transform.PrependStringSanitizer(
	transform.LatinExtendedSanitize,
	SimpleSimilarityScore,
)
