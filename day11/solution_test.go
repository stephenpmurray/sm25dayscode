package day11

import (
	"reflect"
	"testing"
)

func TestFloorPlan_GetFLoorplan(t *testing.T) {
	type fields struct {
		Old  [][]rune
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
				Old:  tt.fields.Old,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			p.GetFLoorplan(tt.args.input)
			if !reflect.DeepEqual(p.Old, tt.wantOld) {
				t.Errorf("Test GetFloorplan() want p.Old = %v \n got %v", tt.wantOld, p.Old)
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
		Old  [][]rune
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
				Old: [][]rune{
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
				Old: [][]rune{
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
				Old: [][]rune{
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
				Old: [][]rune{
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
				Old: [][]rune{
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
				Old:  tt.fields.Old,
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
		Old  [][]rune
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
				Old: [][]rune{
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
				Old:  tt.fields.Old,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			if got := fp.AnsOne(); got != tt.want {
				t.Errorf("FloorPlan.AnsOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloorPlan_applyNewRules(t *testing.T) {
	type fields struct {
		Old  [][]rune
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
			"stay unocccipied central seat",
			fields{
				[][]rune{
					{'L', '.', '.', '.', '#'},
					{'.', '.', '.', '#', '.'},
					{'#', '.', 'L', '.', '.'},
					{'.', '.', '.', '.', '.'},
					{'.', '.', '.', '#', '.'},
				},
				5,
				5,
			},
			args{2, 2},
			'L',
		},
		{
			"become unocccipied central seat",
			fields{
				[][]rune{
					{'L', '.', '.', '.', '#'},
					{'.', '.', '.', '#', '.'},
					{'#', '.', '#', 'L', '.'},
					{'.', '.', '#', '.', '.'},
					{'#', '.', '.', '.', '#'},
				},
				5,
				5,
			},
			args{2, 2},
			'L',
		},
		{
			"example 2",
			fields{
				[][]rune{
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
					{'.', 'L', '.', 'L', '.', '#', '.', '#', '.'},
					{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				},
				3,
				9,
			},
			args{1, 1},
			'L',
		},
		{
			"example 3",
			fields{
				[][]rune{
					{'.', '#', '.', '#', '.'},
					{'#', '.', '.', '.', '#'},
					{'.', '.', 'L', '.', '.'},
					{'#', '.', '.', '.', '#'},
					{'.', '#', '.', '#', '.'},
				},
				5,
				5,
			},
			args{2, 2},
			'#',
		},
		{
			"edge case 1",
			fields{
				[][]rune{
					{'#', '.', '#', '#'},
					{'#', '.', '#', '#'},
					{'#', '#', '.', '.'},
				},
				4,
				4,
			},
			args{0, 2},
			'L',
		},
		{
			"edge case 2",
			fields{
				[][]rune{
					{'#', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', '#'},
					{'#', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
					{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
					{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
					{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
					{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
					{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
					{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', '#'},
					{'#', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
					{'#', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', '#'},
				},
				10,
				10,
			},
			args{0, 3},
			'#',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := &FloorPlan{
				Old:  tt.fields.Old,
				rows: tt.fields.rows,
				cols: tt.fields.cols,
			}
			if gotRet := fp.applyNewRules(tt.args.r, tt.args.c); gotRet != tt.wantRet {
				t.Errorf("FloorPlan.applyNewRules() = %v, want %v", string(gotRet), string(tt.wantRet))
			}
		})
	}
}

func Test_checkDirection(t *testing.T) {
	type args struct {
		Old [][]rune
		x   int
		y   int
		dx  int
		dy  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"check on diagonal top right",
			args{
				Old: [][]rune{
					{'L', '.', '.', '.', '#'},
					{'.', '.', '.', '#', '.'},
					{'#', '.', 'L', '.', '.'},
					{'.', '.', '.', '.', '.'},
					{'.', '.', '.', '#', '.'},
				},
				x: 2, y: 2, dx: 1, dy: -1,
			},
			true,
		},
		{
			"check on right horizontal",
			args{
				Old: [][]rune{
					{'L', '.', '.', '.', '#'},
					{'.', '.', '.', '#', '.'},
					{'#', '.', '#', '.', '.'},
					{'.', '.', '.', '.', '.'},
					{'.', '.', '.', '#', '.'},
				},
				x: 2, y: 2, dx: 1, dy: 0,
			},
			false,
		},
		{
			"check on right edge case",
			args{
				Old: [][]rune{
					{'.', '.', '#', '.', '.'},
					{'.', '.', '#', '.', '.'},
					{'.', '.', '#', '.', '#'},
					{'.', '.', '#', '.', '.'},
					{'.', '.', '#', '.', '.'},
				},
				x: 4, y: 2, dx: 1, dy: 1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDirection(tt.args.Old, tt.args.x, tt.args.y, tt.args.dx, tt.args.dy); got != tt.want {
				t.Errorf("checkDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
