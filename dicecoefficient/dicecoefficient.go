package dicecoefficient

import "sort"

/*
The Sørensen–Dice Coefficient (sometimes called the Dice Similarity Coefficient; DSC)
is one of the metrics to measure the similarity between two given strings (each provided
as a slice of rune characters). This function computes this value by first converting
each string into a multi-set (i.e. a bag) of bigrams, then it determines how much both
bags overlap. The concrete formula for the Dice Similarity Coefficient is
    2 * n_itx / (n_fst + n_snd)
where
    - n_fst = the number of bigrams in the first string
    - n_snd = the number of bigrams in the second string
    - n_itx = the number of bigrams in the intersection of both bags
*/
func DiceSimilarityCoefficient(fst, snd string) float64 {
	fstRunes := []rune(fst)
	sndRunes := []rune(snd)
	fstBigramSeq := newSortedBigramSequence(fstRunes)
	sndBigramSeq := newSortedBigramSequence(sndRunes)

	numIntersections := 0
	for i, j := 0, 0; i < fstBigramSeq.Len() || j < sndBigramSeq.Len(); {
		if i == fstBigramSeq.Len() {
			j++
		} else if j == sndBigramSeq.Len() {
			i++
		} else if less(fstBigramSeq, i, sndBigramSeq, j) {
			i++
		} else if less(sndBigramSeq, j, fstBigramSeq, i) {
			j++
		} else {
			i++
			j++
			numIntersections++
		}
	}
	return float64(2*numIntersections) / float64(fstBigramSeq.Len()+sndBigramSeq.Len())
}

/*
This data type contains the information about a sequence of all possible bigrams found
in a referenced string (provided as a slice of rune characters) called runesData.
The sequence itself is a permutation of indices (for i = 0 to len(runesData), inclusive)
stored within the attribute indexMap. Particularly, the i-th bigram of this sequence is
a pair of runes located at positions indexMap[i]-1 and indexMap[i].

This sequence of bigrams can be sorted to represent different sequences of bigrams.
It also has several methods to read the actual bigrams at a given index.
*/
type bigramSequence struct {
	runesData []rune
	indexMap  []int
}

/*
The i-th bigram of the sequence of bigrams.
*/
func (bgSeq bigramSequence) bigram(i int) (rune, rune) {
	remappedIndex := bgSeq.indexMap[i]
	var left, right rune
	if remappedIndex == 0 {
		left = ' '
	} else {
		left = bgSeq.runesData[remappedIndex-1]
	}
	if remappedIndex == len(bgSeq.runesData) {
		right = ' '
	} else {
		right = bgSeq.runesData[remappedIndex]
	}
	return left, right
}

/*
The number of bigrams presented in the string.
Part of 'Sorter' interface.
*/
func (bgSeq bigramSequence) Len() int {
	return len(bgSeq.indexMap)
}

/*
Swaps the given positions of two bigrams in the sequence.
Part of 'Sorter' and 'swapper' interfaces.
*/
func (bgSeq bigramSequence) Swap(i, j int) {
	bgSeq.indexMap[i], bgSeq.indexMap[j] = bgSeq.indexMap[j], bgSeq.indexMap[i]
}

/*
Compare two bigrams at given positions to see which one comes first lexicographically.
Part of the 'Sorter' interface.
*/
func (bgSeq bigramSequence) Less(i, j int) bool {
	return less(&bgSeq, i, &bgSeq, j)
}

/*
Compare two bigrams from two sequences (potentially the same sequence) to see which
one comes first lexicographically.
*/
func less(bgSeqA *bigramSequence, indexA int, bgSeqB *bigramSequence, indexB int) bool {
	leftRuneA, rightRuneA := bgSeqA.bigram(indexA)
	leftRuneB, rightRuneB := bgSeqB.bigram(indexB)
	return leftRuneA < leftRuneB || leftRuneA == leftRuneB && rightRuneA < rightRuneB
}

/*
Based on a given string (provided as a slice of rune characters), this function first
generate a new proxy object representing a sequence of all possible bigrams. Then this
sequence of bigrams is sorted lexicographically before it is returned.
*/
func newSortedBigramSequence(runesData []rune) *bigramSequence {
	runesLength := len(runesData)
	indexMap := make([]int, runesLength+1)
	for i := 0; i <= runesLength; i++ {
		indexMap[i] = i
	}
	bgSeq := &bigramSequence{runesData, indexMap}
	sort.Sort(bgSeq)
	return bgSeq
}
