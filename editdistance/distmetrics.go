package editdistance

import (
	"github.com/abhabongse/fuzzymatch-go/runedata"
)

/*
StringDistanceFunction is a type alias for string edit-distance functions.
*/
type StringDistanceFunction = func(string, string) float64

/*
RunePair is a struct type for a pair of runes. This type is introduced solely
so that a pair of runes can be used as keys to a map structure.
*/
type RunePair = struct{ c, d rune }

/*
RuneDistanceMetric is an umbrella type alias representing a symmetric function which
computes a distance score (penalty) between a given pair of runes. There are two
different kinds of distance metrics: substitution penalty and transposition penalty.


Substitution penalty

Substitution penalty refers to the cost of character substitutions (which includes
insertions and deletions as special cases). Specifically, a function of this kind
takes in two runes as input arguments and returns the penalty value between 0
(meaning that two runes are identical) and 1 (meaning that two runes are totally
distinct). The insertion and deletion penalties are represented by assigning byte 0
as one of the input argument to the function.


Transposition penalty

Transposition penalty refers to the cost of adjacent character transpositions.
Specifically, a function of this kind takes in two runes as input argument and
returns the penalty value between 0 (meaning that both runes are interchangeable)
and 2 (meaning that two runes are hardly mistaken with each other).
*/
type RuneDistanceMetric = func(rune, rune) float64

/*
UnitDist is an indicator function which returns 1 if two input characters differ
and returns 0 otherwise. This function can be used for both the substitution and the
transposition penalties.

UnitDist is a symmetric function, meaning that the order of the input arguments do
not matter.

For example, UnitDist('a', 'b') is the penalty for substituting character 'a' with
another character 'b' (which is 1 in this case) whereas UnitDist('c', 0) is the
penalty for an insertion or a deletion of character 'c'.
*/
func UnitDist(c, d rune) float64 {
	if c == d {
		return 0
	}
	return 1
}

/*
PartialSubstitutionErrorDist is a function returning substitution penalties with values
between 0 and 1, specialized for Thai rune characters. If erroneous substitutions
between rune characters c and d are more likely, then the penalty will be smaller.
*/
func PartialSubstitutionErrorDist(c, d rune) float64 {
	if c == d {
		return 0 // Exact match
	}
	if value, ok := substitutionErrorTable[RunePair{c, d}]; ok {
		return value
	}
	if value, ok := substitutionErrorTable[RunePair{d, c}]; ok {
		return value
	}
	return 1 // Full substitution penalty
}

/*
substitutionErrorTable maps a pair of rune characters to their substitution errors.
*/
var substitutionErrorTable = map[RunePair]float64{

	// The following 3 rules describes pairs of characters Thai people use interchangeably
	RunePair{runedata.ThaiCharacterDoChada, runedata.ThaiCharacterToPatak}:              0.9,
	RunePair{runedata.ThaiCharacterSoSala, runedata.ThaiCharacterSoRusi}:                0.9,
	RunePair{runedata.ThaiCharacterSoSala, runedata.ThaiCharacterSoSua}:                 0.9,
	RunePair{runedata.ThaiCharacterSaraAiMaimalai, runedata.ThaiCharacterSaraAiMaimuan}: 0.8,

	// The following 5 rules given leniency to users who miss some ascending characters.
	RunePair{runedata.ThaiCharacterMaiEk, 0}:       0.6,
	RunePair{runedata.ThaiCharacterMaiTho, 0}:      0.6,
	RunePair{runedata.ThaiCharacterMaiTri, 0}:      0.6,
	RunePair{runedata.ThaiCharacterMaiChattawa, 0}: 0.6,
	RunePair{runedata.ThaiCharacterThanthakhat, 0}: 0.6,

	// The following 3 rules are common tonal confusions even among Thai people
	RunePair{runedata.ThaiCharacterMaiEk, runedata.ThaiCharacterMaiTho}:     0.6,
	RunePair{runedata.ThaiCharacterMaiTho, runedata.ThaiCharacterMaiTri}:    0.6,
	RunePair{runedata.ThaiCharacterMaiTri, runedata.ThaiCharacterMaitaikhu}: 0.6,
}

/*
PartialTranspositionErrorDist is a function returning transposition penalties with values
between 0 and 2, specialized for Thai rune characters. If erroneous transpositions
between rune characters c and d are more likely, then the penalty will be smaller.
*/
func PartialTranspositionErrorDist(c, d rune) float64 {
	if c == d {
		return 0 // Exact match
	}
	if value, ok := transpositionErrorTable[RunePair{c, d}]; ok {
		return value
	}
	if value, ok := transpositionErrorTable[RunePair{d, c}]; ok {
		return value
	}
	return 1 // Full transposition penalty
}

/*
transpositionErrorTable maps a pair of rune characters to their transposition error penalties.
*/
var transpositionErrorTable = map[RunePair]float64{
	// The following 4 rules are transpositions between SaraAm and tonal marks
	RunePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiEk}:       0.3,
	RunePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiTho}:      0.3,
	RunePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiTri}:      0.3,
	RunePair{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiChattawa}: 0.3,
}
