package fuzzymatch

import (
	"fmt"
	"testing"

	"github.com/abhabongse/fuzzymatch-go/editdistance"
)

func ExampleSimilarityScore() {
	var SimilarityScore func(string, string) float64
	// Construct a basic string similarity score function (these lines are equivalent)
	SimilarityScore = NewSimilarityScoreFunction()
	SimilarityScore = NewSimilarityScoreFunction(
		StringCanonicalization(DefaultCanonicalizeString),
		CandidateGeneration(DefaultGenerateCandidates),
		RuneDistancePenalties(editdistance.UnitDist, editdistance.UnitDist),
		CombinationWeights(1.0, 0.0),
	)

	// Constructed string similarity score function can be applied to pairs of strings
	score := SimilarityScore("saturday", "sunday")
	fmt.Println(score)
}

func TestCombinationWeights(t *testing.T) {
	type args struct {
		optimalAlignmentWeight float64
		diceSimilarityWeight   float64
	}
	tests := []struct {
		name        string
		args        args
		shouldPanic bool
	}{
		{"default", args{1.0, 0.0}, false},
		{"positive", args{2.0, 3.0}, false},
		{"all zeros", args{0.0, 0.0}, true},
		{"negative", args{3.0, -2.0}, true},
		// TODO: Add test cases.
	}

	checkPanic := func(args args, shouldPanic bool) {
		defer func(){
			if r := recover(); (r != nil) != shouldPanic {
				t.Errorf("CombinationWeights() panic status = %v, expected %v", r != nil, shouldPanic)
			}
		}()
		NewSimilarityScoreFunction(CombinationWeights(args.optimalAlignmentWeight, args.diceSimilarityWeight))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Try to run the function and recover from panic
			checkPanic(tt.args, tt.shouldPanic)
		})
	}
}
