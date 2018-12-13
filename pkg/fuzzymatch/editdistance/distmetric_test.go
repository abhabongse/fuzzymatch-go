package editdistance

import (
	"github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch/runedata"
	"testing"
)

func TestFuzzySubstErrorDist(t *testing.T) {
	type args struct {
		c rune
		d rune
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"identical", args{'ก', 'ก'}, 0},
		{"totally different #1", args{'ก', 'ข'}, 1},
		{"totally different #2", args{'ก', 'ข'}, 1},
		{"similar consonant #1", args{'ศ', 'ษ'}, 0.9},
		{"similar consonant #2", args{'ษ', 'ศ'}, 0.9},
		{"missing tonal #1", args{0, runedata.ThaiCharacterMaiEk}, 0.6},
		{"missing tonal #2", args{runedata.ThaiCharacterMaiEk, 0}, 0.6},
		{"missing tonal #1", args{runedata.ThaiCharacterMaiTho, runedata.ThaiCharacterMaiEk}, 0.6},
		{"missing tonal #2", args{runedata.ThaiCharacterMaiEk, runedata.ThaiCharacterMaiTho}, 0.6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuzzySubstErrorDist(tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("FuzzySubstErrorDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFuzzyTransErrorDist(t *testing.T) {
	type args struct {
		c rune
		d rune
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"identical", args{'ก', 'ก'}, 0},
		{"totally different #1", args{'ก', 'ข'}, 1},
		{"totally different #2", args{'ก', 'ข'}, 1},
		{"sara am with tonal #1", args{runedata.ThaiCharacterMaiTri, runedata.ThaiCharacterSaraAm}, 0.3},
		{"sara am with tonal #2", args{runedata.ThaiCharacterSaraAm, runedata.ThaiCharacterMaiTri}, 0.3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuzzyTransErrorDist(tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("FuzzyTransErrorDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
