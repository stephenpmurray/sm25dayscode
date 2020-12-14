package day03

import (
	"reflect"
	"testing"
)

func TestPuzzle_readTreeLine(t *testing.T) {
	type fields struct {
		Map [][]rune
	}
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantMap [][]rune
	}{
		{
			name:    "Read line correctly to Map struct",
			args:    args{"..##.."},
			wantMap: [][]rune{{'.', '.', '#', '#', '.', '.'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Map: tt.fields.Map,
			}
			p.readTreeLine(tt.args.line)
			if !reflect.DeepEqual(tt.wantMap, p.Map) {
				t.Errorf("Problem in TestPuzzle_readTreeLine() Got %v want %v", p.Map, tt.wantMap)
			}
		})
	}
}

func TestPuzzle_ReadAllTreeLines(t *testing.T) {
	type fields struct {
		Map [][]rune
	}
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantMap [][]rune
	}{
		{
			name:    "Correctly read input to p.Map[]",
			args:    args{[]string{"..##..", "##...."}},
			wantMap: [][]rune{{'.', '.', '#', '#', '.', '.'}, {'#', '#', '.', '.', '.', '.'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Map: tt.fields.Map,
			}
			p.ReadAllTreeLines(tt.args.lines)
			if !reflect.DeepEqual(tt.wantMap, p.Map) {
				t.Errorf("Problem in TestPuzzle_ReadAllTreeLine() Got %v want %v", p.Map, tt.wantMap)
			}
		})
	}
}

func TestPuzzle_Ans(t *testing.T) {
	type fields struct {
		Map [][]rune
	}
	type args struct {
		strdR int
		strdC int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Return correct sum",
			fields: fields{
				[][]rune{
					{'.', '.', '.', '.'},
					{'.', '.', '.', '#'},
				},
			},
			args: args{1, 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Map: tt.fields.Map,
			}
			if got := p.Ans(tt.args.strdR, tt.args.strdC); got != tt.want {
				t.Errorf("Puzzle.AnsOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
