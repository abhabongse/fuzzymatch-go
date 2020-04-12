package diacritics

import (
	"golang.org/x/text/transform"
	"testing"
)

func asciiFold(input string) string {
	output, _, _ := transform.String(AsciiFoldTransformer, input)
	return output
}

func TestAsciiFoldTransformer(t *testing.T) {
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
		{"special compatibility", args{"n² æ"}, "n2 ae"},
		{"latin untouched", args{"normal text"}, "normal text"},
		{"thai untouched", args{"ที่นู่นนั่นมีเป็นใช้"}, "ที่นู่นนั่นมีเป็นใช้"},
		{
			"long input",
			args{"café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig café könig n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ n² æ œ"},
			"cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig cafe konig n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe n2 ae oe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asciiFold(tt.args.str); got != tt.want {
				t.Errorf("asciiFold() = %v, want %v", got, tt.want)
			}
		})
	}
}
