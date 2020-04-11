package preset

import (
	"github.com/abhabongse/fuzzymatch-go/candidate/nametitle"
	"github.com/abhabongse/fuzzymatch-go/editdist"
	editdistThai "github.com/abhabongse/fuzzymatch-go/editdist/thai"
	"github.com/abhabongse/fuzzymatch-go/factory"
	"github.com/abhabongse/fuzzymatch-go/sanitary"
	sanitaryThai "github.com/abhabongse/fuzzymatch-go/sanitary/thai"
)

// DefaultSimilarityScore computes the similarity score between two input strings.
// Two input strings will be directly compared under optimal alignment distance metric
// without any pre-processing, and the resulting distance will be re-normalized to
// a similarity score between 0 and 1 (inclusive).
var DefaultSimilarityScore = factory.MakeSimilarityScoreFunction()

// PlainSimilarityScore computes the similarity score between two input strings
// but each input string will be sanitized before they are compared to each other.
var PlainSimilarityScore = factory.MakeSimilarityScoreFunction(
	factory.StringSanitizer(sanitary.LatinExtendedSanitize),
)

// ThaiOptimalAlignmentDist computes the edit distance between two strings
// under optimal alignment distance metric with penalty function
// customized especially for Thai character sets.
var ThaiOptimalAlignmentDist = editdist.MakeOptimalAlignmentDistFunction(
	editdistThai.SubstPenalty, editdistThai.TransPenalty,
)

// ThaiStringSimilarity computes the normalized string similarity
// based on ThaiOptimalAlignmentDist function.
var ThaiStringSimilarity = editdist.MakeStringSimilarityFunction(ThaiOptimalAlignmentDist)

// ThaiNameSimilarityScore computes the similarity score between two input strings
// with extra functionalities:
// (1) Each input string will be sanitized
//     (removing accent symbols from latin characters, removing repeated Thai tonal marks, etc.)
// (2) English and Thai titles are attempted removal (Mr., Miss, etc.)
//     to generate extra candidates.
// (3) ThaiStringSimilarity is used to compute similarity scores between each pair of candidates.
//     Some leniency is given to more common errors.
var ThaiNameSimilarityScore = factory.MakeSimilarityScoreFunction(
	factory.StringSanitizer(sanitaryThai.Sanitize),
	factory.CandidatesGenerator(nametitle.GenerateNamesWithoutTitles),
	factory.SimilarityComputer(ThaiStringSimilarity),
)
