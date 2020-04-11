package preset

import (
	"github.com/abhabongse/fuzzymatch-go/candidate"
	editdistExtra "github.com/abhabongse/fuzzymatch-go/legacy_editdist/thai"
	"github.com/abhabongse/fuzzymatch-go/factory"
	"github.com/abhabongse/fuzzymatch-go/sanitary"
	sanitaryExtra "github.com/abhabongse/fuzzymatch-go/sanitary/extra"
)

/*
PlainSimilarityScore is a function which computes the basic similarity score between two
input strings. Both strings are sanitized before they are compared to each other. The
returned score value is a floating point value between 0 (meaning that two strings are
very distinct) and 1 (meaning that two strings are very similar).

Note that the Optimal Alignment distance score utilizes the standard unit rune distance
metrics for both the substitution and the transposition penalties. Also note that the
final similarity score is computed from (1) the simplified Optimal Alignment distance
score and (2) the Sørensen–Dice coefficient; both scores are combined at the ratio 1:2
respectively.
*/
var PlainSimilarityScore = factory.NewSimilarityScoreFunc(
	factory.StringSanitization(sanitary.LatinExtendedSanitize),
	factory.LinearCombinedScore(1.0, 2.0),
)

/*
ThaiNameSimilarityScore is a string similarity score function customized for names
written in Thai scripts. There are two major differences between this function and
PlainSimilarityScore: (1) this function accounts for possible discrepancy in the
appearance of salutation titles; and (2) this function is powered by non-unit rune
distance metrics for edit operations — some leniency is given to more common errors.
*/
var ThaiNameSimilarityScore = factory.NewSimilarityScoreFunc(
	factory.StringSanitization(sanitaryExtra.ThaiSanitize),
	factory.CandidatesGeneration(candidate.NamesWithoutTitles),
	factory.OptimalAlignmentEditDistance(editdistExtra.ThaiSubstPenalty, editdistExtra.ThaiTransPenalty),
	factory.LinearCombinedScore(2.0, 3.0),
)
