/*
Package fuzzymatch implements an approximate string matching algorithm to determine the similarity
score between two given strings. In most cases, you would be interested in the function SimilarityScore.
*/
package fuzzymatch

/*
A definitive function which computes the similarity score between two input strings.
The returned score value is a floating point between 0 (strings are very distinct)
and 1 (strings are identical).
*/
func SimilarityScore(fst, snd string) float64 {
	fstRunes := normalizeString(fst)
	sndRunes := normalizeString(snd)

	// Breaking ties to save memory in the long run
	if len(fstRunes) < len(sndRunes) {
		fstRunes, sndRunes = sndRunes, fstRunes
	}

	return 0
}

/*
Normalize an input string using various methods and return as a slice of runes.
*/
func normalizeString(str string) []rune {
	// TODO: introduce multiple string normalization functions
	str = NormalizeWhiteSpaces(str)
	// ...

	runesData := []rune(str)
	return runesData
}
