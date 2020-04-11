package factory

// String sanitization function to be applied to each input string
type StringSanitizerFunction = func(string) string

// Candidate generation function to be applied to each input string
type CandidatesGeneratorFunction = func(string) []string

// Edit distance function to be applied to both strings
type SimilarityComputerFunction = func(string, string) float64
