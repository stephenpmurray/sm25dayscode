package day04

import (
	"reflect"
	"testing"
)

func TestPuzzle_ReadPassports(t *testing.T) {
	type fields struct {
		Records []Record
	}
	type args struct {
		inputs []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantRecord []Record
	}{
		{
			name: "Read input correctly to p.Record map",
			args: args{
				[]string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
					"",
					"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
					"hcl:#cfa07d byr:1929",
				},
			},
			wantRecord: []Record{
				{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd"},
				{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			p.ReadPassports(tt.args.inputs)
			if !reflect.DeepEqual(p.Records, tt.wantRecord) {
				t.Errorf("Test_Puzzle_ReadPassports(): want \n %v \n got \n %v", tt.wantRecord, p.Records)
			}
		})
	}
}

func TestPuzzle_AnsOne(t *testing.T) {
	type fields struct {
		Records []Record
	}
	tests := []struct {
		name   string
		fields fields
		wantC  int
	}{
		{
			"Returns correct count for a given p.Records field",
			fields{
				[]Record{
					{"ecl": "true", "pid": "true", "eyr": "true", "hcl": "true", "byr": "true", "iyr": "true", "cid": "true", "hgt": "true"},
					{"iyr": "true", "ecl": "true", "cid": "true", "eyr": "true", "pid": "true", "hcl": "true", "byr": "true"},
					{"ecl": "true", "pid": "true", "eyr": "true", "hcl": "true", "byr": "true", "iyr": "true", "hgt": "true"},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			if gotC := p.AnsOne(); gotC != tt.wantC {
				t.Errorf("Puzzle.AnsOne() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestPuzzle_isYrOK(t *testing.T) {
	type fields struct {
		Records []Record
	}
	type args struct {
		in  string
		low int
		hgh int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Check byr in range low",
			args:    args{"1920", 1920, 2002},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Check iyr out of range low",
			args:    args{"1945", 2010, 2020},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Check eyr out of range high",
			args:    args{"2031", 2020, 2030},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Return error when string cannot convert to int",
			args:    args{"asdf", 2020, 2030},
			want:    false,
			wantErr: true,
		},
		{
			name:    "check iyr val for last example",
			args:    args{"2010", 2010, 2020},
			want:    true,
			wantErr: false,
		},
		{
			name:    "check byr val for last example",
			args:    args{"1944", 1920, 2002},
			want:    true,
			wantErr: false,
		},
		{
			name:    "check eyr val for last example",
			args:    args{"2021", 2020, 2030},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			got, err := p.isYrOK(tt.args.in, tt.args.low, tt.args.hgh)
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.isYrOK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Puzzle.isYrOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_isHeightOK(t *testing.T) {
	type fields struct {
		Records []Record
	}
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Return success with good cm input",
			args:    args{"166cm"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Return error with bad input",
			args:    args{"zzzz"},
			want:    false,
			wantErr: true,
		}, {
			name:    "Return failure with OOB inches input",
			args:    args{"80in"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Return success with extra input",
			args:    args{"158cm"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "return failure when no units given",
			args:    args{"83"},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			got, err := p.isHeightOK(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.isHeightOK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Puzzle.isHeightOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_isPIDOK(t *testing.T) {
	type fields struct {
		Records []Record
	}
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "return success for good input",
			args:    args{"773498192"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "return error if input cannot convert to int",
			args:    args{"773d981a2"},
			want:    false,
			wantErr: true,
		},
		{
			name:    "return false when len(input) != 9",
			args:    args{"7734981921"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "return success for extra input",
			args:    args{"093154719"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			got, err := p.isPIDOK(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.isPIDOK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Puzzle.isPIDOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_isHCLOK(t *testing.T) {
	type fields struct {
		Records []Record
	}
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "return success for good input",
			args:    args{"#1f2a"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "return false if does not start with a hash",
			args:    args{"1f2a"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "return error when cannot convert to hex",
			args:    args{"#12fh"},
			want:    false,
			wantErr: true,
		},
		{
			name:    "return success for good input",
			args:    args{"#b6652a"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			got, err := p.isHCLOK(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.isHCLOK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Puzzle.isHCLOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_isECLOK(t *testing.T) {
	type fields struct {
		Records []Record
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "return success when first is match",
			args: args{"amb"},
			want: true,
		},
		{
			name: "return success when last is match",
			args: args{"oth"},
			want: true,
		},
		{
			name: "return failure when no match",
			args: args{"asf"},
			want: false,
		},
		{
			name: "return success for extra input",
			args: args{"blu"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			if got := p.isECLOK(tt.args.in); got != tt.want {
				t.Errorf("Puzzle.isECLOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPuzzle_AnsTwo(t *testing.T) {
	type fields struct {
		Records []Record
	}
	tests := []struct {
		name    string
		fields  fields
		wantC   int
		wantErr bool
	}{
		{
			"return failure for all invalid passports",
			fields{
				[]Record{
					{"eyr": "1972", "cid": "100", "hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926"},
					{"iyr": "2019", "hcl": "#602927", "eyr": "1967", "hgt": "170cm", "ecl": "grn", "pid": "012533040", "byr": "1946"},
					{"hcl": "dab227", "iyr": "2012", "ecl": "brn", "hgt": "182cm", "pid": "021572410", "eyr": "2020", "byr": "1992", "cid": "277"},
					{"hgt": "59cm", "ecl": "zzz", "eyr": "2038", "hcl": "74454a", "iyr": "2023", "pid": "3556412378", "byr": "2007"},
				},
			},
			0,
			false,
		},
		{
			"return success for valid passports",
			fields{
				[]Record{
					{"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012", "eyr": "2030", "byr": "1980", "hcl": "#623a2f"},
					{"eyr": "2029", "ecl": "blu", "cid": "129", "byr": "1989", "iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165cm"},
					{"hcl": "#888785", "hgt": "164cm", "byr": "2001", "iyr": "2015", "cid": "88", "pid": "545766238", "ecl": "hzl", "eyr": "2022"},
					{"iyr": "2010", "hgt": "158cm", "hcl": "#b6652a", "ecl": "blu", "byr": "1944", "eyr": "2021", "pid": "093154719"},
				},
			},
			4,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Puzzle{
				Records: tt.fields.Records,
			}
			gotC, err := p.AnsTwo()
			if (err != nil) != tt.wantErr {
				t.Errorf("Puzzle.AnsTwo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotC != tt.wantC {
				t.Errorf("Puzzle.AnsTwo() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
