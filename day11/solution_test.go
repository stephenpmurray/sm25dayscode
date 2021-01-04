package day11

import (
	"reflect"
	"testing"
)

func TestFloorPlan_GetFLoorplan(t *testing.T) {
	type fields struct {
		old  [][]rune
		rows int
		cols int
	}
	type args struct {
		input []string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantOld  [][]rune
		wantRows int
		wantCols int
	}{
		{
			name: "correctly read example input",
			args: args{
				[]string{
					"L.LL",
					"LLLL",
				},
			},
			wantOld: [][]rune{
				{'L', '.', 'L', 'L'},
				{'L', 'L', 'L', 'L'},
			},
			wantRows: 2,
			wantCols: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FloorPlan{
				old:  tt.fields.old,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			p.GetFLoorplan(tt.args.input)
			if !reflect.DeepEqual(p.old, tt.wantOld) {
				t.Errorf("Test GetFloorplan() want p.Old = %v \n got %v", tt.wantOld, p.old)
			}
			if p.rows != tt.wantRows {
				t.Errorf("Test GetFloorplan() want p.rows = %v \n got %v", tt.wantRows, p.rows)
			}
			if p.cols != tt.wantCols {
				t.Errorf("Test GetFloorplan() want p.cols = %v \n got %v", tt.wantCols, p.cols)
			}
		})
	}
}

func TestFloorPlan_applyRules(t *testing.T) {
	type fields struct {
		old  [][]rune
		rows int
		cols int
	}
	type args struct {
		r int
		c int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRet rune
	}{
		{
			"full > empty",
			fields{
				old: [][]rune{
					{'#', '.', 'L'},
					{'L', '#', '.'},
					{'#', '#', '#'},
				},
				rows: 3,
				cols: 3,
			},
			args{1, 1},
			'L',
		},
		{
			"empty > full",
			fields{
				old: [][]rune{
					{'#', '.', 'L'},
					{'L', 'L', '.'},
					{'#', '#', '#'},
				},
				rows: 3,
				cols: 3,
			},
			args{0, 2},
			'#',
		},
		{
			"handle corner case bottom right",
			fields{
				old: [][]rune{
					{'#', '.', 'L'},
					{'L', '#', '.'},
					{'#', '#', '#'},
				},
				rows: 3,
				cols: 3,
			},
			args{2, 2},
			'#',
		},
		{
			"don't change on empty space (.)",
			fields{
				old: [][]rune{
					{'#', '#', '#'},
					{'#', '#', '#'},
					{'#', '.', '#'},
				},
				rows: 3,
				cols: 3,
			},
			args{2, 1},
			'.',
		},
		{
			"empty seat stays empty when next to occupied",
			fields{
				old: [][]rune{
					{'.', 'L', '#'},
					{'.', '.', '.'},
					{'.', '.', '.'},
				},
				rows: 3,
				cols: 3,
			},
			args{0, 1},
			'L',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := &FloorPlan{
				old:  tt.fields.old,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			if gotRet := fp.applyRules(tt.args.r, tt.args.c); gotRet != tt.wantRet {
				t.Errorf("FloorPlan.applyRules() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_countSeats(t *testing.T) {
	type args struct {
		in [][]rune
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		{
			"correctly count seats in a floorplan",
			args{
				[][]rune{
					{'#', '#', '#', '#', '.'},
					{'.', '#', '#', '#', '#'},
					{'#', '#', '.', '#', '#'},
				},
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := countSeats(tt.args.in); gotC != tt.wantC {
				t.Errorf("countSeats() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestFloorPlan_AnsOne(t *testing.T) {
	type fields struct {
		old  [][]rune
		rows int
		cols int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"complete example successfully",
			fields{
				old: [][]rune{
					{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
					{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
					{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
					{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
					{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
					{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
					{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
					{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
					{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
					{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
				},
				rows: 10,
				cols: 10,
			},
			37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := &FloorPlan{
				old:  tt.fields.old,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			if got := fp.AnsOne(); got != tt.want {
				t.Errorf("FloorPlan.AnsOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
