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
RecombineThaiGrams recombines two kinds of bigrams into single characters:
(1) nikhahit + sara-aa = sara-am; and (2) sara-e + sara-e = sara-ae.
*/
func RecombineThaiGrams(str string) string {
	// TODO: implement this function in terms of a string transformer
	// TODO: separate table configuration from function

	preSaraAm := string([]rune{runedata.ThaiCharacterNikhahit, runedata.ThaiCharacterSaraAa})
	postSaraAm := string([]rune{runedata.ThaiCharacterSaraAm})
	str = strings.Replace(str, preSaraAm, postSaraAm, -1)

	preSaraAe := string([]rune{runedata.ThaiCharacterSaraE, runedata.ThaiCharacterSaraE})
	postSaraAe := string([]rune{runedata.ThaiCharacterSaraAe})
	str = strings.Replace(str, preSaraAe, postSaraAe, -1)

	return str
}
