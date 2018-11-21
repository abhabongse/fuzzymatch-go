package fuzzymatch

import "testing"

func TestSimilarityScore(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty & empty", args{"",""},0},
		{"empty & something", args{"","sunday"},0},
		{"something & empty", args{"saturday",""},0},
		{"days of week", args{"saturday","sunday"},0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimilarityScore(tt.args.fst, tt.args.snd); got != tt.want {
				t.Errorf("SimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
