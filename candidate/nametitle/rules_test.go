package nametitle

import (
	"reflect"
	"sort"
	"testing"
)

func TestNamesWithoutTitles(t *testing.T) {
	type args struct {
		name string
	}
	SortedNamesWithoutTitles := func(name string) []string {
		results := GenerateNamesWithoutTitles(name)
		sort.Strings(results)
		return results
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty", args{""}, []string{""}},
		{"unchanged", args{"สวัสดี"}, []string{"สวัสดี"}},
		{"basic", args{"นายกรัฐมนตรี"}, []string{"กรัฐมนตรี", "นายกรัฐมนตรี"}},
		{"Thai ms/mrs #1", args{"นางสาวสยาม"}, []string{"นางสาวสยาม", "สยาม", "สาวสยาม"}},
		{"Thai ms/mrs #2", args{"นางสาว สยาม"}, []string{"นางสาว สยาม", "สยาม", "สาว สยาม"}},
		{"Thai ms/mrs #3", args{"นาง สาวสยาม"}, []string{"นาง สาวสยาม", "สาวสยาม"}},
		{"Thai ms/mrs #3", args{"นาง สาว สยาม"}, []string{"นาง สาว สยาม", "สาว สยาม"}},
		{"Thai mstr #1", args{"ดช ด่วน ทันใจ"}, []string{"ดช ด่วน ทันใจ", "ด่วน ทันใจ"}},
		{"Thai mstr #2", args{"ด.ช. ด่วน ทันใจ"}, []string{"ด.ช. ด่วน ทันใจ", "ด่วน ทันใจ"}},
		{"Thai mstr #3", args{"ด.ช.ด่วน ทันใจ"}, []string{"ด.ช.ด่วน ทันใจ", "ด่วน ทันใจ"}},
		{"English mr #1", args{"mr lightning speed"}, []string{"lightning speed", "mr lightning speed"}},
		{"English mr #2", args{"mr.lightning speed"}, []string{"lightning speed", "mr.lightning speed"}},
		{"English mr #3", args{"mr. lightning speed"}, []string{"lightning speed", "mr. lightning speed"}},
		{"English mister", args{"mister lightning speed"}, []string{"lightning speed", "mister lightning speed"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedNamesWithoutTitles(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateNamesWithoutTitles() = %v, want %v", got, tt.want)
			}
		})
	}
}
