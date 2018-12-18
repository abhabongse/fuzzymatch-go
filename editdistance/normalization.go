package editdistance

import "math"

/*
MakeNormalized converts a typical edit-distance computation function in the
string space into a normalized version. Thus, the returned score of the
new edit-distance function should always be between 0 (meaning two strings
are very similar) and 1 (meaning two strings are very different).
*/
func MakeNormalized(computeDistance StringDistanceFunction) StringDistanceFunction {
	computeNormalizedDistance := func(fst, snd string) float64 {
		dist := computeDistance(fst, snd)
		fstLength := computeDistance(fst, "")
		sndLength := computeDistance("", snd)
		score := dist / math.Max(fstLength, sndLength)
		if math.IsNaN(score) {
			score = 0.0 // both strings are empty
		}
		return score
	}
	return computeNormalizedDistance
}

/*
NormalizedSimpleAlignmentDistance is the normalized version of the SimpleAlignmentDistance
scoring function whose outputs are guaranteed to be between 0 (meaning that strings are
very similar) and 1 (meaning that strings very distinct). The original distance score
is normalized against the max length of two strings.
*/
var NormalizedSimpleAlignmentDistance = MakeNormalized(SimpleAlignmentDistance)

/*
NormalizedFuzzyAlignmentDistance
*/
