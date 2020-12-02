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
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName

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

func checkList(values []int) (summed [2]int, err error) {

	for idx, n := range values {
		for _, m := range values[idx+1:] {
			if n+m == 2020 {
				return [2]int{n, m}, nil
			}
		}

	}
	return summed, errors.New("didn't find a pair that sums to 2020")
}

func main() {
	// import input list to array
	values, err := importInput("input")
	if err != nil {
		log.Fatal("failed to import file, exiting...")
	}

	// check list for matching pairs
	summed, err := checkList(values)
	if err != nil {
		fmt.Println(summed)
	}
}
