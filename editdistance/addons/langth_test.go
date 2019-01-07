package addons

import (
	"github.com/abhabongse/fuzzymatch-go/runedata"
	"math"
	"testing"
)

func TestThaiSubstitutionErrorDist(t *testing.T) {
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
			if got := ThaiSubstitutionErrorDist(tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("ThaiSubstitutionErrorDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThaiTranspositionErrorDist(t *testing.T) {
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
			if got := ThaiTranspositionErrorDist(tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("ThaiTranspositionErrorDist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThaiOptimalAlignmentDistance(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"English: weekends", args{"saturday", "sunday"}, 3},
		{"English: greetings", args{"hello", "hola"}, 3},
		{"English: empty", args{"", "hi"}, 2},
		{"English: transpose thursday", args{"thrust", "thursday"}, 4},
		{"Thai: unit substitution #1", args{"กขคงจฉมยร", "คงจฉชมยรลว"}, 5},
		{"Thai: unit substitution #2", args{"สองสามสี่", "ลองถามพี่"}, 3},
		{"Thai: non-unit substitution #1", args{"สองสามสี่", "ลองถามสิ"}, 3.6},
		{"Thai: non-unit substitution #2", args{"ไฟอากาศนำ้ค่ะ", "ใฝอากาสน้ำค้ะ"}, 3.6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThaiOptimalAlignmentDistance(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("ThaiOptimalAlignmentDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizedThaiOptimalAlignmentDistance(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"English: weekends", args{"saturday", "sunday"}, 0.375},
		{"English: greetings", args{"hello", "hola"}, 0.6},
		{"English: empty", args{"", "hi"}, 1},
		{"English: transpose thursday", args{"thrust", "thursday"}, 0.5},
		{"Thai: unit substitution #1", args{"กขคงจฉมยร", "คงจฉชมยรลว"}, 0.5},
		{"Thai: unit substitution #2", args{"สองสามสี่", "ลองถามพี่"}, 3.0 / 8.6},
		{"Thai: non-unit substitution #1", args{"สองสามสี่", "ลองถามสิ"}, 3.6 / 8.6},
		{"Thai: non-unit substitution #2", args{"ไฟอากาศนำ้ค่ะ", "ใฝอากาสน้ำค้ะ"}, 3.6 / 12.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizedThaiOptimalAlignmentDistance(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("NormalizedThaiOptimalAlignmentDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
