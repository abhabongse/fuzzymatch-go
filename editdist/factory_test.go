package editdist

import (
	"math"
	"testing"
)

var SimpleLevenshteinDist = MakeLevenshteinDistFunction(UnitPenalty)
var SimpleOptimalAlignmentDist = MakeOptimalAlignmentDistFunction(UnitPenalty, UnitPenalty)
var SimpleLevenshteinStringSimilarity = MakeStringSimilarityFunction(SimpleLevenshteinDist)
var SimpleOptimalAlignmentStringSimilarity = MakeStringSimilarityFunction(SimpleOptimalAlignmentDist)

func TestSimpleLevenshteinStringSimilarity(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"weekends", args{"saturday", "sunday"}, 1 - 0.375},
		{"greetings", args{"hello", "hola"}, 1 - 0.6},
		{"empty", args{"", "hi"}, 1 - 1},
		{"transpose thursday", args{"thrust", "thursday"}, 1 - 0.625},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleLevenshteinStringSimilarity(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("SimpleAlignmentNormDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleOptimalAlignmentStringSimilarity(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"weekends", args{"saturday", "sunday"}, 1 - 0.375},
		{"greetings", args{"hello", "hola"}, 1 - 0.6},
		{"empty", args{"", "hi"}, 1 - 1},
		{"transpose thursday", args{"thrust", "thursday"}, 1 - 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleOptimalAlignmentStringSimilarity(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("SimpleAlignmentNormDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
