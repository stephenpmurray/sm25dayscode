package day11

import (
	"reflect"
)

type FloorPlan struct {
	old  [][]rune
	rows int
	cols int
}

func (fp *FloorPlan) GetFLoorplan(input []string) {

	// populate floorplan struct from input []string
	for _, s := range input {
		fp.old = append(fp.old, []rune(s))

	}
	fp.rows = len(fp.old)
	fp.cols = len(fp.old[0]) // assume all rows same length

	return
}

func (fp *FloorPlan) applyRules(r, c int) (ret rune) {
	// function to apply rules to an index of fp.old and return the new value
	var rdx, cdx int

	if fp.old[r][c] == '.' {
		ret = '.'
	}

	// count hashes (seats) in surrounding area
	seats := 0
	for _, dc := range []int{-1, 0, 1} {
		for _, dr := range []int{-1, 0, 1} {
			rdx = r + dr
			cdx = c + dc

			// handle OOB cases:
			if rdx < 0 || cdx < 0 || rdx >= fp.rows || cdx >= fp.cols {
				continue
			}
			//handle self case:
			if dr == 0 && dc == 0 {
				continue
			}
			if fp.old[rdx][cdx] == '#' {
				seats += 1
			}
		}
	}

	// apply rules:
	if fp.old[r][c] == 'L' {
		if seats == 0 {
			ret = '#'
		} else {
			ret = 'L'
		}
	}
	if fp.old[r][c] == '#' {
		if seats >= 4 {
			ret = 'L'
		} else {
			ret = '#'
		}

	}
	return
}

func (fp *FloorPlan) applyRulesAll(new [][]rune) [][]rune {
	// apply rules
	for i := 0; i < fp.rows; i++ {
		for j := 0; j < fp.cols; j++ {
			new[i][j] = fp.applyRules(i, j)
		}

	}

	return new
}

func countSeats(in [][]rune) (c int) {
	// count all instances of '#' in a 2d rune array

	for _, line := range in {
		for _, r := range line {
			if r == '#' {
				c += 1
			}
		}
	}
	return c
}

func (fp *FloorPlan) AnsOne() int {
	// applies rules to to floorplan until it stabilises
	var new [][]rune

	for {
		// make a new empty array
		new = make([][]rune, fp.rows)
		for i := range new {
			new[i] = make([]rune, fp.cols)
		}
		new = fp.applyRulesAll(new)

		if reflect.DeepEqual(fp.old, new) {
			return countSeats(fp.old)
		} else {
			// note that a deep copy is required, DONT copy structs :))))
			fp.old = new
			continue
		}

	}

}
