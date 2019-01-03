package canonical

import (
	"github.com/abhabongse/fuzzymatch-go/runedata"
	"strings"
	"unicode"
)

/*
ReSpace removes leading and trailing white-spaces, then it reduces all inter-word
white-spaces into a single normal space.
*/
func ReSpace(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

/*
RecombineThaiGrams searches for Thai characters written as a non-canonical bigrams
and turn them into its appropriate canonical form. There are 2 patterns:
(1) nikhahit + sara-aa = sara-am; and (2) sara-e + sara-e = sara-ae.
*/
func RecombineThaiGrams(str string) string {
	// TODO (backlog): implement this function in terms of a string transformer

	for _, pattern := range thaiRecombinationTable {
		str = strings.Replace(str, pattern.oldString, pattern.newString, -1)
	}
	return str
}

/*
thaiRecombinationTable is a list of recombination patterns: each pattern describes
an old substring portion which should be replaced with a new substring.
*/
var thaiRecombinationTable = []struct{ oldString, newString string }{
	{
		// Replacing nikhahit + sara-aa = sara-am
		string([]rune{runedata.ThaiCharacterNikhahit, runedata.ThaiCharacterSaraAa}),
		string([]rune{runedata.ThaiCharacterSaraAm}),
	},
	{
		// Replacing sara-e + sara-e = sara-ae
		string([]rune{runedata.ThaiCharacterSaraE, runedata.ThaiCharacterSaraE}),
		string([]rune{runedata.ThaiCharacterSaraAe}),
	},
}

/*
RemoveThaiRepeatedAccidents attempts to remove multiple Thai ascending and descending
characters when they are "accidentally" repeated in succession.

For example, ThaiCharacterMaiEk is never to be repeated more than once. So when the
input string contains two or more consecutive ThaiCharacterMaiEk, then all of them
are moved except one.
*/
func RemoveThaiRepeatedAccidents(str string) string {
	inputRunes := []rune(str)
	outputRunes := make([]rune, 0, len(inputRunes))
	prevRune := rune(0)

	for _, c := range inputRunes {
		if !(c == prevRune && unicode.In(c, runedata.ThaiNonSpacingMarks)) {
			outputRunes = append(outputRunes, c)
		}
		prevRune = c
	}
	return string(outputRunes)
}
