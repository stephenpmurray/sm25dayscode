package day10

import (
	"errors"
	"log"
)

// list of adapters:
// can take  V-1,V-2,V-3, V+3
// Outlet Vo = 0

// problem 1:
// using EVERY device,
// find number of 1-volt diffs * 3-volt diffs
// Basic and Boring Solution:
// sort array
// iterate over array and sum diffs of 1 or 3

func AnsOne(sorted []int) (ans int, err error) {
	var diff, ones, threes int

	// insert a 0 for the power outlet
	sorted = append([]int{0}, sorted...)

	for i, v := range sorted[:len(sorted)-1] {
		diff = sorted[i+1] - v
		switch diff {
		case 0:
			break
		case 1:
			ones += 1
			break
		case 2:
			// n.b. input does not contain 2s(!)
			break
		case 3:
			threes += 1
			break
		default:
			return 0, errors.New("Cannot complete: joltage diff not in range 0-3")
		}
	}
	// increment threes for devices builtin:
	threes += 1
	return ones * threes, nil
}

// problem 2:
// find EVERY possible path from outlet to device
// what is the total number of paths?

// interesting solution:
// sort array of ints
// there are no diffs = 2, only 1s and 3s
// at each diff = 3, the path narrows to 1 possibility

func calcPaths(n int) int {
	// takes the number of ints in a sequence and returns the total
	// possibilities according to pn = p(n-1) + p(n-2) + p(n-3)
	p3 := 1
	p2 := 1
	p1 := 2
	p := p1 + p2 + p3

	switch n {
	case 0:
		log.Fatal("trying to calculate paths for 0-len sequence")
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	default:
		for i := 4; i <= n; i++ {
			p = p1 + p2 + p3
			p3 = p2
			p2 = p1
			p1 = p
		}
	}

	return p

}

func AnsTwo(sorted []int) int {
	prod := 1
	run := 1
	var diff int

	// insert a 0 for the power outlet
	sorted = append([]int{0}, sorted...)

	// iterate over array and look for runs of diff=1
	for i, v := range sorted[:len(sorted)-1] {
		diff = sorted[i+1] - v
		if diff == 1 {
			run += 1
		}
		if diff == 3 {
			prod *= calcPaths(run)
			run = 1
		}
	}

	return prod * calcPaths(run)
}
