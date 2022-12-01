package main

import (
	"embed"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

//go:embed *
var f embed.FS

func partOne() {
	data, _ := f.ReadFile("input")
	splitted := strings.Split(string(data), "\n\n")

	biggest := 0
	for _, s := range splitted {
		total := 0
		for _, j := range strings.Split(s, "\n") {
			atoi, err := strconv.Atoi(j)
			if err != nil {
				continue
			}
			total+=atoi
		}
		if total > biggest {
			biggest = total
		}
	}
	fmt.Println(biggest)
}

func partTwo() {
	data, _ := f.ReadFile("input")
	splitted := strings.Split(string(data), "\n\n")

	elves := make([]int, 0, len(splitted))

	for _, s := range splitted {
		elf := 0
		for _, j := range strings.Split(s, "\n") {
			atoi, err := strconv.Atoi(j)
			if err != nil {
				continue
			}
			elf+=atoi
		}
		elves = append(elves, elf)
	}
	sort.Ints(elves)
	total := 0
	for i:=1; i<=3; i++ {
		total+=elves[len(elves)-i]
	}
	fmt.Println(total)
}

func main() {
	partOne()
	fmt.Println("--------")
	partTwo()
}
