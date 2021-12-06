package day06

import (
	"aoc-2021/days/utils"
	"fmt"
)

// func resolveLight(days int, timer int64) int64 {
// 	lights := []int64{timer}
// 	tmp := []int64{}
// 	news := []int64{}
// 	for i := 0; i < days; i++ {
// 		for _, v := range lights {
// 			if v == 0 {
// 				tmp = append(tmp, 6)
// 				news = append(news, 8)
// 			} else {
// 				tmp = append(tmp, v-1)
// 			}
// 		}
// 		lights = append([]int64{}, tmp...)
// 		lights = append(lights, news...)
// 		tmp = []int64{}
// 		news = []int64{}
// 	}

// 	return int64(len(lights))
// }

func fishGenerationCount(afterDay int64, startingInternalTimer int64) int64 {
	bornCalendar := make([]int64, afterDay)
	fishCounter := int64(1)

	for i := startingInternalTimer; i < int64(len(bornCalendar)); i += 7 {
		bornCalendar[i]++
	}

	for i := 0; i < len(bornCalendar); i++ {
		borns := bornCalendar[i]
		fishCounter += borns
		for j := i + 9; j < len(bornCalendar); j += 7 {
			bornCalendar[j] += borns
		}
	}

	return fishCounter
}

// Run run
func Run(data []byte) {
	fishs := utils.ListOfSeparatedInt64(data, ",")

	sum := int64(0)
	for _, f := range fishs {
		sum += fishGenerationCount(256, f)
	}

	fmt.Printf("%d\n", sum)
	// fmt.Printf("%d\n", fishGenerationCount(18, 2))
}
