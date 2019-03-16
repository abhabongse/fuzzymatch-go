package candidate

import (
	"regexp"
	"sort"
)

// TODO: we do not care about prefixes and suffixes; leave only the stem

/*
NameAndTitles is a type struct representing the splitting between the actual
bare name from the prefixed and suffixed titles.
*/
type NameAndTitles = struct{ Bare, Prefix, Suffix string }

/*
SplitTitles uses the provided regular expression patterns to parse the given
name in order to generate all possibilities to split prefixed and suffixed
titles. The output is a slice of NameAndTitles structures in the sorted
order of the composite key (Bare, Prefix, Suffix).
*/
func SplitTitles(patterns []*regexp.Regexp, name string) []NameAndTitles {

	// One possible candidate is empty titles (always assumed)
	results := []NameAndTitles{{name, "", ""}}

	// Attempt to parse name with each pattern
	for _, pattern := range patterns {
		r := pattern.FindStringSubmatch(name)
		if len(r) > 0 {
			results = append(results, NameAndTitles{r[2], r[1], r[3]})
		}
	}

	// Sort the output sequence using the composite key (Bare, Prefix, Suffix)
	sort.Slice(results,
		func(i, j int) bool {
			return results[i].Bare < results[j].Bare || results[i].Bare == results[j].Bare &&
				(results[i].Prefix < results[j].Prefix || results[i].Prefix == results[j].Prefix &&
					results[i].Suffix < results[j].Suffix)
		},
	)
	return results
}

/*
MakeBareNameExtractor constructs a function which attempts to extract all possible
bare names (names without prefixed and suffixed titles) from a name provided as
the only input argument. Input names will be parsed according to the given regular
expression patterns.
*/
func MakeBareNamesExtractor(patterns []*regexp.Regexp) func(string) []string {
	return func(name string) []string {
		tokenizedCandidates := SplitTitles(patterns, name)
		bareNames := make([]string, 0, len(tokenizedCandidates))
		for _, candidate := range tokenizedCandidates {
			bareNames = append(bareNames, candidate.Bare)
		}
		return bareNames
	}
}

/*
NamesWithoutTitles is a preset function which generates a sequence of all possible
bare names (or names without prefixed or suffixed titles). It is built upon the function
DecomposeName with the patterns from DefaultTitledNamePatterns.
*/
var NamesWithoutTitles = MakeBareNamesExtractor(DefaultTitledNamePatterns)

/*
DefaultTitledNamePatterns is a sequence of all compiled regular expression objects which
separates bare names from their prefixed and suffixed titles in a full name string.

As of current, only common English and Thai prefixed titles are handled.
*/
var DefaultTitledNamePatterns = []*regexp.Regexp{
	// English full prefixed titles: space separator required
	regexp.MustCompile("^(mister)(?: )(.*)()$"),
	regexp.MustCompile("^(miss)(?: )(.*)()$"),
	regexp.MustCompile("^(master)(?: )(.*)()$"),
	// English abbreviated prefixed titles: separator required
	regexp.MustCompile("^(mr)(?: |\\. |\\.)(.*)()$"),
	regexp.MustCompile("^(mrs)(?: |\\. |\\.)(.*)()$"),
	regexp.MustCompile("^(ms)(?: |\\. |\\.)(.*)()$"),
	// Thai full prefixed titles: separator optional
	regexp.MustCompile("^(นาย)(?: ?)(.*)()$"),
	regexp.MustCompile("^(นาง)(?: ?)(.*)()$"),
	regexp.MustCompile("^(นางสาว)(?: ?)(.*)()$"),
	regexp.MustCompile("^(เด็กชาย)(?: ?)(.*)()$"),
	regexp.MustCompile("^(เด็กหญิง)(?: ?)(.*)()$"),
	// Thai abbreviated prefixed titles: space separator required
	regexp.MustCompile("^(ดช)(?: )(.*)()$"),
	regexp.MustCompile("^(ดญ)(?: )(.*)()$"),
	// Thai dot-abbreviated prefixed titles: separator optional
	regexp.MustCompile("^(ด\\.ช\\.)(?: ?)(.*)()$"),
	regexp.MustCompile("^(ด\\.ญ\\.)(?: ?)(.*)()$"),
}
