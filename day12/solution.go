package day12

import (
	"errors"
	"strconv"
)

type Instruct struct {
	action rune
	n      int
}

type Position struct {
	facing int
	x      int
	y      int
}

func GetInput(in []string) (i []Instruct, err error) {
	var n int

	for _, s := range in {
		n, err = strconv.Atoi(s[1:])
		if err != nil {
			return []Instruct{}, err
		}
		i = append(i, Instruct{action: rune(s[0]), n: n})
	}
	return i, nil
}

func (p *Position) ApplyInstruct(inst []Instruct) (m int, err error) {
	// returns the Manhattan distance after applying the instructions
	p.facing = 90
	// for each instruction
	//interpret instruction to x, y vals
	for _, i := range inst {
		err = p.interpretInstruct(i)
		if err != nil {
			return 0, err
		}
	}

	// find the manhattan distance
	m = manhattan(p.x, p.y)

	return m, nil
}

func manhattan(x int, y int) int {
	return intAbs(x) + intAbs(y)
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (p *Position) interpretInstruct(i Instruct) (err error) {
	// interpret action rune and apply to struct

	// normal coords apply
	switch i.action {
	case 'N':
		p.y += i.n
		return nil
	case 'S':
		p.y -= i.n
		return nil
	case 'E':
		p.x += i.n
		return nil
	case 'W':
		p.x -= i.n
		return nil
	case 'F':
		// note that facings are all increments of 90
		// TODO: refactor to avoid code duplication:
		switch p.facing {
		case 0:
			p.y += i.n
			return nil
		case 90:
			p.x += i.n
			return nil
		case 180:
			p.y -= i.n
			return nil
		case 270:
			p.x -= i.n
			return nil
		}
	case 'L':
		p.facing = calcFacing(p.facing, i.n, i.action)
		return nil
	case 'R':
		p.facing = calcFacing(p.facing, i.n, i.action)
		return nil
	default:
		return errors.New(string(i.action) + "not a valid instruction for travel")
	}
	return errors.New("not a valid instruction for travel")
}

func calcFacing(curr, val int, dir rune) int {

	if dir == 'R' {
		curr += val
		// note: change 360 -> 0
		if curr >= 360 {
			curr -= 360
		}
	}
	if dir == 'L' {
		curr -= val
		if curr < 0 {
			curr += 360
		}
	}
	return curr
}

// part 2

// TODO: refactor to avoid repetition

func (p *Position) interpretInstruct2(i Instruct, w *Position) (err error) {

	//NOTE: w is a position relative to p
	// avoiding floating point maths

	// 270 is the same as -90, etc..
	if i.n == 270 && i.action == 'L' {
		i.n = 90
		i.action = 'R'
	}
	if i.n == 270 && i.action == 'R' {
		i.n = 90
		i.action = 'L'
	}

	switch i.action {
	case 'N':
		w.y += i.n
		return nil
	case 'S':
		w.y -= i.n
		return nil
	case 'E':
		w.x += i.n
		return nil
	case 'W':
		w.x -= i.n
		return nil
	case 'F':
		for j := 0; j < i.n; j++ {
			// check that this is performing correct number of iterations?
			p.x += w.x
			p.y += w.y
		}
		return nil
	case 'L':
		// if 180, flip the signs
		if i.n == 180 || i.n == 0 {
			w.x = -w.x
			w.y = -w.y
			return nil
		}
		if intAbs(p.facing) == 90 {
			w.x, w.y = -w.y, w.x
			return nil
		}

	case 'R':
		// if 180, flip the signs
		if i.n == 180 || i.n == 0 {
			w.x = -w.x
			w.y = -w.y
			return nil
		}
		if intAbs(p.facing) == 90 {
			w.x, w.y = w.y, -w.x
			return nil
		}

	default:
		return errors.New("Could not parse i.action")
	}

	return errors.New("Could not parse i.action")
}

func (p *Position) ApplyInstruct2(inst []Instruct, w Position) (m int, err error) {
	// returns the Manhattan distance after applying the instructions
	p.facing = 90
	w.x = 10
	w.y = 1

	// for each instruction
	//interpret instruction to x, y vals
	for _, i := range inst {
		err = p.interpretInstruct2(i, &w)
		if err != nil {
			return 0, err
		}
	}

	// find the manhattan distance
	m = manhattan(p.x, p.y)
	return m, nil
}
