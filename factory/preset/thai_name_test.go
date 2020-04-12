package preset

import (
	"math"
	"testing"
)

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
		{"days of week", args{"saturday", "sunday"}, 5.0 / 8.0},
		{"identical after norm #1", args{"\nการแข่งขันจำ schön", "การเเข่ง\x01ขันจํา schon"}, 1},
		{"identical after norm #2", args{"\nการแข่งขันจำ schœn", "การเเข่ง\x01ขันจํา schoen"}, 1},
		{"identical after norm #3", args{"ที่ของเรา", "ทีีีีี่่่่่่ของเรา"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThaiNameSimilarityScore(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("ThaiNameSimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
