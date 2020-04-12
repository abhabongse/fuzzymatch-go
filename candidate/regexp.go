package candidate

import "regexp"

// MakeRegexpGenerator constructs a function which attempts to extract
// all possible text from the input string provided as the only input argument.
// Input strings will be parsed according to the given regular expression patterns.
func MakeRegexpGenerator(patterns []*regexp.Regexp) func(string) []string {
	return func(input string) []string {
		return generateCandidateWithRegexps(input, patterns)
	}
}

// generateCandidateWithRegexps uses each of the provided regular expression patterns
// to match and extract a given text from a given string.
// The output is a sorted slice of all extract strings that matched one of the patterns.
// Unmatched patterns will be skipped.
// Only the first matched group from each pattern is extracted.
func generateCandidateWithRegexps(input string, patterns []*regexp.Regexp) []string {
	results := make([]string, 0)
	for _, pattern := range patterns {
		r := pattern.FindStringSubmatch(input)
		if len(r) > 1 {
			results = append(results, r[1])
		}
	}
	return results
}
