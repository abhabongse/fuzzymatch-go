package editdist

import (
	"math"
)

/*
MakeNormalized converts a typical edit-distance computation function in the string space
into its normalized version. Therefore, the returned score of the new edit-distance
function should always be between 0 (meaning two strings are very similar) and 1
(meaning two strings are very different).

Implementation details: the denominator of the fraction is not the length of the longer
string. The reason for this is that some insertion/deletion errors incur sub-unit penalties.
Without the size-fitting denominator, a malicious user may attack by saturating those
insertions/deletions in order to decrease the total edit distances.
*/
func MakeNormalized(computeDist StringDistFunc) StringDistFunc {
	computeNormDist := func(fst, snd string) float64 {
		dist := computeDist(fst, snd)
		fstCapacity := computeDist(fst, "")
		sndCapacity := computeDist("", snd)
		score := dist / math.Max(fstCapacity, sndCapacity)
		if math.IsNaN(score) {
			score = 0.0 // both strings are empty
		}
		return score
	}
	return computeNormDist
}

/*
SimpleAlignmentNormDist is the normalized version of the SimpleAlignmentDist
scoring function whose outputs are guaranteed to be between 0 (meaning that strings are
very similar) and 1 (meaning that strings very distinct). The original distance score
is normalized against the max length of two strings.
*/
var SimpleAlignmentNormDist = MakeNormalized(SimpleAlignmentDist)
