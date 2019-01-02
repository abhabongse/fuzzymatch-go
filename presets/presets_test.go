package presets

import (
	"fmt"
	"math"
	"testing"
)

func TestPlainSimilarityScore(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty & empty", args{"", ""}, 1},
		{"empty & something", args{"", "sunday"}, 0},
		{"something & empty", args{"saturday", ""}, 0},
		{"days of week", args{"saturday", "sunday"}, 13.0 / 24.0},
		{"identical after norm #1", args{"\nการแข่งขันจำ schön", "การเเข่ง\x01ขันจํา schon"}, 1},
		{"identical after norm #2", args{"ที่ของเรา", "ทีีีีี่่่่่่ของเรา"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlainSimilarityScore(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("PlainSimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThaiNameSimilarityScore(t *testing.T) {
	type args struct {
		fst string
		snd string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty & empty", args{"", ""}, 1},
		{"empty & something", args{"", "sunday"}, 0},
		{"something & empty", args{"saturday", ""}, 0},
		{"days of week", args{"saturday", "sunday"}, 11.0 / 20.0},
		{"identical after norm #1", args{"\nการแข่งขันจำ schön", "การเเข่ง\x01ขันจํา schon"}, 1},
		{"identical after norm #2", args{"ที่ของเรา", "ทีีีีี่่่่่่ของเรา"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThaiNameSimilarityScore(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("ThaiNameSimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSimilarityScore() {
	// Find the similarity score between "saturday" and "sunday"
	score := PlainSimilarityScore("saturday", "sunday")
	fmt.Println(score)
}
