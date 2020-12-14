package day06

func ProcInput(in []string) (groups [][]rune) {
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

func AnsTwo(groups [][]rune) (sum int) {
	sum = 0

	for _, g := range groups {
		sum += len(g)
	}

	return sum
}
