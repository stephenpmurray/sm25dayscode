package day03

const tree = rune('#')
const space = rune('.')

type Puzzle struct {
	Map [][]rune
}

func (p *Puzzle) readTreeLine(line string) {
	p.Map = append(p.Map, []rune(line))
}

func (p *Puzzle) ReadAllTreeLines(lines []string) {
	for _, line := range lines {
		p.readTreeLine(line)
	}
}

func (p *Puzzle) Ans(strdR, strdC int) int {
	rdx := 0
	cdx := 0
	r := len(p.Map)
	c := len(p.Map[0])
	trees := 0

	for {
		// get to the bottom
		if rdx >= r {
			break
		}
		// wraparound
		if cdx >= c {
			cdx -= c
		}

		if p.Map[rdx][cdx] == tree {
			trees += 1
		}
		cdx += strdC
		rdx += strdR
	}
	return trees
}
