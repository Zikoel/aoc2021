package day05

import (
	"aoc-2021/days/utils"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int64
	y int64
}

type line struct {
	a point
	b point
}

func (p point) toString() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func decodePoint(str string) point {
	parts := strings.Split(str, ",")
	x, err := strconv.ParseInt(parts[0], 10, 64)
	y, er := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || er != nil {
		panic("unable to convert")
	}

	return point{x, y}
}

func decodeLine(str string) line {
	parts := strings.Split(str, " -> ")
	a := decodePoint(parts[0])
	b := decodePoint(parts[1])
	return line{a, b}
}

func flipLine(l line, dir string) line {
	var resultLine line
	if dir == "h" {
		resultLine = line{point{l.b.x, l.a.y}, point{l.a.x, l.b.y}}
	}
	if dir == "v" {
		resultLine = line{point{l.a.x, l.b.y}, point{l.b.x, l.a.y}}
	}

	return resultLine
}

func listAllLinePoints(l line) []point {
	points := []point{}
	line := l
	if l.a.x == l.b.x {
		if l.a.y > l.b.y {
			line = flipLine(l, "v")
		}
		for y := line.a.y; y <= line.b.y; y++ {
			points = append(points, point{x: line.a.x, y: y})
		}
	}

	if l.a.y == l.b.y {
		if l.a.x > l.b.x {
			line = flipLine(l, "h")
		}
		for x := line.a.x; x <= line.b.x; x++ {
			points = append(points, point{x: x, y: line.a.y})
		}
	}

	return points
}

// Run run
func Run(data []byte) {
	linesStr := utils.ListOfStrings(data)
	lines := []line{}
	for _, str := range linesStr {
		lines = append(lines, decodeLine(str))
	}

	//fileter lines
	horVerLines := []line{}
	for _, line := range lines {
		if line.a.x == line.b.x || line.a.y == line.b.y {
			horVerLines = append(horVerLines, line)
		}
	}

	points := make(map[string]int)
	for _, line := range horVerLines {
		linePoints := listAllLinePoints(line)
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
