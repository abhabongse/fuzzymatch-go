package canonical

import (
	"testing"
)

func TestStripNonPrintTransformer(t *testing.T) {
	StripNonPrint := func(str string) string { return ApplyTransformers(str, StripNonPrintTransformer) }
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
		{"non-normal spaces", args{"1\n2\r3"}, "1\n2\r3"},
		{"control codes", args{"Y\x01M\x02C\x1FA"}, "YMCA"},
		{"assigned thai", args{"x\xe0\xb8\x81\xe0\xb8\xb2x"}, "xกาx"},
		{"unassigned thai", args{"x\xe0\xb8\x80\xe0\xb8\xbcx"}, "xx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripNonPrint(tt.args.str); got != tt.want {
				t.Errorf("StripNonPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNormalSpaceTransformer(t *testing.T) {
	ToNormalSpace := func(str string) string { return ApplyTransformers(str, ToNormalSpaceTransformer) }
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"leading", args{"\t space in front"}, "  space in front"},
		{"trailing", args{"space at back \r\t "}, "space at back    "},
		{"leading & trailing", args{" \t lonely words \r  "}, "   lonely words    "},
		{"inter-word", args{"well\t\ndone"}, "well  done"},
		{"no change", args{"this is perfectly fine"}, "this is perfectly fine"},
		{"all three", args{" \t100\r200\n "}, "  100 200  "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNormalSpace(tt.args.str); got != tt.want {
				t.Errorf("ToNormalSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAccentsTransformer(t *testing.T) {
	RemoveAccents := func(str string) string { return ApplyTransformers(str, RemoveAccentsTransformer) }
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

func TestToLowerTransformer(t *testing.T) {
	ToLower := func(str string) string { return ApplyTransformers(str, ToLowerTransformer) }
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
		{"thai", args{"การทำงาน"}, "การทำงาน"},
		{"greek & cyrillic", args{"ΣΩ ЂЩ"}, "σω ђщ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.args.str); got != tt.want {
				t.Errorf("ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
