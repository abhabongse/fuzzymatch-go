package diacritics

import (
	"golang.org/x/text/transform"
	"testing"
)

func stripDiacriticalMarks(input string) string {
	output, _, _ := transform.String(StripDiacriticalMarksTransformer, input)
	return output
}

func TestStripDiacriticalMarkTransformer(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"combined", args{"café könig"}, "cafe konig"},
		{"decomposed", args{"café könig"}, "cafe konig"},
		{"special compatibility", args{"n²"}, "n2"},
		{"latin untouched", args{"normal text"}, "normal text"},
		{"thai untouched", args{"ที่นู่นนั่นมีเป็นใช้"}, "ที่นู่นนั่นมีเป็นใช้"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripDiacriticalMarks(tt.args.str); got != tt.want {
				t.Errorf("stripDiacriticalMarks() = %v, want %v", got, tt.want)
			}
		})
	}
}
