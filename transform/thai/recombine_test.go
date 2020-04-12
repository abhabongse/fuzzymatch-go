package thai

import (
	"golang.org/x/text/transform"
	"testing"
)

func bigramRecombine(input string) string {
	output, _, _ := transform.String(BigramRecombineTransformer, input)
	return output
}

func TestBigramRecombineTransformer(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"normal thai", args{"สำแดงแรงดี"}, "สำแดงแรงดี"},
		{"double sara-e", args{"สำเเดงแรงดี"}, "สำแดงแรงดี"},
		{"attach sara-am", args{"สําแดงแรงดี"}, "สำแดงแรงดี"},
		{"both kinds", args{"สําแดงเเรงดี"}, "สำแดงแรงดี"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bigramRecombine(tt.args.str); got != tt.want {
				t.Errorf("bigramRecombine() = %v, want %v", got, tt.want)
			}
		})
	}
}
