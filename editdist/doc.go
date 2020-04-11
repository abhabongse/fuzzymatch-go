// Package editdist provides a set of string comparison functions
// to compute the distance between a pair of strings
// under various distance metrics in string space.
// The higher the distance, the greater the difference between both strings.
//
//
// Levenshtein Distance
//
// In the original Levenshtein distance metric,
// the distance between any two string is measured by the minimum "edit operations"
// required to transform one string into the other.
// These edit operations are limited to (1) an insertion of a character,
// (2) a deletion of a character, and (3) a substitution of a character by another character.
//
// For example, when using UnitPenalty indicator function as the character substitution penalty,
// the Levenshtein distance between "Hello" and "Hola" is 3,
// whereas the distance between "Hi" and "" (empty string) is 2.
//
//
// Damerau–Levenshtein Distance
//
// The Damerau–Levenshtein distance metric is an improvement upon the original Levenshtein distance
// by allowing another kind of edit operation: a transposition of two adjacent characters.
//
// For example, when using UnitPenalty indicator function
// as both the character substitution and character transposition penalty,
// the Damerau–Levenshtein distance between "Thrust" and "Thursday" is 4.
// Without transposition penalty, the distance under the original Levenshtein metric
// would have become 5.
//
// Note: The efficient implementation of Damerau-Levenshtein distance is hard to come by,
// hence we use the weaker version of Damerau-Levenshtein distance called Optimal Alignment,
// which is described next.
//
//
// Optimal Alignment Distance
//
// The optimal alignment distance is the restricted version of the Damerau–Levenshtein distance;
// specifically, each rune character in both the original and the target string
// is subjected to at most one edit operation,
// and only characters that are adjacent in the original input are allowed to be transposed.
//
// Therefore, Damerau–Levenshtein distance between "trout" and "turn" is 3
// (as shown by "trout" → "trut" → "turt" → "turn").
// But the second operation would not be allowed in the optimal alignment metric,
// and thus yield a worse distance of 4 (as shown by "trout" → "tout" → "tut" → "turt" → "turn").
package editdist
