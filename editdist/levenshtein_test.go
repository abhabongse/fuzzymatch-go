package editdist

import "testing"

func TestLevenshteinDist(t *testing.T) {
	type args struct {
		fst               string
		snd               string
		substDistFunction RunePenaltyFunction
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"weekends", args{"saturday", "sunday", UnitPenalty}, 3},
		{"greetings", args{"hello", "hola", UnitPenalty}, 3},
		{"empty", args{"", "hi", UnitPenalty}, 2},
		{"transpose thursday", args{"thrust", "thursday", UnitPenalty}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevenshteinDist(tt.args.fst, tt.args.snd, tt.args.substDistFunction); got != tt.want {
				t.Errorf("LevenshteinDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
