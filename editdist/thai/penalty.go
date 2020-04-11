package thai

import (
	runedataThai "github.com/abhabongse/fuzzymatch-go/runedata/thai"
)

/*
runePair is a struct type for a pair of runes. This type is introduced solely
so that a pair of runes can be used as keys to a map structure.
*/
type runePair = struct{ c, d rune }

// thaiSubstPenaltyTable maps a pair of rune characters to their substitution errors.
var thaiSubstPenaltyTable = map[runePair]float64{
	// The following 3 rules describes pairs of characters Thai people use interchangeably
	runePair{runedataThai.CharacterDoChada, runedataThai.CharacterToPatak}:              0.9,
	runePair{runedataThai.CharacterSoSala, runedataThai.CharacterSoRusi}:                0.9,
	runePair{runedataThai.CharacterSoSala, runedataThai.CharacterSoSua}:                 0.9,
	runePair{runedataThai.CharacterSaraAiMaimalai, runedataThai.CharacterSaraAiMaimuan}: 0.8,

	// The following 5 rules given leniency to users who miss some ascending characters.
	runePair{runedataThai.CharacterMaiEk, 0}:       0.6,
	runePair{runedataThai.CharacterMaiTho, 0}:      0.6,
	runePair{runedataThai.CharacterMaiTri, 0}:      0.6,
	runePair{runedataThai.CharacterMaiChattawa, 0}: 0.6,
	runePair{runedataThai.CharacterThanthakhat, 0}: 0.6,

	// The following 3 rules are common tonal confusions even among Thai people
	runePair{runedataThai.CharacterMaiEk, runedataThai.CharacterMaiTho}:     0.6,
	runePair{runedataThai.CharacterMaiTho, runedataThai.CharacterMaiTri}:    0.6,
	runePair{runedataThai.CharacterMaiTri, runedataThai.CharacterMaitaikhu}: 0.6,
}

// thaiTransPenaltyTable maps a pair of rune characters to their transposition error penalties.
var thaiTransPenaltyTable = map[runePair]float64{
	// The following 4 rules are transpositions between SaraAm and tonal marks
	runePair{runedataThai.CharacterSaraAm, runedataThai.CharacterMaiEk}:       0.3,
	runePair{runedataThai.CharacterSaraAm, runedataThai.CharacterMaiTho}:      0.3,
	runePair{runedataThai.CharacterSaraAm, runedataThai.CharacterMaiTri}:      0.3,
	runePair{runedataThai.CharacterSaraAm, runedataThai.CharacterMaiChattawa}: 0.3,
}

/*
SubstPenalty is a function returning substitution penalties with values
between 0 and 1, specialized for Thai rune characters. If erroneous substitutions
between rune characters c and d are more likely, then the penalty will be smaller.
*/
func SubstPenalty(c, d rune) float64 {
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
TransPenalty is a function returning transposition penalties with values
between 0 and 2, specialized for Thai rune characters. If erroneous transpositions
between rune characters c and d are more likely, then the penalty will be smaller.
*/
func TransPenalty(c, d rune) float64 {
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
