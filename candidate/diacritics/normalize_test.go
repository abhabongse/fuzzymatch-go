package diacritics

import (
	"github.com/abhabongse/fuzzymatch-go/factory"
	"testing"
)

func TestStripDiacriticalMarkTransformer(t *testing.T) {
	stripDiacriticalMarks := factory.MakeStringTransformFunction(StripDiacriticalMarksTransformer)
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
