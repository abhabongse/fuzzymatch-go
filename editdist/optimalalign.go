// This source file contains the implementation of optimal alignment distance metric.

package editdist

import "math"

// OptimalAlignmentDistance computes the "optimal alignment distance"
// between two provided string, viewed as a sequence of rune characters.
//
// Character substitution penalty function substPenalty has to be a symmetric function
// which takes in two rune character and outputs their substitution penalty between 0 an 1.
// If two identical runes are provided, ideally the output penalty has to be 0.
// As a special case to account for character insertions and deletions,
// if one of the input runes is rune 0, then substDistFunction should return
// the insertion/deletion penalty of the other rune provided to the function.
//
// Adjacent character transposition penalty function transPenalty also has to be
// a symmetric function which takes in two non-zero rune characters
// and outputs their transposition penalty score between 0 and 2.
// Please note that if for a pair of runes, x and y, satisfies the equation
//     transDistFunction(x, y) >= 2 * substDistFunction(x, y)
// then the transposition of those characters are effectively ignored
// since it is cheaper to just perform two substitutions in a row instead.
//
// Time complexity
//
// The implementation used to compute the optimal alignment distance
// is a simple dynamic programming algorithm that takes O(|fst| * |snd|) time
// where |fst| and |snd| are the lengths of two input strings respectively.
// Additionally, the memory usage for this function is within the order of O(|snd|).
func OptimalAlignmentDist(fst, snd string, substPenalty, transPenalty RunePenaltyFunction) float64 {
	// Convert string into slice of runes
	fstRunes := []rune(fst)
	sndRunes := []rune(snd)

	// Set up the dynamic programming table maintaining
	// only the last 3 rows of the computation.
	// Each rune with zero-index i of the first string
	// corresponds to the row i+1 of the dynamic programming table,
	// whereas each rune with zero-index j of the second string
	// corresponds to the column j+1 of the table.
	// Note that row 0 and column 0 is reserved for empty prefix initializations.
	table := [3][]float64{
		make([]float64, len(sndRunes)+1),
		make([]float64, len(sndRunes)+1),
		make([]float64, len(sndRunes)+1),
	}

	// Initialize row 0 of the dynamic programming table
	// by repeatedly inserting runes in order.
	table[0][0] = 0
	for zj, d := range sndRunes {
		// For definitions of d, zj, pj, j, see the next part of the code
		_, pj, j := resolveColIndex3(zj)
		table[0][j] = table[0][pj] + substPenalty(0, d)
	}

	// Fill in the dynamic programming table row-by-row
	var pc rune = 0 // rune from the previous row, initialized empty
	for zi, c := range fstRunes {
		ppi, pi, i := resolveRowIndex3(zi)
		// c = each rune in the first string
		// zi = zero-indexing over the original first string space
		// ppi, pi, i = the before previous, the previous, and the current
		//     row index in the dynamic programming table space

		table[i][0] = table[pi][0] + substPenalty(c, 0)

		var pd rune = 0 // rune from the previous column, initialized empty
		for zj, d := range sndRunes {
			ppj, pj, j := resolveColIndex3(zj)
			// d = each rune in the second string
			// zj = zero-indexing over the original second string space
			// ppj, pj, j = the before previous, the previous, and the current
			//     column index in the dynamic programming table space

			// Compute the minimum score for the rune insertions, deletions,
			// and substitutions only.
			table[i][j] = math.Min(
				table[pi][pj]+substPenalty(c, d),
				math.Min(
					table[pi][j]+substPenalty(c, 0),
					table[i][pj]+substPenalty(0, d),
				),
			)
			// Now consider the transposition penalty if the last two runes
			// of both strings are indeed transposition.
			if pc == d && pd == c {
				table[i][j] = math.Min(
					table[i][j],
					table[ppi][ppj]+transPenalty(c, d),
				)
			}
			pd = d
		}
		pc = c
	}
	// Extract the actual distance score
	_, _, i := resolveRowIndex3(len(fstRunes) - 1)
	_, _, j := resolveColIndex3(len(sndRunes) - 1)
	return table[i][j]
}
