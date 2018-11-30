package editdistance

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
