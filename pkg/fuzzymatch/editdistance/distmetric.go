package editdistance

/*
SubstDistMetric is a type representing the distance score (penalty) function for
edit operations, particularly substitutions as well as insertions and deletions.
The function takes in two runes as input arguments and returns the penalty value
between 0 (meaning that two runes are identical) and 1 (meaning that two runes
are totally distinct).
*/
type SubstDistMetric = func(rune, rune) float64

/*
TransDistMetric is a type representing the distance score (penalty) function for
edit operations, particularly transpositions of adjacent characters. The function
takes in two runes as input arguments and returns the penalty value between 0
(meaning that two runes are highly likely to be transposed) and 1 (meaning that
two runes are hardly ever transposed).
*/
type TransDistMetric = func(rune, rune) float64

/*
UnitDist is a indicator function which returns 1 if two input characters differ and
returns 0 otherwise. To represent the distance score (or the penalty) for insertion
and deletion operations, we use the rune character 0 to represent the absence of a
character input. This function is guaranteed to be symmetric (i.e. the order of
arguments do not matter).

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
