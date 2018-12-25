package candidacy

import (
	"reflect"
	"testing"
)

func TestGenerateSalutationDecomposites(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []Decomposite
	}{
		{"empty", args{""}, []Decomposite{{"", ""}}},
		{"unchanged", args{"สวัสดี"}, []Decomposite{
			{"", "สวัสดี"},
		}},
		{"basic", args{"นายกรัฐมนตรี"}, []Decomposite{
			{"", "นายกรัฐมนตรี"},
			{"นาย", "กรัฐมนตรี"},
		}},
		// TODO: add more comprehensive tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: sort slices before comparison
			if got := GenerateSalutationDecomposites(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSalutationDecomposites() = %v, want %v", got, tt.want)
			}
		})
	}
}
