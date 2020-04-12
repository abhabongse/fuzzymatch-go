package editdist

import (
	"math"
)

// OptimalAlignmentDistance computes the original "Levenshtein distance"
// between two provided string, viewed as a sequence of rune characters.
//
// Character substitution penalty function substPenalty has to be a symmetric function
// which takes in two rune character and outputs their substitution penalty between 0 an 1.
// If two identical runes are provided, ideally the output penalty has to be 0.
// As a special case to account for character insertions and deletions,
// if one of the input runes is rune 0, then substDistFunction should return
// the insertion/deletion penalty of the other rune provided to the function.
//
// Time complexity
//
// The implementation used to compute the optimal alignment distance
// is a simple dynamic programming algorithm that takes O(|fst| * |snd|) time
// where |fst| and |snd| are the lengths of two input strings respectively.
// Additionally, the memory usage for this function is within the order of O(|snd|).
func LevenshteinDist(fst, snd string, substPenalty RunePenaltyFunction) float64 {
	// Convert string into slice of runes
	fstRunes := []rune(fst)
	sndRunes := []rune(snd)

	// Set up the dynamic programming table maintaining
	// only the last 2 rows of the computation.
	// Each rune with zero-index i of the first string
	// corresponds to the row i+1 of the dynamic programming table,
	// whereas each rune with zero-index j of the second string
	// corresponds to the column j+1 of the table.
	// Note that row 0 and column 0 is reserved for empty prefix initializations.
	table := [2][]float64{
		make([]float64, len(sndRunes)+1),
		make([]float64, len(sndRunes)+1),
	}

	// Initialize row 0 of the dynamic programming table
	// by repeatedly inserting runes in order.
	table[0][0] = 0
	for zj, d := range sndRunes {
		// For definitions of d, zj, pj, j, see the next part of the code
		pj, j := resolveColIndex(zj)
		table[0][j] = table[0][pj] + substPenalty(0, d)
	}

	// Fill in the dynamic programming table row-by-row
	for zi, c := range fstRunes {
		pi, i := resolveRowIndex(zi)
		// c = each rune in the first string
		// zi = zero-indexing over the original first string space
		// pi, i = the previous, and the current row index
		//		in the dynamic programming table space

		table[i][0] = table[pi][0] + substPenalty(c, 0)

		for zj, d := range sndRunes {
			pj, j := resolveColIndex(zj)
			// d = each rune in the second string
			// zj = zero-indexing over the original second string space
			// pj, j = the previous, and the current column index
			//		in the dynamic programming table space

			// Compute the minimum score for the rune insertions, deletions,
			// and substitutions only.
			table[i][j] = math.Min(
				table[pi][pj]+substPenalty(c, d),
				math.Min(
					table[pi][j]+substPenalty(c, 0),
					table[i][pj]+substPenalty(0, d),
				),
			)
		}
	}
	// Extract the actual distance score
	_, i := resolveRowIndex(len(fstRunes) - 1)
	_, j := resolveColIndex(len(sndRunes) - 1)
	return table[i][j]
}
