package main

import (
	"aoc-2021/days/day01"
	"aoc-2021/days/day02"
	"aoc-2021/days/day03"
	"aoc-2021/days/day04"
	"aoc-2021/days/day05"
	"aoc-2021/days/day06"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var day string = ""
	var file string = ""

	flag.StringVar(&day, "day", "", "day to run")
	flag.StringVar(&day, "d", "", "day to run (shorthand)")
	flag.StringVar(&file, "file", "", "file to  read input from")
	flag.StringVar(&file, "f", "", "file to  read input from (shorthand)")

	flag.Parse()

	if day == "" {
		flag.Usage()
		os.Exit(1)
	}

	var data []byte
	var err error
	if file == "" {
		data, err = ioutil.ReadAll(os.Stdin)
	} else if file != "" {
		data, err = ioutil.ReadFile(file)
	}

	if err != nil {
		log.Fatal("Impossible read input")
	}

	switch day {
	case "01e02":
		day01.Run2(data)
	case "02":
		day02.Run(data)
	case "02e02":
		day02.Run2(data)
	case "03":
		day03.Run(data)
	case "03e02":
		day03.Run2(data)
	case "04":
		day04.Run(data)
	case "05":
		day05.Run(data)
	case "05e02":
		day05.Run2(data)
	case "06":
		day06.Run(data)
	default:
		fmt.Println("Day not present")
	}
}
