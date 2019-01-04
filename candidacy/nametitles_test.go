package candidacy

import (
	"reflect"
	"testing"
)

func TestSplitTitlesWithDefaultTitledNamePatterns(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []NameAndTitles
	}{
		{"empty", args{""}, []NameAndTitles{
			{"", "", ""},
		}},
		{"unchanged", args{"สวัสดี"}, []NameAndTitles{
			{"สวัสดี", "", ""},
		}},
		{"basic", args{"นายกรัฐมนตรี"}, []NameAndTitles{
			{"กรัฐมนตรี", "นาย", ""},
			{"นายกรัฐมนตรี", "", ""},
		}},
		{"Thai ms/mrs #1", args{"นางสาวสยาม"}, []NameAndTitles{
			{"นางสาวสยาม", "", ""},
			{"สยาม", "นางสาว", ""},
			{"สาวสยาม", "นาง", ""},
		}},
		{"Thai ms/mrs #2", args{"นางสาว สยาม"}, []NameAndTitles{
			{"นางสาว สยาม", "", ""},
			{"สยาม", "นางสาว", ""},
			{"สาว สยาม", "นาง", ""},
		}},
		{"Thai ms/mrs #3", args{"นาง สาวสยาม"}, []NameAndTitles{
			{"นาง สาวสยาม", "", ""},
			{"สาวสยาม", "นาง", ""},
		}},
		{"Thai mstr #1", args{"ดช ด่วน ทันใจ"}, []NameAndTitles{
			{"ดช ด่วน ทันใจ", "", ""},
			{"ด่วน ทันใจ", "ดช", ""},
		}},
		{"Thai mstr #2", args{"ด.ช. ด่วน ทันใจ"}, []NameAndTitles{
			{"ด.ช. ด่วน ทันใจ", "", ""},
			{"ด่วน ทันใจ", "ด.ช.", ""},
		}},
		{"Thai mstr #3", args{"ด.ช.ด่วน ทันใจ"}, []NameAndTitles{
			{"ด.ช.ด่วน ทันใจ", "", ""},
			{"ด่วน ทันใจ", "ด.ช.", ""},
		}},
		{"English mr #1", args{"mr lightning speed"}, []NameAndTitles{
			{"lightning speed", "mr", ""},
			{"mr lightning speed", "", ""},
		}},
		{"English mr #2", args{"mr.lightning speed"}, []NameAndTitles{
			{"lightning speed", "mr", ""},
			{"mr.lightning speed", "", ""},
		}},
		{"English mr #3", args{"mr. lightning speed"}, []NameAndTitles{
			{"lightning speed", "mr", ""},
			{"mr. lightning speed", "", ""},
		}},
		{"English mister", args{"mister lightning speed"}, []NameAndTitles{
			{"lightning speed", "mister", ""},
			{"mister lightning speed", "", ""},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitTitles(DefaultTitledNamePatterns, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultExtractBareNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
