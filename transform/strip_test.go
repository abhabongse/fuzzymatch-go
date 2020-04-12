package transform

import (
	"golang.org/x/text/transform"
	"testing"
)

func stripNonPrint(input string) string {
	output, _, _ := transform.String(StripNonPrintTransformer, input)
	return output
}

func respace(input string) string {
	output, _, _ := transform.String(RespaceTransformer, input)
	return output
}

func TestStripNonPrintTransformer(t *testing.T) {
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
			if got := stripNonPrint(tt.args.str); got != tt.want {
				t.Errorf("stripNonPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRespaceTransformer(t *testing.T) {
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
			if got := respace(tt.args.str); got != tt.want {
				t.Errorf("respace() = %v, want %v", got, tt.want)
			}
		})
	}
}
