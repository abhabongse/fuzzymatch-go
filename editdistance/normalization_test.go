package editdistance

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
			if got := NormalizedSimpleAlignmentDistance(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("NormalizedSimpleAlignmentDistance() = %v, want %v", got, tt.want)
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
