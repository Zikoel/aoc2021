package day01

import (
	"fmt"
	"strconv"
	"strings"
)

func shift(list []int64, val int64) []int64 {
	l := make([]int64, len(list))
	for idx := range l {
		if idx == 0 {
			l[idx] = val
		} else {
			l[idx] = list[idx-1]
		}
	}
	return l
}

func listSum(list []int64) int64 {
	result := int64(0)
	for _, v := range list {
		result += v
	}
	return result
}

// Run run
func Run2(data []byte) {
	valuesStr := strings.Split(string(data), "\n")
	values := []int64{}
	for _, str := range valuesStr {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic("Unable to convert value")
		}

		values = append(values, val)
	}

	window := make([]int64, 3)
	incCount := -1
	lastWindowSum := int64(0)
	for idx, v := range values {
		window = shift(window, v)

		if idx < len(window)-1 {
			continue
		}

		sum := listSum(window)

		if (sum - lastWindowSum) > 0 {
			incCount++
		}

		lastWindowSum = sum
	}

	fmt.Printf("Result: %d\n", incCount)
}
