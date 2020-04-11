package factory

// Options is a type struct that stores the configuration information
// regarding how to compute string similarity score between two input strings,
// particularly,
// (1) how strings are sanitized,
// (2) how variants/candidates are generates,
// (3) how edit distances between both strings are computed, and
// (4) how edit distance sub-score and the Dice coefficient sub-score
//     will be combined to compute the final score, etc.
type Options struct {
	// String sanitization function to be applied to each input string
	sanitizeString StringSanitizerFunction
	// Candidate generation function to be applied to each input string
	generateCandidates CandidatesGeneratorFunction
	// Edit distance function to be applied to both strings
	computeSimilarity SimilarityComputerFunction
}

// OptionSetter is a type alias for function that modifies the Options type struct.
// Functions of this type can be used to configure
// how to compute the overall string similarity scores
// between any two input strings.
type OptionSetter = func(*Options)

// StringSanitizer assigns the function that would be used
// to clean up each of the input strings before they are subsequently compared.
// Specifically, this wrappedFunc will receive a string as the only input argument
// and it should return the sanitized string of the input.
func StringSanitizer(sanitizeString StringSanitizerFunction) OptionSetter {
	return func(config *Options) {
		config.sanitizeString = sanitizeString
	}
}

// CandidatesGenerator assigns the function that would be used
// to generate all normalization variants of the already-sanitized input string.
// Specifically, this function will receive a string as the only input argument
// and it should return a slice of strings each indicating a possible variant of the input string.
func CandidatesGenerator(generateCandidates CandidatesGeneratorFunction) OptionSetter {
	return func(config *Options) {
		config.generateCandidates = generateCandidates
	}
}

// SimilarityComputer assigns the given function as the function
// to compute the edit distance between two input strings.
// Warning: this function should return 1 when both strings are identical
// and should return 0 when both string are totally distinct.
func SimilarityComputer(computeSimilarity SimilarityComputerFunction) OptionSetter {
	return func(config *Options) {
		config.computeSimilarity = computeSimilarity
	}
}
