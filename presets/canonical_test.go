package presets

import "testing"

func TestThaiStringCanonicalize(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"unchanged", args{"sunday"}, "sunday"},
		{"thai unchanged", args{"การแข่งขัน"}, "การแข่งขัน"},
		{"fixed #1", args{"การเเข่งขันจําคำตอบที่่่่่่่่ café"}, "การแข่งขันจำคำตอบที่ cafe"},
		{"fixed #2", args{"การ\x01เเข่ง\x01ขันจําคำตอบที่่่่่่่่  \n\n cAFé\r"}, "การแข่งขันจำคำตอบที่ cafe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThaiStringCanonicalize(tt.args.str); got != tt.want {
				t.Errorf("ThaiStringCanonicalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
