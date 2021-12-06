package day04

import (
	"aoc-2021/days/utils"
	"fmt"
	"strconv"
	"strings"
)

type row struct {
	unmarked []int64
	marked   []int64
}

type board struct {
	rows []row
}

func (r row) isWinning() bool {
	return len(r.marked) >= 5
}

func (b board) isWinning() bool {
	for _, r := range b.rows {
		if r.isWinning() {
			return true
		}
	}
	return false
}

func (b board) unamerked() []int64 {
	result := []int64{}
	for _, r := range b.rows {
		result = append(result, r.unmarked...)
	}

	return result
}

func (r *row) markNumber(n int64) {
	if utils.ListContains64(r.unmarked, n) {
		r.unmarked = utils.ListRemoveVal64(r.unmarked, n)
		r.marked = append(r.marked, n)
	}
}

func (b *board) markNumber(n int64) {
	for i := range b.rows {
		b.rows[i].markNumber(n)
	}
}

func decodeNumbers(str string) []int64 {
	vals := []int64{}
	for _, s := range strings.Split(str, ",") {
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic("unable to convert")
		}
		vals = append(vals, val)
	}

	return vals
}

func decodeBoard(str string) board {
	rowsStr := strings.Split(str, "\n")
	b := board{[]row{}}
	for _, rowStr := range rowsStr {
		numbersStr := strings.Split(rowStr, " ")
		numbers := []int64{}
		for _, numberStr := range numbersStr {
			if numberStr == "" {
				continue
			}
			n, err := strconv.ParseInt(numberStr, 10, 64)
			if err != nil {
				panic("unable to convert number")
			}
			numbers = append(numbers, n)
		}

		b.rows = append(b.rows, row{numbers, []int64{}})
	}

	return b
}

func decodeInput(data []byte) ([]int64, []board) {
	dataStr := string(data)
	parts := strings.Split(dataStr, "\n\n")
	numbers := []int64{}
	boards := []board{}
	for i, v := range parts {
		if i == 0 {
			numbers = decodeNumbers(v)
		} else {
			b := decodeBoard(v)
			boards = append(boards, b)
		}
	}

	return numbers, boards
}

// Run run
func Run(data []byte) {
	extracted, boards := decodeInput(data)

	winnerFound := false
	for _, n := range extracted {
		for bidx := range boards {
			if winnerFound {
				continue
			}
			boards[bidx].markNumber(n)
			if boards[bidx].isWinning() {
				unmarked := boards[bidx].unamerked()
				sum := utils.ListElementsSum64(unmarked)
				fmt.Printf("%d\n", sum*n)
				winnerFound = true
			}
		}
	}
}
