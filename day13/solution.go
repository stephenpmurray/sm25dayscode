package day13

import (
	"strconv"
	"strings"
)

type Bus struct {
	id int
}

func GetInput(in []string) (t0 int, buses []Bus, err error) {
	t0, err = strconv.Atoi(in[0])
	if err != nil {
		return 0, []Bus{}, err
	}
	s := strings.Split(in[1], ",")
	for _, c := range s {
		if i, err := strconv.Atoi(c); err == nil {
			buses = append(buses, Bus{id: i})
		}
	}

	return t0, buses, nil
}

// part One:

func AnsOne(t0 int, buses []Bus) int {
	wait := buses[0].id
	id := 0

	for _, b := range buses {
		if w := CheckBus(t0, b); w < wait {
			wait = w
			id = b.id
		}
	}
	return wait * id
}

func CheckBus(t0 int, b Bus) int {
	return ((t0/b.id)*b.id + b.id) - t0
}
