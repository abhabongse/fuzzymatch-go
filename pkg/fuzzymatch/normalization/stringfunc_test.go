package normalization

import (
	"testing"
)

func TestReSpace(t *testing.T) {
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
			if got := ReSpace(tt.args.str); got != tt.want {
				t.Errorf("ReSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeThaiGrams(t *testing.T) {
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
			if got := RecombineThaiGrams(tt.args.str); got != tt.want {
				t.Errorf("RecombineThaiGrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveThaiRepeatedAccidents(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"no repeats", args{"การทำงานที่บ้าน"}, "การทำงานที่บ้าน"},
		{"repeated tonal", args{"การทำงานที่่่่่่บ้าน"}, "การทำงานที่บ้าน"},
		{"repeated asc vowel", args{"มีีีีคนทำ"}, "มีคนทำ"},
		{"repeated desc vowel", args{"ดูููสิ"}, "ดูสิ"},
		{"combinations", args{"ห้้าหกปีิีทีีี่แล้วนูููู่่น"}, "ห้าหกปีิีที่แล้วนู่น"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveThaiRepeatedAccidents(tt.args.str); got != tt.want {
				t.Errorf("RemoveThaiRepeatedAccidents() = %v, want %v", got, tt.want)
			}
		})
	}
}
