package day09

import (
	"testing"
)

func TestAnsOne(t *testing.T) {
	type args struct {
		wLen int
		ints []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
		wantErr bool
	}{
		{
			"fails summing test and returns correct index",
			args{
				wLen: 5,
				ints: []int{25, 47, 40, 62, 55, 65,
					95, 102, 117, 150, 182, 127, 219},
			},
			127,
			false,
		},
		{
			"returns correct value when fail window is at end of array",
			args{
				wLen: 5,
				ints: []int{25, 47, 40, 62, 55, 65,
					95, 102, 117, 150, 182, 127},
			},
			127,
			false,
		},
		{
			"returns correct value when fail window is at start of array",
			args{
				wLen: 5,
				ints: []int{95, 102, 117, 150, 182, 127},
			},
			127,
			false,
		},
		{
			"returns error when no failing window found",
			args{
				wLen: 5,
				ints: []int{25, 47, 40, 62, 55, 65,
					95, 102, 117, 150, 182},
			},
			0,
			true,
		},
		{
			"test problem at line 155",
			args{
				25,
				[]int{
					603, 573, 728, 591, 681, 1153, 631, 646, 659, 1082,
					996, 926, 779, 795, 827, 1094, 975, 1079, 1382, 1359,
					1761, 1081, 1194, 1807, 1272, 1863},
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns, err := AnsOne(tt.args.wLen, tt.args.ints)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnsOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAns != tt.wantAns {
				t.Errorf("AnsOne() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func TestCheckWindow(t *testing.T) {
	type args struct {
		wLen int
		n    int
		win  []int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"return true when there is at least one match",
			args{
				5,
				182,
				[]int{65, 95, 102, 117, 150},
			},
			true,
			false,
		},
		{
			"return false when there is no match",
			args{
				5,
				127,
				[]int{95, 102, 117, 150, 182},
			},
			false,
			false,
		},
		{
			"return error when wlen doesnt match len(win)",
			args{5, 127, []int{12, 14, 11}},
			false,
			true,
		},
		{
			"test problem at line 155",
			args{
				25,
				1863,
				[]int{
					603, 573, 728, 591, 681, 1153, 631, 646, 659, 1082,
					996, 926, 779, 795, 827, 1094, 975, 1079, 1382, 1359,
					1761, 1081, 1194, 1807, 1272},
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckWindow(tt.args.wLen, tt.args.n, tt.args.win)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckWindow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindRun(t *testing.T) {
	type args struct {
		target int
		ints   []int
	}
	tests := []struct {
		name      string
		args      args
		wantStart int
		wantEnd   int
	}{
		{
			"returns correct indices of sum run matching target",
			args{
				127,
				[]int{35, 20, 15, 25, 47, 40, 62, 55, 65},
			},
			2,
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd := FindRun(tt.args.target, tt.args.ints)
			if gotStart != tt.wantStart {
				t.Errorf("FindRun() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("FindRun() gotEnd = %v, want %v", gotEnd, tt.wantEnd)
			}
		})
	}
}

func TestGetMinMax(t *testing.T) {
	type args struct {
		start int
		end   int
		arr   []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		{
			"return correct min and max vals",
			args{
				2,
				5,
				[]int{8, 4, 2, 3, 6, 7, 12, 1},
			},
			2,
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := GetMinMax(tt.args.start, tt.args.end, tt.args.arr)
			if gotMin != tt.wantMin {
				t.Errorf("GetMinMax() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("GetMinMax() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
