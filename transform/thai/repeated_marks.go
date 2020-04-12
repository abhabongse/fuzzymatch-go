package thai

import (
	runedataThai "github.com/abhabongse/fuzzymatch-go/runedata/thai"
	"golang.org/x/text/transform"
	"unicode"
	"unicode/utf8"
)

// RemoveRepeatedMarksTransformer is a Unicode stream transformer object
// which removes repeated Thai ascending and descending marks except the first.
var RemoveRepeatedMarksTransformer transform.SpanningTransformer = &removeRepeatedMarksSpanningTransformer{}

type removeRepeatedMarksSpanningTransformer struct {
	prevRune rune
}

func (t *removeRepeatedMarksSpanningTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
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
		if r == t.prevRune && unicode.In(r, runedataThai.NonSpacingMarks) {
			// Skip this tonal character
			nSrc += size
			continue
		}
		t.prevRune = r
		// Write character to destination if capacity allows
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
	return
}

func (t *removeRepeatedMarksSpanningTransformer) Reset() {
	t.prevRune = rune(0)
}

func (t *removeRepeatedMarksSpanningTransformer) Span(src []byte, atEOF bool) (n int, err error) {
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
		if unicode.In(r, runedataThai.NonSpacingMarks) {
			err = transform.ErrEndOfSpan
			break
		}
		n += size
	}
	return
}
