package thai

import (
	"github.com/abhabongse/fuzzymatch-go/runedata/thai"
	"golang.org/x/text/transform"
	"unicode/utf8"
)

// BigramRecombineTransformer is a Unicode stream transformer object
// which removes all leading and trailing white-spaces,
// then it reduces all inter-word white-spaces into a single normal space.
var BigramRecombineTransformer transform.Transformer = bigramRecombineTransformer{}

type bigramRecombineTransformer struct {
	transform.NopResetter
}

func (t bigramRecombineTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
outer:
	for nSrc < len(src) {
		bigram := [2]rune{0, 0}
		totalSize := 0
		for c := 0; c < 2; c++ {
			r, size := rune(0), 0
			// If the accumulation is too much to handle, bail and restart
			if nSrc+totalSize == len(src) {
				// Well, unless we are already at the end of source input
				if atEOF {
					// Channel out all characters and be gone!
					partialBytes := []byte(string(bigram[:c]))
					partialSize := len(partialBytes)
					if nDst+partialSize > len(dst) {
						// Welp! Still not enough capacity in the destination
						err = transform.ErrShortDst
					} else {
						for i := 0; i < partialSize; i++ {
							dst[nDst] = partialBytes[i]
							nDst++
						}
						nSrc += totalSize
					}
				} else {
					// Need to see more bytes, sorry!
					err = transform.ErrShortSrc
				}
				break outer
			}
			if r = rune(src[nSrc+totalSize]); r < utf8.RuneSelf {
				size = 1
			} else if r, size = utf8.DecodeRune(src[nSrc+totalSize:]); size == 1 {
				if !atEOF && !utf8.FullRune(src[nSrc+totalSize:]) {
					err = transform.ErrShortSrc
					break outer
				}
			}
			bigram[c] = r
			totalSize += size
		}
		var targetBytes []byte
		movedSize := 0
		if target, ok := bigramRecombinationTable[bigram]; ok {
			// Copy the entire target into destination
			targetBytes = []byte(target)
			movedSize = totalSize
		} else {
			// Copy only the first rune into destination
			targetBytes = []byte(string(bigram[0]))
			movedSize = len(targetBytes)
		}
		targetSize := len(targetBytes)
		if nDst+targetSize > len(dst) {
			err = transform.ErrShortDst
			break outer
		}
		for i := 0; i < targetSize; i++ {
			dst[nDst] = targetBytes[i]
			nDst++
		}
		nSrc += movedSize
	}
	return
}

var bigramRecombinationTable = map[[2]rune]string{
	[2]rune{thai.CharacterNikhahit, thai.CharacterSaraAa}: string([]rune{thai.CharacterSaraAm}),
	[2]rune{thai.CharacterSaraE, thai.CharacterSaraE}:     string([]rune{thai.CharacterSaraAe}),
}
