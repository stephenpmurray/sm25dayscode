package day12

import (
	"reflect"
	"testing"
)

func TestGetInput(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name    string
		args    args
		wantI   []Instruct
		wantErr bool
	}{
		{
			"correctly read in input",
			args{
				[]string{"F10", "N3", "F7", "R90", "F11"},
			},
			[]Instruct{
				{'F', 10}, {'N', 3}, {'F', 7}, {'R', 90}, {'F', 11},
			},
			false,
		},
		{
			"return error for bad input",
			args{
				[]string{"F10", "N3", "FHH", "R90", "F11"},
			},
			[]Instruct{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := GetInput(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("GetInput() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func Test_calcFacing(t *testing.T) {
	type args struct {
		curr int
		val  int
		dir  rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"correctly apply right facing",
			args{
				0, 90, 'R',
			},
			90,
		},
		{
			"correctly apply right facing > 360",
			args{
				270, 180, 'R',
			},
			90,
		},
		{
			"correctly apply left facing < 0",
			args{
				90, 270, 'L',
			},
			180,
		},
		{
			"ensure that 360 -> 0",
			args{
				270, 90, 'R',
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFacing(tt.args.curr, tt.args.val, tt.args.dir); got != tt.want {
				t.Errorf("calcFacing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_interpretInstruct(t *testing.T) {
	type fields struct {
		facing int
		x      int
		y      int
	}
	type args struct {
		i Instruct
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantPos Position
	}{

		{
			"go north",
			fields{90, 0, 0},
			args{
				Instruct{action: 'N', n: 5},
			},
			false,
			Position{
				facing: 90, x: 0, y: 5,
			},
		},
		{
			"go south",
			fields{90, 0, 0},
			args{
				Instruct{action: 'S', n: 11},
			},
			false,
			Position{
				facing: 90, x: 0, y: -11,
			},
		},
		{
			"go east",
			fields{90, 0, 0},
			args{
				Instruct{action: 'E', n: 11},
			},
			false,
			Position{
				facing: 90, x: 11, y: 0,
			},
		},
		{
			"go west",
			fields{90, 0, 0},
			args{
				Instruct{action: 'W', n: 5},
			},
			false,
			Position{
				facing: 90, x: -5, y: 0,
			},
		},
		{
			"forward east",
			fields{90, 0, 0},
			args{
				Instruct{action: 'F', n: 5},
			},
			false,
			Position{
				facing: 90, x: 5, y: 0,
			},
		},
		{
			"forward south",
			fields{180, 0, 0},
			args{
				Instruct{action: 'F', n: 5},
			},
			false,
			Position{
				facing: 180, x: 0, y: -5,
			},
		},
		{
			"forward west",
			fields{270, 0, 0},
			args{
				Instruct{action: 'F', n: 5},
			},
			false,
			Position{
				facing: 270, x: -5, y: 0,
			},
		},
		{
			"forward north",
			fields{0, 0, 0},
			args{
				Instruct{action: 'F', n: 5},
			},
			false,
			Position{
				facing: 0, x: 0, y: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				facing: tt.fields.facing,
				x:      tt.fields.x,
				y:      tt.fields.y,
			}
			if err := p.interpretInstruct(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Position.interpretInstruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantPos, *p) {
				t.Errorf("Position.interpretInstruct()got = %v, want %v", *p, tt.wantPos)
			}
		})
	}
}

func Test_manhattan(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"return correct val from negative nums",
			args{-7, -1},
			8,
		},
		{
			"return correct val from mixed nums",
			args{17, -1},
			18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manhattan(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("manhattan() = %v, want %v", got, tt.want)
			}
		})
	}
}
