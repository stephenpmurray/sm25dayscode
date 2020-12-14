package day06

import (
	"errors"
	"strings"
)

func ProcInputSets(in []string) (groups [][]rune) {
	set := make(map[rune]bool)
	rns := []rune{}

	for _, line := range in {
		if line == "" {
			// convert from map to runes and make new group:
			for k, _ := range set {
				rns = append(rns, k)
			}
			groups = append(groups, rns)
			set = make(map[rune]bool)
			rns = []rune{}
			continue
		}

		// add each char to set for group of people
		for _, ch := range []rune(line) {
			set[ch] = true
		}
	}
	// if last line is not newline
	for k, _ := range set {
		rns = append(rns, k)
	}
	groups = append(groups, rns)
	return groups
}

func AnsOne(groups [][]rune) (sum int) {
	sum = 0

	for _, g := range groups {
		sum += len(g)
	}

	return sum
}

func ProcInputFull(in []string) (msets []string, ppl []int) {
	// Take slice of strings, and concatenates strings
	// which are not separated by empty string. Counts the
	// number of lines appended to each group and returns as []ppl

	var group string
	c := 0

	for _, line := range in {
		if line == "" {
			msets = append(msets, group)
			group = ""
			ppl = append(ppl, c)
			c = 0
			continue
		}
		group = group + line
		c += 1
	}

	msets = append(msets, group)
	ppl = append(ppl, c)
	return msets, ppl
}

func AnsTwo(sets [][]rune, msets []string, ppl []int) (sum int, err error) {
	if len(sets) != len(msets) || len(sets) != len(ppl) {
		return 0, errors.New("Error: []sets, []msets and []ppl must be equal length")
	}

	for idx, str := range msets {
		for _, ch := range sets[idx] {
			if strings.Count(str, string(ch)) == ppl[idx] {
				sum += 1
			}
		}
	}

	return sum, nil
}
