package day09

import (
	"errors"
	"strconv"
)

func ProcInput(input []string) (ints []int, err error) {
	// convert strings to integers
	var i int

	for _, s := range input {
		i, err = strconv.Atoi(s)
		if err != nil {
			return []int{}, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func CheckWindow(wLen, n int, win []int) (bool, error) {

	if wLen != len(win) {
		return false, errors.New("wlen does not match win length")
	}

	for jdx, j := range win[:wLen-1] {
		for _, k := range win[jdx+1 : wLen] {
			if n == (j + k) {
				// everything ok
				return true, nil
			}
		}

	}
	// fail if there is no match:
	return false, nil
}

func AnsOne(wLen int, ints []int) (ans int, err error) {
	l := len(ints)
	var win []int
	var isOk bool

	// for each window:
	for idx, _ := range ints[:(l - wLen)] {
		win = ints[idx : idx+wLen]
		isOk, err = CheckWindow(wLen, ints[idx+wLen], win)
		if err != nil {
			return 0, err
		}
		if isOk {
			continue
		} else {
			// fmt.Println(idx, len(win), win)
			return ints[idx+wLen], nil

		}
	}

	return 0, errors.New("Could not find window failing summing test")
}

func FindRun(target int, ints []int) (start, end int) {
	// finds start and end indices of sum run equal to target

	// do 1st run to find target OR where target is exceeded
	acc := ints[start]
	for {
		end += 1
		acc += ints[end]
		if acc == target {
			return start, end
		}
		if acc > target {
			break
		}
	}

	for {
		// subtract start val and increment start index
		acc -= ints[start]
		start += 1
		if start == end {
			end += 1
			acc += ints[end]
		}

		// adjust end index and accumulator, check:
		switch {
		case acc == target:
			return start, end
		case end >= len(ints):
			return 0, 0
		case acc < target:
			for {
				end += 1
				acc += ints[end]
				if acc == target {
					return start, end
				}
				if acc > target {
					break
				}
			}
			break
		case acc > target:
			for {
				acc -= ints[end]
				end -= 1
				if acc == target {
					return start, end
				}
				if start == end {
					end += 1
					acc += ints[end]
					break
				}
				if acc < target {
					break
				}
			}
			break
		}

	}
}

func GetMinMax(start, end int, arr []int) (min, max int) {
	min = arr[start]
	max = arr[start]
	for _, v := range arr[start : end+1] {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}

	}
	return min, max
}
