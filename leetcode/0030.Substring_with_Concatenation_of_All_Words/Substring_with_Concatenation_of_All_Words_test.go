package Substring_with_Concatenation_of_All_Words

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"duplicate",
			args{
				"barfoofoobarmanbarbarbarfoofoofoobar",
				[]string{"foo", "bar"},
			},
			[]int{0, 6, 21, 30},
		},
		{
			"simple",
			args{
				"barfoothefoobarman",
				[]string{"foo", "bar"},
			},
			[]int{0, 9},
		},
		{
			"null",
			args{
				"wordgoodgoodgoodbestword",
				[]string{"word", "good", "best", "word"},
			},
			[]int{},
		},
		{
			"same",
			args{
				"aaaaaaaa",
				[]string{"aa", "aa", "aa"},
			},
			[]int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
