package thai

import (
	"golang.org/x/text/transform"
	"testing"
)

func removedRepeatedMarks(input string) string {
	output, _, _ := transform.String(RemoveRepeatedMarksTransformer, input)
	return output
}


func TestRemoveRepeatedMarksTransformer(t *testing.T) {
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
		{
			"long input",
			args{"การทำงานที่บ้าน การทำงานที่่่่่่บ้าน มีีีีคนทำ ดูููสิ ห้้าหกปีิีทีีี่แล้วนูููู่่น การทำงานที่บ้าน การทำงานที่่่่่่บ้าน มีีีีคนทำ ดูููสิ ห้้าหกปีิีทีีี่แล้วนูููู่่น การทำงานที่บ้าน การทำงานที่่่่่่บ้าน มีีีีคนทำ ดูููสิ ห้้าหกปีิีทีีี่แล้วนูููู่่น การทำงานที่บ้าน การทำงานที่่่่่่บ้าน มีีีีคนทำ ดูููสิ ห้้าหกปีิีทีีี่แล้วนูููู่่น การทำงานที่บ้าน การทำงานที่่่่่่บ้าน มีีีีคนทำ ดูููสิ ห้้าหกปีิีทีีี่แล้วนูููู่่น"},
			"การทำงานที่บ้าน การทำงานที่บ้าน มีคนทำ ดูสิ ห้าหกปีิีที่แล้วนู่น การทำงานที่บ้าน การทำงานที่บ้าน มีคนทำ ดูสิ ห้าหกปีิีที่แล้วนู่น การทำงานที่บ้าน การทำงานที่บ้าน มีคนทำ ดูสิ ห้าหกปีิีที่แล้วนู่น การทำงานที่บ้าน การทำงานที่บ้าน มีคนทำ ดูสิ ห้าหกปีิีที่แล้วนู่น การทำงานที่บ้าน การทำงานที่บ้าน มีคนทำ ดูสิ ห้าหกปีิีที่แล้วนู่น",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removedRepeatedMarks(tt.args.str); got != tt.want {
				t.Errorf("removedRepeatedMarks() = %v, want %v", got, tt.want)
			}
		})
	}
}
