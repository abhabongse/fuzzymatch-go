package editdist

import (
	"math"
	"testing"
)

func TestSimpleAlignmentDistance(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"weekends", args{"saturday", "sunday"}, 3},
		{"greetings", args{"hello", "hola"}, 3},
		{"empty", args{"", "hi"}, 2},
		{"transpose thursday", args{"thrust", "thursday"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleAlignmentDist(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("SimpleAlignmentDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
