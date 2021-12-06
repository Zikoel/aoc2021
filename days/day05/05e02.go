package day05

import (
	"aoc-2021/days/utils"
	"fmt"
)

func flipLine2(l line) line {
	return line{l.b, l.a}
}

func listAllLinePoints2(l line) []point {
	points := []point{}
	line := l

	// vertical
	if l.a.x == l.b.x {
		if l.a.y > l.b.y {
			line = flipLine2(l)
		}
		for y := line.a.y; y <= line.b.y; y++ {
			points = append(points, point{x: line.a.x, y: y})
		}
	}

	//horizontal
	if l.a.y == l.b.y {
		if l.a.x > l.b.x {
			line = flipLine2(l)
		}
		for x := line.a.x; x <= line.b.x; x++ {
			points = append(points, point{x: x, y: line.a.y})
		}
	}

	// diag
	if l.a.x != l.b.x && l.a.y != l.b.y {
		incX := int64(0)
		incY := int64(0)

		if l.a.x > l.b.x {
			incX = -1
		} else {
			incX = 1
		}

		if l.a.y > l.b.y {
			incY = -1
		} else {
			incY = 1
		}

		for x, y := line.a.x, line.a.y; x != line.b.x+incX && y != line.b.y+incY; x, y = x+incX, y+incY {
			points = append(points, point{x: x, y: y})
		}
	}

	return points
}

// Run run
func Run2(data []byte) {
	linesStr := utils.ListOfStrings(data)
	lines := []line{}
	for _, str := range linesStr {
		lines = append(lines, decodeLine(str))
	}

	points := make(map[string]int)
	for _, line := range lines {
		linePoints := listAllLinePoints2(line)
		for _, lp := range linePoints {
			points[lp.toString()]++
		}
	}

	count := 0
	for _, p := range points {
		if p >= 2 {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}
