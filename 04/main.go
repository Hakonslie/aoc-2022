package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	input, _ := f.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	fmt.Println(p2(lines))
}

func convertDisregardError(n string) int {
	i, _ := strconv.Atoi(n)
	return i
}

func p1(lines []string) int {
	ans := 0
	for _, l := range lines {
		pair := strings.Split(l, ",")
		elfOne := strings.Split(pair[0], "-")
		elfTwo := strings.Split(pair[1], "-")

		overlapping := false
		if convertDisregardError(elfOne[0]) >= convertDisregardError(elfTwo[0]) && convertDisregardError(elfOne[1]) <= convertDisregardError(elfTwo[1]) {
			overlapping = true
		} else if convertDisregardError(elfTwo[0]) >= convertDisregardError(elfOne[0]) && convertDisregardError(elfTwo[1]) <= convertDisregardError(elfOne[1]) {
			overlapping = true
		}
		if overlapping {
			ans++
		}
	}
	return ans
}

func p2(lines []string) int {
	ans := 0
	for _, l := range lines {
		pair := strings.Split(l, ",")
		elfOne := strings.Split(pair[0], "-")
		elfTwo := strings.Split(pair[1], "-")

		elfOne1 := convertDisregardError(elfOne[0])
		elfOne2 := convertDisregardError(elfOne[1])
		elfTwo1 := convertDisregardError(elfTwo[0])
		elfTwo2 := convertDisregardError(elfTwo[1])

		overlapping := false
		if elfOne1 >= elfTwo1 && elfOne1 <= elfTwo2 {
			overlapping = true
		}
		if elfOne2 >= elfTwo1 && elfOne2 <= elfTwo2 {
			overlapping = true
		}
		if elfOne1 >= elfTwo1 && elfOne2 <= elfTwo2 {
			overlapping = true
		}
		if elfTwo1 >= elfOne1 && elfTwo2 <= elfOne2 {
			overlapping = true
		}
		fmt.Printf("pair: %s is overlapping: %v \n", pair, overlapping)

		if overlapping {
			ans++
		}
	}

	return ans

}
