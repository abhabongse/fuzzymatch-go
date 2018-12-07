package editdistance

/*
RuneDistanceMetric is an umbrella type representing a symmetric function which
computes a distance score (penalty) between a given pair of runes. There are two
different kinds of distance metrics: substitution penalty and transposition penalty.


Substitution penalty

Substitution penalty refers to the cost of character substitutions (which includes
insertions and deletions as special cases). Specifically, a function of this kind
takes in two runes as input arguments and returns the penalty value between 0
(meaning that two runes are identical) and 1 (meaning that two runes are totally
distinct). The insertion and deletion penalties are represented by assigning byte 0
as one of the input argument to the function.


Transposition penalty

Transposition penalty refers to the cost of adjacent character transpositions.
Specifically, a function of this kind takes in two runes as input argument and
returns the penalty value between 0 (meaning that both runes are interchangeable)
and 2 (meaning that two runes are hardly mistaken with each other).
*/
type RuneDistanceMetric = func(rune, rune) float64

/*
UnitDist is an indicator function which returns 1 if two input characters differ
and returns 0 otherwise. This function can be used for both the substitution and the
transposition penalties.

UnitDist is a symmetric function, meaning that the order of the input arguments do
not matter.

For example, UnitDist('a', 'b') is the penalty for substituting character 'a' with
another character 'b' (which is 1 in this case) whereas UnitDist('c', 0) is the
penalty for an insertion or a deletion of character 'c'.
*/
func UnitDist(c, d rune) float64 {
	if c == d {
		return 0
	} else {
		return 1
	}
}

/*
ErrorTolerantSubstDist is a function returning substitution penalties with values
between 0 and 1. If erroneous substitutions between rune characters c and d are
more likely, then the penalty will be smaller.
*/
func ErrorTolerantSubstDist(c, d rune) float64 {
	return UnitDist(c, d)
}

/*
ErrorTolerantSubstDist is a function returning transposition penalties with values
between 0 and 2. If erroneous transpositions between rune characters c and d are
more likely, then the penalty will be smaller.
*/
func ErrorTolerantTransDist(c, d rune) float64 {
	return UnitDist(c, d)
}
