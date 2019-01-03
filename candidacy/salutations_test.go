package candidacy

import (
	"reflect"
	"testing"
)

func TestDecomposeNameWithDefaultSalutationPatterns(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []Decomposite
	}{
		{"empty", args{""}, []Decomposite{
			{"", ""},
		}},
		{"unchanged", args{"สวัสดี"}, []Decomposite{
			{"", "สวัสดี"},
		}},
		{"basic", args{"นายกรัฐมนตรี"}, []Decomposite{
			{"", "นายกรัฐมนตรี"},
			{"นาย", "กรัฐมนตรี"},
		}},
		{"Thai ms/mrs #1", args{"นางสาวสยาม"}, []Decomposite{
			{"", "นางสาวสยาม"},
			{"นาง", "สาวสยาม"},
			{"นางสาว", "สยาม"},
		}},
		{"Thai ms/mrs #2", args{"นางสาว สยาม"}, []Decomposite{
			{"", "นางสาว สยาม"},
			{"นาง", "สาว สยาม"},
			{"นางสาว", "สยาม"},
		}},
		{"Thai ms/mrs #3", args{"นาง สาวสยาม"}, []Decomposite{
			{"", "นาง สาวสยาม"},
			{"นาง", "สาวสยาม"},
		}},
		{"Thai mstr #1", args{"ดช ด่วน ทันใจ"}, []Decomposite{
			{"", "ดช ด่วน ทันใจ"},
			{"ดช", "ด่วน ทันใจ"},
		}},
		{"Thai mstr #2", args{"ด.ช. ด่วน ทันใจ"}, []Decomposite{
			{"", "ด.ช. ด่วน ทันใจ"},
			{"ด.ช.", "ด่วน ทันใจ"},
		}},
		{"Thai mstr #3", args{"ด.ช.ด่วน ทันใจ"}, []Decomposite{
			{"", "ด.ช.ด่วน ทันใจ"},
			{"ด.ช.", "ด่วน ทันใจ"},
		}},
		{"English mr #1", args{"mr lightning speed"}, []Decomposite{
			{"", "mr lightning speed"},
			{"mr", "lightning speed"},
		}},
		{"English mr #2", args{"mr.lightning speed"}, []Decomposite{
			{"", "mr.lightning speed"},
			{"mr", "lightning speed"},
		}},
		{"English mr #3", args{"mr. lightning speed"}, []Decomposite{
			{"", "mr. lightning speed"},
			{"mr", "lightning speed"},
		}},
		{"English mister", args{"mister lightning speed"}, []Decomposite{
			{"", "mister lightning speed"},
			{"mister", "lightning speed"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecomposeName(tt.args.name, DefaultSalutationPatterns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecomposeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
