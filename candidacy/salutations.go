package candidacy

import (
	"regexp"
	"sort"
)

/*
Decomposition is a type struct representing the splitting between the salutation
title (Salute) and the bare name sans the salutation (Bare).
*/
type Decomposite = struct{ Salute, Bare string }

/*
DecomposeName attempts to parse the given name string using various decomposition
patterns. The output of this function is a slice of Decomposite structures in
the sorted order of the attributes Salute and Bare, respectively.
*/
func DecomposeName(patterns []*regexp.Regexp, name string) []Decomposite {
	// One possible candidate is the empty salutation (always assumed)
	decomposites := []Decomposite{{"", name}}

	// Attempt to parse name with each pattern
	for _, pattern := range patterns {
		result := pattern.FindStringSubmatch(name)
		if len(result) > 0 {
			decomposites = append(decomposites, Decomposite{result[1], result[2]})
		}
	}

	// Sort the output sequence using the composite key (Salute, Bare)
	sort.Slice(decomposites,
		func(i, j int) bool {
			return decomposites[i].Salute < decomposites[j].Salute || decomposites[i].Salute == decomposites[j].Salute && decomposites[i].Bare < decomposites[j].Bare
		},
	)
	return decomposites
}

/*
ExtractBareNames is a preset function which generates a sequence of all possible
bare names (or names without salutation titles). It is built upon the function
DecomposeName with the patterns from DefaultSalutationPatterns.
*/
func ExtractBareNames(name string) []string {
	salutationDecomposites := DecomposeName(DefaultSalutationPatterns, name)
	bareNames := make([]string, 0, len(salutationDecomposites))
	for _, decomposite := range salutationDecomposites {
		bareNames = append(bareNames, decomposite.Bare)
	}
	return bareNames
}

/*
DefaultSalutationPatterns is a sequence of all compiled regular expression
objects which separates salutation titles from the actual full name part.

As of current, only common English and Thai salutation titles are handled.
*/
var DefaultSalutationPatterns = []*regexp.Regexp{
	// English full salutation titles: space separator required
	regexp.MustCompile("^(mister)(?: )(.*)$"),
	regexp.MustCompile("^(miss)(?: )(.*)$"),
	regexp.MustCompile("^(master)(?: )(.*)$"),
	// English abbreviated salutation titles: separator required
	regexp.MustCompile("^(mr)(?: |\\. |\\.)(.*)$"),
	regexp.MustCompile("^(mrs)(?: |\\. |\\.)(.*)$"),
	regexp.MustCompile("^(ms)(?: |\\. |\\.)(.*)$"),
	// Thai full salutation titles: separator optional
	regexp.MustCompile("^(นาย)(?: ?)(.*)$"),
	regexp.MustCompile("^(นาง)(?: ?)(.*)$"),
	regexp.MustCompile("^(นางสาว)(?: ?)(.*)$"),
	regexp.MustCompile("^(เด็กชาย)(?: ?)(.*)$"),
	regexp.MustCompile("^(เด็กหญิง)(?: ?)(.*)$"),
	// Thai abbreviated salutation titles: space separator required
	regexp.MustCompile("^(ดช)(?: )(.*)$"),
	regexp.MustCompile("^(ดญ)(?: )(.*)$"),
	// Thai dot-abbreviated salutation titles: separator optional
	regexp.MustCompile("^(ด\\.ช\\.)(?: ?)(.*)$"),
	regexp.MustCompile("^(ด\\.ญ\\.)(?: ?)(.*)$"),
}
