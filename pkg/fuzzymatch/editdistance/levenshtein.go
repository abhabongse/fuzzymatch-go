package editdistance

import "math"

/*
SimpleAlignmentDistance is a simplified version of the OptimalAlignmentDistance function.
It assumes that all edit operations (insertions, deletions, substitutions, and adjacent
character transpositions) will incur unit penalties.
*/
func SimpleAlignmentDistance(fst, snd string) float64 {
	return OptimalAlignmentDistance(fst, snd, UnitDist, UnitDist)
}

/*
OptimalAlignmentDistance computes the "optimal alignment distance" between two given
string, viewed as a sequence of rune characters.

An optimal alignment distance is a restriction over the Damerau–Levenshtein distance,
which in turn is a generalization of the original Levenshtein distance.


Levenshtein distance

In the original original Levenshtein distance metric, the distance between two given
strings is measured by the minimum "edit operations" required to transform one string
into the other. Those edit operations is limited to (1) an insertion of a character,
(2) a deletion of a character, and (3) a substitution of a character by another
character.

The penalty for these edit operations is described by the input argument called the
substDistFunc function; this function has to be a symmetric function that takes
in two rune characters and outputs their substitution penalty score between 0 and 1.
If two identical runes are presented, the output penalty has to be 0. As a special case,
if one of the input argument is rune 0, then substDistFunc should return the
insertion or the deletion penalty of the other rune character. For an example of what to
expect from the substDistFunc, see the distance metric function UnitDist.

For example, under the UnitDist character substitution metric, the Levenshtein distance
between "Hello" and "Hola" is 3, whereas the Levenshtein distance between "Hi" and ""
(empty string) is 2.


Damerau–Levenshtein distance

The Damerau–Levenshtein distance was improved upon the original Levenshtein distance
by allowing another kind of edit operation: a transposition of two adjacent characters.

The penalty for this particular edit operations is governed by another input argument
called the transDistFunc function. This function also has to be a symmetric
function that takes in two rune characters and outputs their transposition penalty
score between 0 and 2. Note that if two runes, x and y, satisfy the equation
    transDistFunc(x, y) >= 2 * substDistFunc(x, y)
then transpositions are effectively disallowed for these two particular rune
characters, x and y, since it would be cheaper to perform two substitutions instead.

For example, under the UnitDist character substitution and transposition metrics, the
Damerau–Levenshtein distance between "Thrust" and "Thursday" is 4. Without
transposition, the Levenshtein distance would have become 5.


Optimal alignment distance

The optimal alignment distance is the restricted version of the Damerau–Levenshtein
distance; specifically, each rune character in both the original and the target string
is subjected to at most one edit operation, and only characters that are adjacent in
the original input are allowed to be transposed.

Therefore, Damerau–Levenshtein distance between "trout" and "turn" is 3 (as shown by
"trout" → "trut" → "turt" → "turn"), but the second operation would not be allowed
in the optimal alignment metric, and thus yield a worse distance of 4 (as shown by
"trout" → "tout" → "tut" → "turt" → "turn").


Time complexity

This algorithm we used to compute the optimal alignment distance is a dynamic
programming algorithm that takes O(|fst| * |snd|) where |fst| and |snd| are the length
of two input strings, respectively. Additionally, the memory usage for this function
is within the order of O(|snd|).
*/
func OptimalAlignmentDistance(
	fst, snd string,
	substDistFunc, transDistFunc RuneDistanceMetric,
) float64 {
	// Convert string into slice of runes
	fstRunes := []rune(fst)
	sndRunes := []rune(snd)

	// Set up the dynamic programming table maintaining only the last 3 rows
	// of the computation. Each rune with zero-index i of the first string
	// corresponds to the row i+1 of the dynamic programming table, whereas
	// each rune with zero-index j of the second string corresponds to the
	// column j+1 of the table. Note that row 0 and column 0 is reserved for
	// empty prefix initializations.
	table := [3][]float64{
		make([]float64, len(sndRunes)+1),
		make([]float64, len(sndRunes)+1),
		make([]float64, len(sndRunes)+1),
	}

	// Initialize row 0 of the dynamic programming table by repeatedly
	// inserting runes in order.
	table[0][0] = 0
	for zj, d := range sndRunes {
		// For definitions of d, zj, pj, j, see the next part of the code
		_, pj, j := resolveColIndex(zj)
		table[0][j] = table[0][pj] + substDistFunc(0, d)
	}

	// Fill in the dynamic programming table row-by-row
	var pc rune = 0 // rune from the previous row, initialized empty
	for zi, c := range fstRunes {
		ppi, pi, i := resolveRowIndex(zi)
		// c = each rune in the first string
		// zi = zero-indexing over the original first string space
		// ppi, pi, i = the before previous, the previous, and the current
		//     row index in the dynamic programming table space

		table[i][0] = table[pi][0] + substDistFunc(c, 0)

		var pd rune = 0 // rune from the previous column, initialized empty
		for zj, d := range sndRunes {
			ppj, pj, j := resolveColIndex(zj)
			// d = each rune in the second string
			// zj = zero-indexing over the original second string space
			// ppj, pj, j = the before previous, the previous, and the current
			//     column index in the dynamic programming table space

			// Compute the minimum score for the rune insertions, deletions,
			// and substitutions only.
			table[i][j] = math.Min(
				table[pi][pj]+substDistFunc(c, d),
				math.Min(
					table[pi][j]+substDistFunc(c, 0),
					table[i][pj]+substDistFunc(0, d),
				),
			)
			// Now consider the transposition penalty if the last two runes
			// of both strings are indeed transposition.
			if pc == d && pd == c {
				table[i][j] = math.Min(
					table[i][j],
					table[ppi][ppj]+transDistFunc(c, d),
				)
			}
			pd = d
		}
		pc = c
	}
	// Extract the actual distance score
	_, _, i := resolveRowIndex(len(fstRunes) - 1)
	_, _, j := resolveColIndex(len(sndRunes) - 1)
	return table[i][j]
}

/*
A helper function which turns a zero-indexing over the original string
space into a triplet of indices corresponding to the before previous,
the previous, and the current row of the dynamic programming table from
the function OptimalAlignmentDistance, respectively. Returned indices
are in modulus 3.
*/
func resolveRowIndex(index int) (int, int, int) {
	return (index - 1) % 3, index % 3, (index + 1) % 3
}

/*
A helper function which turns a zero-indexing over the original string
space into a triplet of indices corresponding to the before previous,
the previous, and the current column of the dynamic programming table from
the function OptimalAlignmentDistance, respectively.
*/
func resolveColIndex(index int) (int, int, int) {
	return index - 1, index, index + 1
}
