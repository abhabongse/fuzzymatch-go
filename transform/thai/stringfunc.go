package thai

import (
	"github.com/abhabongse/fuzzymatch-go/runedata/thai"
	"strings"
	"unicode"
)

// RecombineThaiGrams searches for Thai characters
// written as a non-canonical bigrams
// and turn them into its appropriate canonical form.
// There are 2 patterns:
// (1) nikhahit + sara-aa = sara-am; and (2) sara-e + sara-e = sara-ae.
//
// TODO: implement this function in terms of a string transformer
func RecombineThaiGrams(str string) string {
	for _, pattern := range thaiRecombinationTable {
		str = strings.Replace(str, pattern.oldString, pattern.newString, -1)
	}
	return str
}

// thaiRecombinationTable is a list of recombination patterns:
// each pattern describes an old substring portion
// which should be replaced with a new substring.
var thaiRecombinationTable = []struct{ oldString, newString string }{
	{
		// Replacing nikhahit + sara-aa = sara-am
		string([]rune{thai.CharacterNikhahit, thai.CharacterSaraAa}),
		string([]rune{thai.CharacterSaraAm}),
	},
	{
		// Replacing sara-e + sara-e = sara-ae
		string([]rune{thai.CharacterSaraE, thai.CharacterSaraE}),
		string([]rune{thai.CharacterSaraAe}),
	},
}

// RemoveThaiRepeatedAccidents attempts to remove multiple
// Thai ascending and descending characters
// when they are "accidentally" repeated in succession.
//
// For example, ThaiCharacterMaiEk is never to be repeated more than once.
// So when the input string contains two or more consecutive ThaiCharacterMaiEk,
// then all of them are moved except the first one.
//
// TODO: implement this function in terms of a string transformer
func RemoveThaiRepeatedAccidents(str string) string {
	inputRunes := []rune(str)
	outputRunes := make([]rune, 0, len(inputRunes))
	prevRune := rune(0)

	for _, c := range inputRunes {
		if !(c == prevRune && unicode.In(c, thai.NonSpacingMarks)) {
			outputRunes = append(outputRunes, c)
		}
		prevRune = c
	}
	return string(outputRunes)
}
