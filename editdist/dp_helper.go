package editdist

// resolveRowIndex3 is a helper function which turns a zero-based index
// over the original string space into a triplet of indices
// corresponding to the before previous, the previous, and the current row
// of the dynamic programming table from the function OptimalAlignmentDist, respectively.
// Returned indices are in modulus 3.
func resolveRowIndex(index int) (int, int) {
	return index % 2, (index + 1) % 2
}

// resolveColIndex3 is a helper function which turns a zero-based index
// over the original string space into a triplet of indices
// corresponding to the before previous, the previous, and the current column
// of the dynamic programming table from the function OptimalAlignmentDist, respectively.
func resolveColIndex(index int) (int, int) {
	return index, index + 1
}

// resolveRowIndex3 is a helper function which turns a zero-based index
// over the original string space into a triplet of indices
// corresponding to the before previous, the previous, and the current row
// of the dynamic programming table from the function OptimalAlignmentDist, respectively.
// Returned indices are in modulus 3.
func resolveRowIndex3(index int) (int, int, int) {
	return (index - 1) % 3, index % 3, (index + 1) % 3
}

// resolveColIndex3 is a helper function which turns a zero-based index
// over the original string space into a triplet of indices
// corresponding to the before previous, the previous, and the current column
// of the dynamic programming table from the function OptimalAlignmentDist, respectively.
func resolveColIndex3(index int) (int, int, int) {
	return index - 1, index, index + 1
}
