package day04

import (
	"strconv"
	"strings"
)

type Puzzle struct {
	Records []Record
}

type Record map[string]string

func (p *Puzzle) ReadPassports(inputs []string) {
	r := Record{}
	var kvs []string

	f := func(c rune) bool {
		return c == rune(':') || c == rune(' ')
	}

	for _, line := range inputs {

		if line == "" {
			p.Records = append(p.Records, r)
			r = Record{}
			continue
		}

		kvs = strings.FieldsFunc(line, f)
		// fmt.Printf("%v", kvs)
		for idx, key := range kvs {
			if idx%2 == 0 {
				r[key] = kvs[idx+1]
			}
		}

	}
	p.Records = append(p.Records, r)
}

func (p *Puzzle) AnsOne() (c int) {
	var flds int
	var cid bool
	for _, r := range p.Records {
		flds = len(r)
		_, cid = r["cid"]
		if (!cid && flds == 7) || flds == 8 {
			c += 1
		}

	}
	return c
}

func (p *Puzzle) isYrOK(in string, low, hgh int) (bool, error) {
	// check if Year is valid
	yr, err := strconv.Atoi(in)
	if err != nil {
		return false, err
	}

	if yr < low || yr > hgh {
		return false, nil
	}

	return true, nil
}

func (p *Puzzle) isHeightOK(in string) (bool, error) {
	// hgt (Height) - a number followed by either cm or in:
	// 		If cm, the number must be at least 150 and at most 193.
	// 		If in, the number must be at least 59 and at most 76.
	u := strings.TrimLeft(in, "0123456789")
	if u == "" {
		return false, nil
	}
	h, err := strconv.Atoi(strings.TrimRight(in, "cmin"))
	if err != nil {
		return false, err
	}
	if u == "cm" && (h < 150 || h > 193) {
		return false, nil
	}
	if u == "in" && (h < 59 || h > 76) {
		return false, nil
	}
	return true, nil
}

func (p *Puzzle) isHCLOK(in string) (bool, error) {
	// check Hair Colour starts with a hash and is valid hex
	hash := "#"
	_, err := strconv.ParseInt(strings.TrimLeft(in, hash), 16, 64)
	ok := strings.HasPrefix(in, hash)
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}

func (p *Puzzle) isECLOK(in string) bool {
	// check Eye Color is one of: amb blu brn gry grn hzl oth.
	set := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, s := range set {
		if in == s {
			return true
		}
	}
	return false
}

func (p *Puzzle) isPIDOK(in string) (bool, error) {
	// check that PID is 9 int chars long
	_, err := strconv.Atoi(in)
	if err != nil || len(in) != 9 {
		return false, err
	}
	return true, nil
}

func (p *Puzzle) AnsTwo() (c int, err error) {
	var flds int
	var cid, ok bool

	for _, r := range p.Records {
		// check that passports has valid number of fields
		flds = len(r)
		_, cid = r["cid"]
		if (cid && (flds == 7)) || (flds < 7) {
			continue
		}

		// byr (Birth Year) - four digits; at least 1920 and at most 2002.
		ok, err = p.isYrOK(r["byr"], 1920, 2002)
		if !ok || err != nil {
			continue
		}
		// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		ok, err = p.isYrOK(r["iyr"], 2010, 2020)
		if !ok || err != nil {
			continue
		}
		// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		ok, err = p.isYrOK(r["eyr"], 2020, 2030)
		if !ok || err != nil {
			continue
		}

		// hgt (Height) - a number followed by either cm or in
		ok, err = p.isHeightOK(r["hgt"])
		if !ok || err != nil {
			continue
		}

		// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		ok, err = p.isHCLOK(r["hcl"])
		if !ok || err != nil {
			continue
		}

		// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		ok = p.isECLOK(r["ecl"])
		if !ok {
			continue
		}

		// pid (Passport ID) - a nine-digit number, including leading zeroes.
		ok, err := p.isPIDOK(r["pid"])
		if !ok || err != nil {
			continue
		}

		// increase count if we get here
		c += 1

	}
	return c, nil
}
