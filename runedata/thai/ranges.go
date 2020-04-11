// This source file contains additional RangeTable for subset of Thai characters.

package thai

import (
	"golang.org/x/text/unicode/rangetable"
	"unicode"
)

// AscendingVowels is a collection of Thai vowels residing above Thai consonants
var AscendingVowels = &unicode.RangeTable{
	R16: []unicode.Range16{
		{toRange16[CharacterMaiHanAkat], toRange16[CharacterMaiHanAkat], 1},
		{toRange16[CharacterSaraI], toRange16[CharacterSaraUee], 1},
	},
}

// ThaiDescending Vowels is a collection of Thai vowels residing below Thai consonants
var DescendingVowels = &unicode.RangeTable{
	R16: []unicode.Range16{
		{toRange16[CharacterSaraU], toRange16[CharacterSaraUu], 1},
	},
}

// AscendingOthers is a collection of Thai non-vowel characters residing above Thai consonants
var AscendingOthers = &unicode.RangeTable{
	R16: []unicode.Range16{
		{toRange16[CharacterMaitaikhu], toRange16[CharacterFongman], 1},
	},
}

// NonSpacingMarks is a collection of all non-spacing mark (Mn) characters in Thai Unicode block
var NonSpacingMarks = rangetable.Merge(AscendingVowels, DescendingVowels, AscendingOthers)
