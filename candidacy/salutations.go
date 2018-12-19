package candidacy

/*
Decomposition is a type struct for the split between the salutation and
the bare name sans the salutation.
*/
type Decomposite = struct{ salute, bare string }

/*
GenerateSalutationDecomposites generates a sequence of all possible splits
(i.e. decomposites) of salutation from a given name input.

For this function, only basic English and Thai salutations are concerned:
mr, mrs, ms, mister, miss, master, นาย, นาง, นางสาว, เด็กชาย, เด็กหญิง.
*/
func GenerateSalutationDecomposites(name string) []Decomposite {
	candidates := make([]Decomposite, 0)
	candidates = append(candidates, Decomposite{"", name})

	// TODO: Check for English salutations

	// TODO: Check for Thai salutations

	return candidates
}

var basicEnglishSalutations = []string{
	"mr",
	"mrs",
	"ms",
	"mister",
	"miss",
	"master",
}

var basicThaiSalutations = []string{
	"นาย",
	"นาง",
	"นางสาว",
	"เด็กชาย",
	"เด็กหญิง",
}
