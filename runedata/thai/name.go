// This source file contains definitions about each character in Thai Unicode block.

package thai

var CharacterKoKai = '\u0e01'
var CharacterKhoKhai = '\u0e02'
var CharacterKhoKhuat = '\u0e03'
var CharacterKhoKhwai = '\u0e04'
var CharacterKhoKhon = '\u0e05'
var CharacterKhoRakhang = '\u0e06'
var CharacterNgoNgu = '\u0e07'
var CharacterChoChan = '\u0e08'
var CharacterChoChing = '\u0e09'
var CharacterChoChang = '\u0e0a'
var CharacterSoSo = '\u0e0b'
var CharacterChoChoe = '\u0e0c'
var CharacterYoYing = '\u0e0d'
var CharacterDoChada = '\u0e0e'
var CharacterToPatak = '\u0e0f'
var CharacterThoThan = '\u0e10'
var CharacterThoNangmontho = '\u0e11'
var CharacterThoPhuthao = '\u0e12'
var CharacterNoNen = '\u0e13'
var CharacterDoDek = '\u0e14'
var CharacterToTao = '\u0e15'
var CharacterThoThung = '\u0e16'
var CharacterThoThahan = '\u0e17'
var CharacterThoThong = '\u0e18'
var CharacterNoNu = '\u0e19'
var CharacterBoBaimai = '\u0e1a'
var CharacterPoPla = '\u0e1b'
var CharacterPhoPhung = '\u0e1c'
var CharacterFoFa = '\u0e1d'
var CharacterPhoPhan = '\u0e1e'
var CharacterFoFan = '\u0e1f'
var CharacterPhoSamphao = '\u0e20'
var CharacterMoMa = '\u0e21'
var CharacterYoYak = '\u0e22'
var CharacterRoRua = '\u0e23'
var CharacterRu = '\u0e24'
var CharacterLoLing = '\u0e25'
var CharacterLu = '\u0e26'
var CharacterWoWaen = '\u0e27'
var CharacterSoSala = '\u0e28'
var CharacterSoRusi = '\u0e29'
var CharacterSoSua = '\u0e2a'
var CharacterHoHip = '\u0e2b'
var CharacterLoChula = '\u0e2c'
var CharacterOAng = '\u0e2d'
var CharacterHoNokhuk = '\u0e2e'
var CharacterPaiyannoi = '\u0e2f'
var CharacterSaraA = '\u0e30'
var CharacterMaiHanAkat = '\u0e31'
var CharacterSaraAa = '\u0e32'
var CharacterSaraAm = '\u0e33'
var CharacterSaraI = '\u0e34'
var CharacterSaraIi = '\u0e35'
var CharacterSaraUe = '\u0e36'
var CharacterSaraUee = '\u0e37'
var CharacterSaraU = '\u0e38'
var CharacterSaraUu = '\u0e39'
var CharacterPhinthu = '\u0e3a'
var CurrencySymbolBaht = '\u0e3f'
var CharacterSaraE = '\u0e40'
var CharacterSaraAe = '\u0e41'
var CharacterSaraO = '\u0e42'
var CharacterSaraAiMaimuan = '\u0e43'
var CharacterSaraAiMaimalai = '\u0e44'
var CharacterLakkhangyao = '\u0e45'
var CharacterMaiyamok = '\u0e46'
var CharacterMaitaikhu = '\u0e47'
var CharacterMaiEk = '\u0e48'
var CharacterMaiTho = '\u0e49'
var CharacterMaiTri = '\u0e4a'
var CharacterMaiChattawa = '\u0e4b'
var CharacterThanthakhat = '\u0e4c'
var CharacterNikhahit = '\u0e4d'
var CharacterYamakkan = '\u0e4e'
var CharacterFongman = '\u0e4f'
var DigitZero = '\u0e50'
var DigitOne = '\u0e51'
var DigitTwo = '\u0e52'
var DigitThree = '\u0e53'
var DigitFour = '\u0e54'
var DigitFive = '\u0e55'
var DigitSix = '\u0e56'
var DigitSeven = '\u0e57'
var DigitEight = '\u0e58'
var DigitNine = '\u0e59'
var CharacterAngkhankhu = '\u0e5a'
var CharacterKhomut = '\u0e5b'

// toRange16 remaps each character in Thai Unicode block from the rune type back
// to the unsigned 16-bit integer type (used in the RangeTable definitions).
var toRange16 = map[rune]uint16{
	CharacterKoKai:          0x0e01,
	CharacterKhoKhai:        0x0e02,
	CharacterKhoKhuat:       0x0e03,
	CharacterKhoKhwai:       0x0e04,
	CharacterKhoKhon:        0x0e05,
	CharacterKhoRakhang:     0x0e06,
	CharacterNgoNgu:         0x0e07,
	CharacterChoChan:        0x0e08,
	CharacterChoChing:       0x0e09,
	CharacterChoChang:       0x0e0a,
	CharacterSoSo:           0x0e0b,
	CharacterChoChoe:        0x0e0c,
	CharacterYoYing:         0x0e0d,
	CharacterDoChada:        0x0e0e,
	CharacterToPatak:        0x0e0f,
	CharacterThoThan:        0x0e10,
	CharacterThoNangmontho:  0x0e11,
	CharacterThoPhuthao:     0x0e12,
	CharacterNoNen:          0x0e13,
	CharacterDoDek:          0x0e14,
	CharacterToTao:          0x0e15,
	CharacterThoThung:       0x0e16,
	CharacterThoThahan:      0x0e17,
	CharacterThoThong:       0x0e18,
	CharacterNoNu:           0x0e19,
	CharacterBoBaimai:       0x0e1a,
	CharacterPoPla:          0x0e1b,
	CharacterPhoPhung:       0x0e1c,
	CharacterFoFa:           0x0e1d,
	CharacterPhoPhan:        0x0e1e,
	CharacterFoFan:          0x0e1f,
	CharacterPhoSamphao:     0x0e20,
	CharacterMoMa:           0x0e21,
	CharacterYoYak:          0x0e22,
	CharacterRoRua:          0x0e23,
	CharacterRu:             0x0e24,
	CharacterLoLing:         0x0e25,
	CharacterLu:             0x0e26,
	CharacterWoWaen:         0x0e27,
	CharacterSoSala:         0x0e28,
	CharacterSoRusi:         0x0e29,
	CharacterSoSua:          0x0e2a,
	CharacterHoHip:          0x0e2b,
	CharacterLoChula:        0x0e2c,
	CharacterOAng:           0x0e2d,
	CharacterHoNokhuk:       0x0e2e,
	CharacterPaiyannoi:      0x0e2f,
	CharacterSaraA:          0x0e30,
	CharacterMaiHanAkat:     0x0e31,
	CharacterSaraAa:         0x0e32,
	CharacterSaraAm:         0x0e33,
	CharacterSaraI:          0x0e34,
	CharacterSaraIi:         0x0e35,
	CharacterSaraUe:         0x0e36,
	CharacterSaraUee:        0x0e37,
	CharacterSaraU:          0x0e38,
	CharacterSaraUu:         0x0e39,
	CharacterPhinthu:        0x0e3a,
	CurrencySymbolBaht:      0x0e3f,
	CharacterSaraE:          0x0e40,
	CharacterSaraAe:         0x0e41,
	CharacterSaraO:          0x0e42,
	CharacterSaraAiMaimuan:  0x0e43,
	CharacterSaraAiMaimalai: 0x0e44,
	CharacterLakkhangyao:    0x0e45,
	CharacterMaiyamok:       0x0e46,
	CharacterMaitaikhu:      0x0e47,
	CharacterMaiEk:          0x0e48,
	CharacterMaiTho:         0x0e49,
	CharacterMaiTri:         0x0e4a,
	CharacterMaiChattawa:    0x0e4b,
	CharacterThanthakhat:    0x0e4c,
	CharacterNikhahit:       0x0e4d,
	CharacterYamakkan:       0x0e4e,
	CharacterFongman:        0x0e4f,
	DigitZero:               0x0e50,
	DigitOne:                0x0e51,
	DigitTwo:                0x0e52,
	DigitThree:              0x0e53,
	DigitFour:               0x0e54,
	DigitFive:               0x0e55,
	DigitSix:                0x0e56,
	DigitSeven:              0x0e57,
	DigitEight:              0x0e58,
	DigitNine:               0x0e59,
	CharacterAngkhankhu:     0x0e5a,
	CharacterKhomut:         0x0e5b,
}
