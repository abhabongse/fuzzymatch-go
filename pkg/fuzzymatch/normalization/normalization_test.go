package normalization

import (
	"testing"
)

func TestNormalizeWhiteSpaces(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"leading", args{"\t space in front"}, "space in front"},
		{"trailing", args{"space at back \r\t "}, "space at back"},
		{"leading & trailing", args{" \t lonely words \r  "}, "lonely words"},
		{"inter-word", args{"well\t\ndone"}, "well done"},
		{"no change", args{"this is perfectly fine"}, "this is perfectly fine"},
		{"all three", args{" \t100\r200\n "}, "100 200"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeWhiteSpaces(tt.args.str); got != tt.want {
				t.Errorf("NormalizeWhiteSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAccents(t *testing.T) {
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
			if got := RemoveAccents(tt.args.str); got != tt.want {
				t.Errorf("RemoveAccents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveNonPrinting(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"all good", args{"hello how good 123"}, "hello how good 123"},
		{"non-normal spaces", args{"1\n2\r3"}, "123"},
		{"control codes", args{"Y\x01M\x02C\x1FA"}, "YMCA"},
		{"assigned thai", args{"x\xe0\xb8\x81\xe0\xb8\xb2x"}, "xกาx"},
		{"unassigned thai", args{"x\xe0\xb8\x80\xe0\xb8\xbcx"}, "xx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveNonPrinting(tt.args.str); got != tt.want {
				t.Errorf("RemoveNonPrinting() = %v, want %v", got, tt.want)
			}
		})
	}
}
