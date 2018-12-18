package editdistance

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
			if got := SimpleAlignmentDistance(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("SimpleAlignmentDistance() = %v, want %v", got, tt.want)
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
