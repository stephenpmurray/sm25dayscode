package day06

import (
	"reflect"
	"testing"
)

func Test_ProcInput(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name       string
		args       args
		wantGroups []int
	}{
		{
			"return correct 2d rune slices for given input",
			args{[]string{"abc", "", "a", "b", "c", "", "a", "a", "", "ab", "ad"}},
			// [][]rune{{'a', 'b', 'c'}, {'a', 'b', 'c'}, {'a'}},
			[]int{294, 294, 97, 295},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroups := ProcInput(tt.args.in)
			// juswt check the summed values to check they're correct:
			i := 0
			iGroups := []int{}
			for _, group := range gotGroups {
				i = 0
				for _, j := range group {
					i += int(j)
				}
				iGroups = append(iGroups, i)
			}
			if !reflect.DeepEqual(iGroups, tt.wantGroups) {
				t.Errorf("ProcInput() = %v, want %v", iGroups, tt.wantGroups)
			}
		})
	}
}

func Test_AnsOne(t *testing.T) {
	type args struct {
		groups [][]rune
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			"return correct sum for 2d rune array",
			args{[][]rune{{'a', 'b', 'c'}, {'a', 'b', 'c'}, {'a', 'b'}}},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := AnsOne(tt.args.groups); gotSum != tt.wantSum {
				t.Errorf("AnsOne() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
