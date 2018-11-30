package bigrams

import (
	"math"
	"testing"
)

func TestDiceSimilarityCoefficient(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty", args{"", ""}, 1},
		{"partially empty", args{"", "hi"}, 0},
		{"weekends", args{"saturday", "sunday"}, 0.5},
		{"circular", args{"abaca", "acaba"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiceSimilarityCoefficient(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("DiceSimilarityCoefficient() = %v, want %v", got, tt.want)
			}
		})
	}
}
