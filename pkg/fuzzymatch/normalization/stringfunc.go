package normalization

import (
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/runedata"
	"strings"
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
	// TODO: implement this function in terms of a string transformer

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
