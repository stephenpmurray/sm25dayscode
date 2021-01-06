package day11

import (
	"reflect"
)

type FloorPlan struct {
	Old  [][]rune
	rows int
	cols int
}

func (fp *FloorPlan) GetFLoorplan(input []string) {

	// populate floorplan struct from input []string
	for _, s := range input {
		fp.Old = append(fp.Old, []rune(s))

	}
	fp.rows = len(fp.Old)
	fp.cols = len(fp.Old[0]) // assume all rows same length

	return
}

func (fp *FloorPlan) applyRules(r, c int) (ret rune) {
	// function to apply rules to an index of fp.Old and return the new value
	var rdx, cdx int

	if fp.Old[r][c] == '.' {
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
			if fp.Old[rdx][cdx] == '#' {
				seats += 1
			}
		}
	}

	// apply rules:
	if fp.Old[r][c] == 'L' {
		if seats == 0 {
			ret = '#'
		} else {
			ret = 'L'
		}
	}
	if fp.Old[r][c] == '#' {
		if seats >= 4 {
			ret = 'L'
		} else {
			ret = '#'
		}

	}
	return
}

func (fp *FloorPlan) applyRulesAll(new [][]rune, rules func(int, int) rune) [][]rune {
	// apply rules
	for i := 0; i < fp.rows; i++ {
		for j := 0; j < fp.cols; j++ {
			new[i][j] = rules(i, j)
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
		new = fp.applyRulesAll(new, fp.applyRules)

		if reflect.DeepEqual(fp.Old, new) {
			return countSeats(fp.Old)
		} else {
			// note that a deep copy is required, DONT copy structs :))))
			fp.Old = new
			continue
		}

	}

}

// answer 2

func checkDirection(Old [][]rune, x, y, dx, dy int) bool {

	lr := len(Old)
	lc := len(Old[0])

	for {
		x += dx
		y += dy
		// check OOB
		if x < 0 || y < 0 || y >= lr || x >= lc {
			return false
		}
		if Old[y][x] == 'L' {
			return false
		}
		if Old[y][x] == '#' {
			return true
		}

	}
}

func (fp *FloorPlan) applyNewRules(r, c int) (ret rune) {

	if fp.Old[r][c] == '.' {
		ret = '.'
	}

	// count hashes (seats) in LOS
	seats := 0
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			if checkDirection(fp.Old, c, r, dx, dy) {
				// fmt.Println("direction true: ", dx, dy)
				seats += 1
			}
		}
	}

	// apply rules:
	if fp.Old[r][c] == 'L' {
		if seats == 0 {
			ret = '#'
		} else {
			ret = 'L'
		}
	}
	if fp.Old[r][c] == '#' {
		if seats >= 5 {
			ret = 'L'
		} else {
			ret = '#'
		}

	}
	return
}

// TODO - lots of duplication, maybe refactor it?

func (fp *FloorPlan) AnsTwo() int {
	// applies rules to to floorplan until it stabilises
	var new [][]rune

	for {
		// make a new empty array
		new = make([][]rune, fp.rows)
		for i := range new {
			new[i] = make([]rune, fp.cols)
		}
		new = fp.applyRulesAll(new, fp.applyNewRules)

		// for i, _ := range new {
		// 	fmt.Println(string(fp.Old[i]), " ", string(new[i]))
		// }
		// fmt.Print("\n")

		if reflect.DeepEqual(fp.Old, new) {
			return countSeats(fp.Old)
		} else {
			// note that a deep copy is required, DONT copy structs :))))
			fp.Old = new
			continue
		}

	}

}
