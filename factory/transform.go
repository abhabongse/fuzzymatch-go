package factory

import "golang.org/x/text/transform"

// MakeStringTransformFunction is a higher-order function which constructs
// a new string transforming function from a transformer object.
func MakeStringTransformFunction(transformer transform.Transformer) StringTransformerFunction {
	return func(input string) string {
		output, _, _ := transform.String(transformer, input)
		return output
	}
}

// PrependStringSanitizerForSimilarityScore is a higher-order function which modifies
// the provided similarityScore function so that each input string to the function
// will be sanitized first via a call to sanitize function.
func PrependStringSanitizerForSimilarityScore(
	sanitize StringTransformerFunction,
	similarityScore SimilarityScoreFunction,
) SimilarityScoreFunction {
	return func(fst, snd string) float64 {
		return similarityScore(sanitize(fst), sanitize(snd))
	}
}
