package candidate

/*
GenerateDefault uses the input string itself as the only candidate.
*/
func GenerateDefault(str string) []string {
	candidates := []string{str}
	return candidates
}
