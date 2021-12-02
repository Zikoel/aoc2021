package utils

// BinarySplit 05
func BinarySplit(code string, space []int) int {
	head := string(code[0])

	// fmt.Printf("code %s space: %v\n", code, space)

	if len(space) == 2 {
		if head == "F" || head == "L" {
			return space[0]
		}
		return space[1]
	}

	half := len(space) / 2

	if head == "B" || head == "R" {
		return BinarySplit(code[1:], space[half:])
	}
	return BinarySplit(code[1:], space[:half])
}

// RemoveDuplicate from a string
func RemoveDuplicate(slice string) string {
	keys := make(map[rune]bool)
	result := ""

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result += string(entry)
		}
	}
	return result
}

// ListContains check if list contains the value val
func ListContains(list []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}

	return false
}

// ListElementsSum sum all element of a list
func ListElementsSum(list []int) int {
	result := 0
	for _, val := range list {
		result += val
	}

	return result
}

// ListMax return the max value contained on the list
func ListMax(list []int) int {
	result := list[0]
	for _, val := range list {
		if val > result {
			result = val
		}
	}

	return result
}

// ListMin return the min value contained on the list
func ListMin(list []int) int {
	result := list[0]
	for _, val := range list {
		if val < result {
			result = val
		}
	}

	return result
}

// ListRemoveVal Remove all occurrence of val from list
func ListRemoveVal(list []int, val int) []int {
	result := []int{}
	for _, v := range list {
		if v == val {
			continue
		}

		result = append(result, v)
	}

	return result
}

// ListRemoveValsList Remove all occurrence of vals from list
func ListRemoveValsList(list, vals []int) []int {
	result := []int{}
	for _, v := range list {
		if ListContains(vals, v) {
			continue
		}

		result = append(result, v)
	}

	return result
}

// ListCount Count the occurence of specified val
func ListCount(list []int, val int) int {
	result := 0
	for _, v := range list {
		if v == val {
			result++
		}
	}

	return result
}
