/*
Package presets is a collection of pre-combined approximate string matching
algorithms which determines a similarity score between two strings.

In most cases, you would be interested in the function 'PlainSimilarityScore'.
*/
package presets

import (
	"github.com/abhabongse/fuzzymatch-go/candidacy"
	"github.com/abhabongse/fuzzymatch-go/editdistance"
	"github.com/abhabongse/fuzzymatch-go/fuzzymatch"
)

/*
PlainSimilarityScore is a definitive function which computes the similarity score between
two input strings. Both strings are sanitized before they are compared to each other.
The returned score value is a floating point value between 0 (meaning that two strings
are very distinct) and 1 (meaning that two strings are very similar).

The final similarity score is computed from (1) the simplified Optimal Alignment distance
score and (2) the Sørensen–Dice coefficient; both scores are combined at the ratio 1:2
respectively.
*/
var PlainSimilarityScore, _ = fuzzymatch.NewSimilarityScoreFunction()

var ThaiNameSimilarityScore, _ = fuzzymatch.NewSimilarityScoreFunction(
	fuzzymatch.CandidateGeneration(candidacy.PossibleBareNames),
	fuzzymatch.RuneDistancePenalties(editdistance.ThaiSubstitutionErrorDist, editdistance.ThaiTranspositionErrorDist),
	fuzzymatch.CombinationWeights(2.0, 3.0),
)
