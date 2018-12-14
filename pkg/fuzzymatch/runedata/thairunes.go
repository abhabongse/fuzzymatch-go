package runedata

import (
	"golang.org/x/text/unicode/rangetable"
	"unicode"
)

var ThaiCharacterKoKai = '\u0e01'
var ThaiCharacterKhoKhai = '\u0e02'
var ThaiCharacterKhoKhuat = '\u0e03'
var ThaiCharacterKhoKhwai = '\u0e04'
var ThaiCharacterKhoKhon = '\u0e05'
var ThaiCharacterKhoRakhang = '\u0e06'
var ThaiCharacterNgoNgu = '\u0e07'
var ThaiCharacterChoChan = '\u0e08'
var ThaiCharacterChoChing = '\u0e09'
var ThaiCharacterChoChang = '\u0e0a'
var ThaiCharacterSoSo = '\u0e0b'
var ThaiCharacterChoChoe = '\u0e0c'
var ThaiCharacterYoYing = '\u0e0d'
var ThaiCharacterDoChada = '\u0e0e'
var ThaiCharacterToPatak = '\u0e0f'
var ThaiCharacterThoThan = '\u0e10'
var ThaiCharacterThoNangmontho = '\u0e11'
var ThaiCharacterThoPhuthao = '\u0e12'
var ThaiCharacterNoNen = '\u0e13'
var ThaiCharacterDoDek = '\u0e14'
var ThaiCharacterToTao = '\u0e15'
var ThaiCharacterThoThung = '\u0e16'
var ThaiCharacterThoThahan = '\u0e17'
var ThaiCharacterThoThong = '\u0e18'
var ThaiCharacterNoNu = '\u0e19'
var ThaiCharacterBoBaimai = '\u0e1a'
var ThaiCharacterPoPla = '\u0e1b'
var ThaiCharacterPhoPhung = '\u0e1c'
var ThaiCharacterFoFa = '\u0e1d'
var ThaiCharacterPhoPhan = '\u0e1e'
var ThaiCharacterFoFan = '\u0e1f'
var ThaiCharacterPhoSamphao = '\u0e20'
var ThaiCharacterMoMa = '\u0e21'
var ThaiCharacterYoYak = '\u0e22'
var ThaiCharacterRoRua = '\u0e23'
var ThaiCharacterRu = '\u0e24'
var ThaiCharacterLoLing = '\u0e25'
var ThaiCharacterLu = '\u0e26'
var ThaiCharacterWoWaen = '\u0e27'
var ThaiCharacterSoSala = '\u0e28'
var ThaiCharacterSoRusi = '\u0e29'
var ThaiCharacterSoSua = '\u0e2a'
var ThaiCharacterHoHip = '\u0e2b'
var ThaiCharacterLoChula = '\u0e2c'
var ThaiCharacterOAng = '\u0e2d'
var ThaiCharacterHoNokhuk = '\u0e2e'
var ThaiCharacterPaiyannoi = '\u0e2f'
var ThaiCharacterSaraA = '\u0e30'
var ThaiCharacterMaiHanAkat = '\u0e31'
var ThaiCharacterSaraAa = '\u0e32'
var ThaiCharacterSaraAm = '\u0e33'
var ThaiCharacterSaraI = '\u0e34'
var ThaiCharacterSaraIi = '\u0e35'
var ThaiCharacterSaraUe = '\u0e36'
var ThaiCharacterSaraUee = '\u0e37'
var ThaiCharacterSaraU = '\u0e38'
var ThaiCharacterSaraUu = '\u0e39'
var ThaiCharacterPhinthu = '\u0e3a'
var ThaiCurrencySymbolBaht = '\u0e3f'
var ThaiCharacterSaraE = '\u0e40'
var ThaiCharacterSaraAe = '\u0e41'
var ThaiCharacterSaraO = '\u0e42'
var ThaiCharacterSaraAiMaimuan = '\u0e43'
var ThaiCharacterSaraAiMaimalai = '\u0e44'
var ThaiCharacterLakkhangyao = '\u0e45'
var ThaiCharacterMaiyamok = '\u0e46'
var ThaiCharacterMaitaikhu = '\u0e47'
var ThaiCharacterMaiEk = '\u0e48'
var ThaiCharacterMaiTho = '\u0e49'
var ThaiCharacterMaiTri = '\u0e4a'
var ThaiCharacterMaiChattawa = '\u0e4b'
var ThaiCharacterThanthakhat = '\u0e4c'
var ThaiCharacterNikhahit = '\u0e4d'
var ThaiCharacterYamakkan = '\u0e4e'
var ThaiCharacterFongman = '\u0e4f'
var ThaiDigitZero = '\u0e50'
var ThaiDigitOne = '\u0e51'
var ThaiDigitTwo = '\u0e52'
var ThaiDigitThree = '\u0e53'
var ThaiDigitFour = '\u0e54'
var ThaiDigitFive = '\u0e55'
var ThaiDigitSix = '\u0e56'
var ThaiDigitSeven = '\u0e57'
var ThaiDigitEight = '\u0e58'
var ThaiDigitNine = '\u0e59'
var ThaiCharacterAngkhankhu = '\u0e5a'
var ThaiCharacterKhomut = '\u0e5b'

/*
ThaiRunes remaps each character in Thai Unicode block from the rune type back
to the unsigned 16-bit integer type (used in the RangeTable definitions).
*/
var ThaiRunes = map[rune]uint16{
	ThaiCharacterKoKai:          0x0e01,
	ThaiCharacterKhoKhai:        0x0e02,
	ThaiCharacterKhoKhuat:       0x0e03,
	ThaiCharacterKhoKhwai:       0x0e04,
	ThaiCharacterKhoKhon:        0x0e05,
	ThaiCharacterKhoRakhang:     0x0e06,
	ThaiCharacterNgoNgu:         0x0e07,
	ThaiCharacterChoChan:        0x0e08,
	ThaiCharacterChoChing:       0x0e09,
	ThaiCharacterChoChang:       0x0e0a,
	ThaiCharacterSoSo:           0x0e0b,
	ThaiCharacterChoChoe:        0x0e0c,
	ThaiCharacterYoYing:         0x0e0d,
	ThaiCharacterDoChada:        0x0e0e,
	ThaiCharacterToPatak:        0x0e0f,
	ThaiCharacterThoThan:        0x0e10,
	ThaiCharacterThoNangmontho:  0x0e11,
	ThaiCharacterThoPhuthao:     0x0e12,
	ThaiCharacterNoNen:          0x0e13,
	ThaiCharacterDoDek:          0x0e14,
	ThaiCharacterToTao:          0x0e15,
	ThaiCharacterThoThung:       0x0e16,
	ThaiCharacterThoThahan:      0x0e17,
	ThaiCharacterThoThong:       0x0e18,
	ThaiCharacterNoNu:           0x0e19,
	ThaiCharacterBoBaimai:       0x0e1a,
	ThaiCharacterPoPla:          0x0e1b,
	ThaiCharacterPhoPhung:       0x0e1c,
	ThaiCharacterFoFa:           0x0e1d,
	ThaiCharacterPhoPhan:        0x0e1e,
	ThaiCharacterFoFan:          0x0e1f,
	ThaiCharacterPhoSamphao:     0x0e20,
	ThaiCharacterMoMa:           0x0e21,
	ThaiCharacterYoYak:          0x0e22,
	ThaiCharacterRoRua:          0x0e23,
	ThaiCharacterRu:             0x0e24,
	ThaiCharacterLoLing:         0x0e25,
	ThaiCharacterLu:             0x0e26,
	ThaiCharacterWoWaen:         0x0e27,
	ThaiCharacterSoSala:         0x0e28,
	ThaiCharacterSoRusi:         0x0e29,
	ThaiCharacterSoSua:          0x0e2a,
	ThaiCharacterHoHip:          0x0e2b,
	ThaiCharacterLoChula:        0x0e2c,
	ThaiCharacterOAng:           0x0e2d,
	ThaiCharacterHoNokhuk:       0x0e2e,
	ThaiCharacterPaiyannoi:      0x0e2f,
	ThaiCharacterSaraA:          0x0e30,
	ThaiCharacterMaiHanAkat:     0x0e31,
	ThaiCharacterSaraAa:         0x0e32,
	ThaiCharacterSaraAm:         0x0e33,
	ThaiCharacterSaraI:          0x0e34,
	ThaiCharacterSaraIi:         0x0e35,
	ThaiCharacterSaraUe:         0x0e36,
	ThaiCharacterSaraUee:        0x0e37,
	ThaiCharacterSaraU:          0x0e38,
	ThaiCharacterSaraUu:         0x0e39,
	ThaiCharacterPhinthu:        0x0e3a,
	ThaiCurrencySymbolBaht:      0x0e3f,
	ThaiCharacterSaraE:          0x0e40,
	ThaiCharacterSaraAe:         0x0e41,
	ThaiCharacterSaraO:          0x0e42,
	ThaiCharacterSaraAiMaimuan:  0x0e43,
	ThaiCharacterSaraAiMaimalai: 0x0e44,
	ThaiCharacterLakkhangyao:    0x0e45,
	ThaiCharacterMaiyamok:       0x0e46,
	ThaiCharacterMaitaikhu:      0x0e47,
	ThaiCharacterMaiEk:          0x0e48,
	ThaiCharacterMaiTho:         0x0e49,
	ThaiCharacterMaiTri:         0x0e4a,
	ThaiCharacterMaiChattawa:    0x0e4b,
	ThaiCharacterThanthakhat:    0x0e4c,
	ThaiCharacterNikhahit:       0x0e4d,
	ThaiCharacterYamakkan:       0x0e4e,
	ThaiCharacterFongman:        0x0e4f,
	ThaiDigitZero:               0x0e50,
	ThaiDigitOne:                0x0e51,
	ThaiDigitTwo:                0x0e52,
	ThaiDigitThree:              0x0e53,
	ThaiDigitFour:               0x0e54,
	ThaiDigitFive:               0x0e55,
	ThaiDigitSix:                0x0e56,
	ThaiDigitSeven:              0x0e57,
	ThaiDigitEight:              0x0e58,
	ThaiDigitNine:               0x0e59,
	ThaiCharacterAngkhankhu:     0x0e5a,
	ThaiCharacterKhomut:         0x0e5b,
}

// TODO: the following three range tables should be used for repeated chars removal

// ThaiAscendingVowels is a collection of Thai vowels residing above Thai consonants
var ThaiAscendingVowels = &unicode.RangeTable{
	R16: []unicode.Range16{
		{ThaiRunes[ThaiCharacterMaiHanAkat], ThaiRunes[ThaiCharacterMaiHanAkat], 1},
		{ThaiRunes[ThaiCharacterSaraI], ThaiRunes[ThaiCharacterSaraUee], 1},
	},
}

// ThaiDescending Vowels is a collection of Thai vowels residing below Thai consonants
var ThaiDescendingVowels = &unicode.RangeTable{
	R16: []unicode.Range16{
		{ThaiRunes[ThaiCharacterSaraU], ThaiRunes[ThaiCharacterSaraUu], 1},
	},
}

// ThaiAscendingOthers is a collection of Thai non-vowel characters residing above Thai consonants
var ThaiAscendingOthers = &unicode.RangeTable{
	R16: []unicode.Range16{
		{ThaiRunes[ThaiCharacterMaitaikhu], ThaiRunes[ThaiCharacterFongman], 1},
	},
}

// ThaiNonSpacingMarks is a collection of all non-spacing mark (Mn) characters in Thai Unicode block
var ThaiNonSpacingMarks = rangetable.Merge(ThaiAscendingVowels, ThaiDescendingVowels, ThaiAscendingOthers)
