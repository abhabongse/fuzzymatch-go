package thai

import (
	"github.com/abhabongse/fuzzymatch-go/runedata/thai"
	"unicode"
)

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
