package day06

import (
	"reflect"
	"testing"
)

func Test_ProcInputSets(t *testing.T) {
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
			gotGroups := ProcInputSets(tt.args.in)
			// just check the summed values to check they're correct:
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
				t.Errorf("ProcInputSets() = %v, want %v", iGroups, tt.wantGroups)
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

func TestProcInputFull(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name      string
		args      args
		wantMsets []string
		wantPpl   []int
	}{
		{
			"return correct 2d rune slices for given input",
			args{[]string{"abc", "", "a", "b", "c", "", "a", "a", "", "ab", "ad"}},
			[]string{"abc", "abc", "aa", "abad"},
			[]int{1, 3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsets, gotPpl := ProcInputFull(tt.args.in)
			if !reflect.DeepEqual(gotMsets, tt.wantMsets) {
				t.Errorf("ProcInputFull() gotMsets = %v, want %v", gotMsets, tt.wantMsets)
			}
			if !reflect.DeepEqual(gotPpl, tt.wantPpl) {
				t.Errorf("ProcInputFull() gotPpl = %v, want %v", gotPpl, tt.wantPpl)
			}
		})
	}
}

func TestAnsTwo(t *testing.T) {
	type args struct {
		sets  [][]rune
		msets []string
		ppl   []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
		wantErr bool
	}{
		{
			"return correct sum when inputs are good",
			args{
				[][]rune{{'a', 'b', 'c'}, {'a', 'b', 'c'}, {'a', 'b', 'c'}, {'a'}, {'b'}},
				[]string{"abc", "abc", "abac", "aaaa", "b"},
				[]int{1, 3, 2, 4, 1},
			},
			6,
			false,
		},
		{
			"return error when input slice lengths are not equal",
			args{
				[][]rune{{'a', 'b'}, {'a'}},
				[]string{"asdff", "asdf"},
				[]int{1, 2, 3},
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, err := AnsTwo(tt.args.sets, tt.args.msets, tt.args.ppl)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnsTwo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSum != tt.wantSum {
				t.Errorf("AnsTwo() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
