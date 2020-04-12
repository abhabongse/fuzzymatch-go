package transform

import (
	"github.com/abhabongse/fuzzymatch-go/runedata"
	"golang.org/x/text/runes"
	"unicode"
)

// ToNormalSpaceTransformer is a Unicode stream transformer object
// which replaces all white space rune characters into a normal space.
var ToNormalSpaceTransformer = runes.Map(func(r rune) rune {
	if unicode.Is(unicode.White_Space, r) {
		return ' '
	}
	return r
})

// ToNormalHyphenTransformer is a Unicode stream transformer object
// which replaces all dash or hyphen characters into a normal hyphen.
var ToNormalHyphenTransformer = runes.Map(func(r rune) rune {
	if unicode.Is(runedata.HyphensAndDashes, r) {
		return '-'
	}
	return r
})
