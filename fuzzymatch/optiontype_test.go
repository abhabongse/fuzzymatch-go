package fuzzymatch

import (
	"testing"
)

func TestPanicLinearCombinedScore(t *testing.T) {
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
		defer func() {
			if r := recover(); (r != nil) != shouldPanic {
				t.Errorf("CombinationWeights() panic status = %v, expected %v", r != nil, shouldPanic)
			}
		}()
		NewSimilarityScoreFunction(LinearCombinedScore(args.optimalAlignmentWeight, args.diceSimilarityWeight))
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Try to run the function and recover from panic
			checkPanic(tt.args, tt.shouldPanic)
		})
	}
}
