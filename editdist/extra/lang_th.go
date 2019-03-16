package extra

import (
	"github.com/abhabongse/fuzzymatch-go/editdist"
	"github.com/abhabongse/fuzzymatch-go/runedata"
)

/*
runePair is a struct type for a pair of runes. This type is introduced solely
so that a pair of runes can be used as keys to a map structure.
*/
type runePair = struct{ c, d rune }

// thaiSubstPenaltyTable maps a pair of rune characters to their substitution errors.
var thaiSubstPenaltyTable = map[runePair]float64{
	// The following 3 rules describes pairs of characters Thai people use interchangeably
	runePair{runedata.ThaiCharacterDoChada, runedata.ThaiCharacterToPatak}:              0.9,
	runePair{runedata.ThaiCharacterSoSala, runedata.ThaiCharacterSoRusi}:                0.9,
	runePair{runedata.ThaiCharacterSoSala, runedata.ThaiCharacterSoSua}:                 0.9,
	runePair{runedata.ThaiCharacterSaraAiMaimalai, runedata.ThaiCharacterSaraAiMaimuan}: 0.8,

	// The following 5 rules given leniency to users who miss some ascending characters.
	runePair{runedata.ThaiCharacterMaiEk, 0}:       0.6,
	runePair{runedata.ThaiCharacterMaiTho, 0}:      0.6,
	runePair{runedata.ThaiCharacterMaiTri, 0}:      0.6,
	runePair{runedata.ThaiCharacterMaiChattawa, 0}: 0.6,
	runePair{runedata.ThaiCharacterThanthakhat, 0}: 0.6,

	// The following 3 rules are common tonal confusions even among Thai people
	runePair{runedata.ThaiCharacterMaiEk, runedata.ThaiCharacterMaiTho}:     0.6,
	runePair{runedata.ThaiCharacterMaiTho, runedata.ThaiCharacterMaiTri}:    0.6,
	runePair{runedata.ThaiCharacterMaiTri, runedata.ThaiCharacterMaitaikhu}: 0.6,
}

// thaiTransPenaltyTable maps a pair of rune characters to their transposition error penalties.
var thaiTransPenaltyTable = map[runePair]float64{
	// The following 4 rules are transpositions between SaraAm and tonal marks
	runePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiEk}:       0.3,
	runePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiTho}:      0.3,
	runePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiTri}:      0.3,
	runePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiChattawa}: 0.3,
}

/*
ThaiSubstPenalty is a function returning substitution penalties with values
between 0 and 1, specialized for Thai rune characters. If erroneous substitutions
between rune characters c and d are more likely, then the penalty will be smaller.
*/
func ThaiSubstPenalty(c, d rune) float64 {
	if c == d {
		return 0 // Exact match
	}
	if value, ok := thaiSubstPenaltyTable[runePair{c, d}]; ok {
		return value
	}
	if value, ok := thaiSubstPenaltyTable[runePair{d, c}]; ok {
		return value
	}
	return 1 // Full substitution penalty
}

/*
ThaiTransPenalty is a function returning transposition penalties with values
between 0 and 2, specialized for Thai rune characters. If erroneous transpositions
between rune characters c and d are more likely, then the penalty will be smaller.
*/
func ThaiTransPenalty(c, d rune) float64 {
	if c == d {
		return 0 // Exact match
	}
	if value, ok := thaiTransPenaltyTable[runePair{c, d}]; ok {
		return value
	}
	if value, ok := thaiTransPenaltyTable[runePair{d, c}]; ok {
		return value
	}
	return 1 // Full transposition penalty
}

/*
ThaiOptimalAlignmentDist is another version of the OptimalAlignmentDist function
customized especially for Thai scripts. Some of the edit operations (insertions, deletions,
substitutions, and adjacent character transpositions) may incur penalties smaller than 1.
*/
var ThaiOptimalAlignmentDist = editdist.MakeOptimalAlignmentDistFunc(ThaiSubstPenalty, ThaiTransPenalty)

/*
ThaiOptimalAlignmentNormDist is the normalized version of the ThaiOptimalAlignmentDist
whose outputs are guaranteed to be between 0 (meaning that strings are very similar) and 1 (meaning
that strings very distinct). The original distance score is normalized against the sum of the
insertion/deletion penalties of one of two strings, whichever is larger.
*/
var ThaiOptimalAlignmentNormDist = editdist.MakeNormalized(ThaiOptimalAlignmentDist)
