// This source file contains data essential to perform latin-based character normalization.
// There are two different strategies:
// - Strategy #1: NFKD -> strip non-letter -> NFKC
// - Strategy #2: Apache Lucene ASCII folding

package runedata

// TODO: other ascii folding data from https://github.com/apache/lucene-solr/blob/master/lucene/analysis/common/src/java/org/apache/lucene/analysis/miscellaneous/ASCIIFoldingFilter.java
