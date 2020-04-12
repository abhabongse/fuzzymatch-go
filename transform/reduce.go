package transform

import (
	"github.com/abhabongse/fuzzymatch-go/factory"
)

// PrependStringSanitizer is a higher-order function
// which modifies the provided similarityScore function
// so that each input string to the function will be sanitized first
// via a call to sanitize function.
func PrependStringSanitizer(
	sanitize factory.StringTransformerFunction,
	similarityScore factory.SimilarityScoreFunction,
) factory.SimilarityScoreFunction {
	return func(fst, snd string) float64 {
		return similarityScore(sanitize(fst), sanitize(snd))
	}
}
