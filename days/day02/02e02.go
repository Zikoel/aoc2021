package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type fda struct {
	f int64
	d int64
	a int64
}

func applyInstruction2(p fda, istruction string) fda {
	parts := strings.Split(istruction, " ")
	val, err := strconv.ParseInt(parts[1], 10, 64)

	if err != nil {
		panic("conv")
	}

	r := p
	switch parts[0] {
	case "forward":
		r.f += val
		r.d += r.a * val
	case "down":
		r.a += val
	case "up":
		r.a -= val
	}

	return r
}

// Run run
func Run2(data []byte) {
	instructions := strings.Split(string(data), "\n")

	p := fda{0, 0, 0}
	for _, i := range instructions {
		p = applyInstruction2(p, i)
	}

	fmt.Printf("%d\n", p.d*p.f)

}
