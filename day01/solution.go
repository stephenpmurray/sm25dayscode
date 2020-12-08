package day01

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func importInput(fileName string) (values []int, err error) {
	var filePath string

	pwd, _ := os.Getwd()
	// check both this folder and day folder in case running as main
	if _, err := os.Stat(pwd + "/" + fileName); os.IsNotExist(err) {
		filePath = pwd + "/day01/" + fileName
	} else {
		filePath = pwd + "/" + fileName
	}
	if err != nil {
		return nil, err
	}

	// open fstream
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err

	}
	defer file.Close()

	// read in lines of file
	s := bufio.NewScanner(file)

	for s.Scan() {
		lineInt, _ := strconv.Atoi(s.Text())
		values = append(values, lineInt)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

func ansOne(values []int) (summed [2]int, err error) {
	lenV := len(values)

	// naive approach
	for idx, n := range values[:lenV] {
		for _, m := range values[idx+1:] {
			if n+m == 2020 {
				return [2]int{n, m}, nil
			}
		}

	}
	return summed, errors.New("didn't find a pair that sums to 2020")
}

func ansTwo(values []int) (summed [3]int, err error) {
	// naive approach
	for idx, n := range values[:] {
		for jdx, m := range values[idx+1:] {
			for _, o := range values[jdx+1:] {
				if n+m+o == 2020 {
					return [3]int{n, m, o}, nil
				}
			}
		}

	}
	return [3]int{}, errors.New("didnt find a triplet that sums to 2020")
}

func RunDay01() {
	// import input list to array
	values, err := importInput("input")
	if err != nil {
		log.Fatal("failed to import file, exiting...")
	}

	// answer 1
	summed, err := ansOne(values)
	if err != nil {
		os.Exit(1)
	}

	product := summed[0] * summed[1]
	fmt.Println(product)

	// answer 2
	triplet, err := ansTwo(values)
	if err != nil {
		os.Exit(1)
	}

	product = triplet[0] * triplet[1] * triplet[2]
	fmt.Println(product)

}
