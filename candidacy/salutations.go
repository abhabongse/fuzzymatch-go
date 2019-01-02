package candidacy

import (
	"regexp"
	"sort"
)

/*
Decomposition is a type struct for the split between the salutation and
the bare name sans the salutation.
*/
type Decomposite = struct{ salute, bare string }

/*
GenerateSalutationDecomposites generates a sequence of all possible splits
(i.e. decomposites) of salutation from a given name input. Output sequences
are sorted according to the extracted salutations first.

For this function, only basic English and Thai salutations are concerned:
mr, mrs, ms, mister, miss, master, นาย, นาง, นางสาว, เด็กชาย, เด็กหญิง.
*/
func GenerateSalutationDecomposites(name string) []Decomposite {
	candidates := make([]Decomposite, 0)
	candidates = append(candidates, Decomposite{"", name})

	for _, regex := range SalutationTitleRegExps {
		result := regex.FindStringSubmatch(name)
		if len(result) > 0 {
			candidates = append(candidates, Decomposite{result[1], result[2]})
		}
	}

	// Sort sequences by salutations first, then by bare name
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].salute < candidates[j].salute || candidates[i].salute == candidates[j].salute && candidates[i].bare < candidates[j].bare
	})
	return candidates
}

/*
SalutationTitleRegExps is a sequence of all compiled regular expression
objects which separates salutation titles from the actual full name part.
*/
var SalutationTitleRegExps = []*regexp.Regexp{
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
