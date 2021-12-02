package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type fd struct {
	f int64
	d int64
}

func applyInstruction(p fd, istruction string) fd {
	parts := strings.Split(istruction, " ")
	val, err := strconv.ParseInt(parts[1], 10, 64)

	if err != nil {
		panic("conv")
	}

	r := p
	switch parts[0] {
	case "forward":
		r.f += val
	case "down":
		r.d += val
	case "up":
		r.d -= val
	}

	return r
}

// Run run
func Run(data []byte) {
	instructions := strings.Split(string(data), "\n")

	p := fd{0, 0}
	for _, i := range instructions {
		p = applyInstruction(p, i)
	}

	fmt.Printf("%d\n", p.d*p.f)

}
