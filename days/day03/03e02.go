package day03

import (
	"aoc-2021/days/utils"
	"fmt"
	"strconv"
)

func filter(list []string, pos int, val string) []string {
	result := []string{}

	for _, v := range list {
		if string(v[pos]) == val {
			result = append(result, v)
		}
	}

	return result
}

func calc(vals []string, mostCommon string) string {

	currentList := vals
	result := ""
	for i := 0; i < 12; i++ {
		zeros := 0
		ones := 0
		for _, r := range currentList {
			v := r[i]
			if string(v) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if mostCommon == "1" {
			if ones >= zeros {
				currentList = filter(currentList, i, "1")
			} else {
				currentList = filter(currentList, i, "0")
			}
		}

		if mostCommon == "0" {
			if zeros <= ones {
				currentList = filter(currentList, i, "0")
			} else {
				currentList = filter(currentList, i, "1")
			}
		}

		if len(currentList) == 1 {
			result = currentList[0]
		}
	}

	return result
}

// Run run
func Run2(data []byte) {
	values := utils.ListOfStrings(data)

	o2r := calc(values, "1")
	co2r := calc(values, "0")

	o2Val, err := strconv.ParseInt(o2r, 2, 64)
	co2Val, er := strconv.ParseInt(co2r, 2, 64)

	fmt.Printf("%d , %d\n", o2Val, co2Val)

	if err != nil || er != nil {
		panic("")
	}

	fmt.Printf("%d\n", o2Val*co2Val)

}
