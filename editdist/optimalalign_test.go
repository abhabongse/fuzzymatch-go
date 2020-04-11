package editdist

import "testing"

func TestOptimalAlignmentDist(t *testing.T) {
	type args struct {
		fst               string
		snd               string
		substDistFunction RunePenaltyFunction
		transDistFunction RunePenaltyFunction
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"weekends", args{"saturday", "sunday", UnitPenalty, UnitPenalty}, 3},
		{"greetings", args{"hello", "hola", UnitPenalty, UnitPenalty}, 3},
		{"empty", args{"", "hi", UnitPenalty, UnitPenalty}, 2},
		{"transpose thursday", args{"thrust", "thursday", UnitPenalty, UnitPenalty}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptimalAlignmentDist(tt.args.fst, tt.args.snd, tt.args.substDistFunction, tt.args.transDistFunction); got != tt.want {
				t.Errorf("OptimalAlignmentDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
