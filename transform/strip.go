package transform

import (
	"github.com/abhabongse/fuzzymatch-go/runedata"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"unicode"
	"unicode/utf8"
)

// StripNonPrintingTransform is a Unicode stream transformer object
// which removes all occurrences of non-printing and non-spacing rune characters
// from a string.
var StripNonPrintTransformer = runes.Remove(runes.NotIn(runedata.PrintsAndWhiteSpaces))

// RespaceTransformer is a Unicode stream transformer object
// which removes all leading and trailing white-spaces,
// then it reduces all inter-word white-spaces into a single normal space.
var RespaceTransformer transform.SpanningTransformer = respaceSpanningTransformer{}

type respaceSpanningTransformer struct {
	afterFirstNonSpace bool
}

func (t respaceSpanningTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for r, size := rune(0), 0; nSrc < len(src); {
		// Attempt to decode the current rune
		if r = rune(src[nSrc]); r < utf8.RuneSelf {
			size = 1
		} else if r, size = utf8.DecodeRune(src[nSrc:]); size == 1 {
			// Invalid rune
			if !atEOF && !utf8.FullRune(src[nSrc:]) {
				err = transform.ErrShortSrc
				break
			}
		}
		// If the current rune is a space, then lookahead and check the next one
		if unicode.IsSpace(r) {
			if nSrc+size == len(src) {
				if atEOF {
					// Apparently this is the final space character,
					// and we could just skip and be done with this!
					nSrc += size
					continue
				} else {
					// Need to see the next rune first, sorry!
					err = transform.ErrShortSrc
					break
				}
			}
			// Attempt to decode the subsequent rune
			r2, size2 := rune(0), 0
			if r2 = rune(src[nSrc+size]); r < utf8.RuneSelf {
				size2 = 1
			} else if r2, size2 = utf8.DecodeRune(src[nSrc+size:]); size2 == 1 {
				// Invalid rune
				if !atEOF && !utf8.FullRune(src[nSrc+size:]) {
					err = transform.ErrShortSrc
					break
				}
			}
			// If the next rune is also a white space
			// or if we are still absolutely in the leading space territory,
			// then we should skip this space character.
			if unicode.IsSpace(r2) || !t.afterFirstNonSpace {
				nSrc += size
				continue
			}
			// Now we are sure that a single normal space
			// should be written to the destination
			// (of course, if the capacity allows)
			if nDst+1 > len(dst) {
				err = transform.ErrShortDst
				break
			}
			dst[nDst] = ' '
			nDst++
			nSrc += size
		} else {
			// Mark that it is no longer leading space territory
			t.afterFirstNonSpace = true
			// Copy over this non-space character if capacity allows
			if nDst+size > len(dst) {
				err = transform.ErrShortDst
				break
			}
			for i := 0; i < size; i++ {
				dst[nDst] = src[nSrc]
				nDst++
				nSrc++
			}
		}
	}
	return
}

func (t respaceSpanningTransformer) Reset() {
	t.afterFirstNonSpace = false
}

func (t respaceSpanningTransformer) Span(src []byte, atEOF bool) (n int, err error) {
	for r, size := rune(0), 0; n < len(src); {
		// Attempt to decode the current rune
		if r = rune(src[n]); r < utf8.RuneSelf {
			size = 1
		} else if r, size = utf8.DecodeRune(src[n:]); size == 1 {
			// Invalid rune
			if !atEOF && !utf8.FullRune(src[n:]) {
				err = transform.ErrShortSrc
			} else {
				err = transform.ErrEndOfSpan
			}
			break
		}
		if unicode.IsSpace(r) {
			err = transform.ErrEndOfSpan
			break
		}
		n += size
	}
	return
}
