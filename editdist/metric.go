package editdist

// StringDistFunction is a type alias for edit distance functions between two strings
type StringDistFunction = func(string, string) float64

// RunePenaltyFunction is an umbrella type alias representing a _symmetric_ function
// which computes a penalty score in the context of two rune characters.
// There are two kinds of penalties: substitution penalty and transposition penalty.
//
// Substitution Penalty
//
// Substitution Penalty refers to the cost of substituting one rune for the other
// (which includes character insertions and deletions as special cases).
// Specifically, a function of this kind takes in two runes as input arguments
// and returns the penalty value within 0 to 1 range,
// where 0 indicates that two runes are identical and 1 indicates that
// two rune are totally distinct and can never be mistakenly mixed up.
// The insertion and deletion penalties are represented by assigning rune 0
// as one of the input arguments to the function.
//
// Transposition Penalty
//
// Transposition Penalty refers to the cost of adjacent character transpositions
// (i.e. swapping location of two adjacent characters in the string).
// Specifically, a function of this kind takes in two runes as input arguments
// and returns the penalty value within 0 to 2 range,
// where 0 indicates that both runes are commonly and indistinctively interchangeable
// and 2 indicates that both runes are completely different
// (and thus the penalty is no different than performing two consecutive substitutions).
type RunePenaltyFunction = func(rune, rune) float64

// UnitPenalty is an indicator function which returns 1
// if two input runes differ and returns 0 otherwise.
// This function can be used for both the substitution and the transposition penalties.
//
// UnitPenalty is a symmetric function by nature,
// meaning that the order of the input arguments do not matter.
//
// For example, UnitPenalty('a', 'b') is the penalty for substituting character 'a'
// with another character 'b' (the output of which is 1 in this particular case).
// On the other hand, UnitPenalty('c', 0) is the penalty for an insertion (or deletion)
// of the character 'c'.
func UnitPenalty(c, d rune) float64 {
	if c == d {
		return 0
	}
	return 1
}
