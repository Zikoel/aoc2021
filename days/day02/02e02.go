package day02

import (
	"aoc-2021/days/utils"
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	x   int64
	y   int64
	aim int64
}

func decodeInstruction(istruction string) (string, int64) {
	parts := strings.Split(istruction, " ")
	val, err := strconv.ParseInt(parts[1], 10, 64)

	if err != nil {
		panic("decode error")
	}

	return parts[0], val
}

func applyIstruction(p position, istruction string) position {

	action, val := decodeInstruction(istruction)

	r := p

	switch action {
	case "forward":
		r.x += val
		r.y += r.aim * val
	case "down":
		r.aim += val
	case "up":
		r.aim -= val
	}

	return r
}

// Run run
func Run2(data []byte) {
	instructions := utils.ListOfStrings(data)

	p := position{0, 0, 0}
	for _, i := range instructions {
		p = applyIstruction(p, i)
	}

	fmt.Printf("%d\n", p.y*p.x)

}
