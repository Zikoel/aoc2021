package day03

import (
	"aoc-2021/days/utils"
	"fmt"
	"strconv"
)

// Run run
func Run(data []byte) {
	values := utils.ListOfStrings(data)

	gamma := ""
	epsilon := ""
	for i := 0; i < 12; i++ {
		zeros := 0
		ones := 0
		for _, r := range values {
			v := r[i]
			if string(v) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaVal, err := strconv.ParseInt(gamma, 2, 64)
	epsilonVal, er := strconv.ParseInt(epsilon, 2, 64)

	if err != nil || er != nil {
		panic("")
	}

	fmt.Printf("%d\n", gammaVal*epsilonVal)

}
