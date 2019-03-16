package editdist

import (
	"math"
	"testing"
)

func TestNormalizedSimpleAlignmentDistance(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"weekends", args{"saturday", "sunday"}, 0.375},
		{"greetings", args{"hello", "hola"}, 0.6},
		{"empty", args{"", "hi"}, 1},
		{"transpose thursday", args{"thrust", "thursday"}, 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleAlignmentNormDist(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("SimpleAlignmentNormDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
