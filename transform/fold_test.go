package transform

import (
	"golang.org/x/text/transform"
	"testing"
)

func caseFold(input string) string {
	output, _, _ := transform.String(CaseFoldingTransformer, input)
	return output
}

func TestCaseFoldingTransformer(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"basic latin", args{"AbCde FGH  ijkl 123"}, "abcde fgh  ijkl 123"},
		{"combined chars", args{"cafÉ kÖnig"}, "café könig"},
		{"decomposed chars", args{"cafÉ kÖnig"}, "café könig"},
		{"decomposed chars", args{"heißen"}, "heissen"},
		{"thai", args{"การทำงาน"}, "การทำงาน"},
		{"greek & cyrillic", args{"ΣΩ ЂЩ"}, "σω ђщ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseFold(tt.args.str); got != tt.want {
				t.Errorf("CaseFold() = %v, want %v", got, tt.want)
			}
		})
	}
}
