package factory

// SimilarityScoreFunction is a type alias describing
// the signature of a function that consumes two input strings
// and computes their string similarity score
// whose value is within the range from 0 to 1
// where 1 indicates that both strings are semantically identical
// and 0 indicates that both strings are totally distinct.
type SimilarityScoreFunction = func(string, string) float64

// StringTransformer is a type alias describing
// the signature of a function that transforms a string into other forms.
// This type is useful to describe string sanitization functions, etc.
type StringTransformerFunction = func(string) string

// CandidatesGeneratorFunction is a type alias describing
// the signature of a function that generates multiple candidates
// where each candidate is a possible true value of the input string.
//
// For example, there may be a string that attempts to strip all titles
// from the name given as the string input of the function.
// However, we cannot be certain if some of the titles are
// actually parts of the real name and not the title itself.
type CandidatesGeneratorFunction = func(string) []string
