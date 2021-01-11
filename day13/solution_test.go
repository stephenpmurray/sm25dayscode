package day13

import (
	"reflect"
	"testing"
)

func TestCheckBus(t *testing.T) {
	type args struct {
		t0 int
		b  Bus
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"correctly identify wait for example bus 59",
			args{
				939, Bus{59},
			},
			5,
		},
		{
			"correctly identify wait for example bus 13",
			args{
				939, Bus{13},
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckBus(tt.args.t0, tt.args.b); got != tt.want {
				t.Errorf("CheckBus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnsOne(t *testing.T) {
	type args struct {
		t0    int
		buses []Bus
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"complete example",
			args{
				939,
				[]Bus{
					{7}, {13}, {59}, {31}, {19},
				},
			},
			295,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnsOne(tt.args.t0, tt.args.buses); got != tt.want {
				t.Errorf("AnsOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInput(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name      string
		args      args
		wantT0    int
		wantBuses []Bus
		wantErr   bool
	}{
		{
			"read in test input correctly",
			args{
				[]string{
					"939",
					"7,13,x,x,59,x,31,19",
				},
			},
			939,
			[]Bus{{7}, {13}, {59}, {31}, {19}},
			false,
		},
		{
			"return error when bad input",
			args{
				[]string{
					"asdf",
					"7,13,x,x,59,x,31,19",
				},
			},
			0,
			[]Bus{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT0, gotBuses, err := GetInput(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotT0 != tt.wantT0 {
				t.Errorf("GetInput() gotT0 = %v, want %v", gotT0, tt.wantT0)
			}
			if !reflect.DeepEqual(gotBuses, tt.wantBuses) {
				t.Errorf("GetInput() gotBuses = %v, want %v", gotBuses, tt.wantBuses)
			}
		})
	}
}
