// Package diacritics is the subpackage of package candidate
// which will attempt to remove diacritical marks from extended latin letters
// based on one of two different strategies.
//
// - Strategy #1: Straight diacritics removal (NFKD -> strip Mn -> NFKC)
// - Strategy #2: Apache Lucene ASCII folding
package diacritics
