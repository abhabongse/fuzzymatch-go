package preset

import (
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
		{"empty & something", args{"      ", "sunday"}, 0},
		{"something & empty", args{"saturday", "  "}, 0},
		{"days of week", args{"saturday", "sunday"}, 5.0 / 8.0},
		{"identical after norm #1", args{"\nmemorization schön", "memor\x01ization schon"}, 1},
		{"identical after norm #2", args{"\nmemorization schön", "memor\x01ization schoen"}, 1},
		{"identical after norm #3", args{"HEIßEN \n now", "heissen now"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlainSimilarityScore(tt.args.fst, tt.args.snd); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("PlainSimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
