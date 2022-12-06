package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	input, _ := f.ReadFile("input")
	fmt.Println(p1(string(input)))
}

// This is also p2. I just swapped out the messageSize
func p1(line string) int {
	ans := 0
	messageSize := 14
bulk:
	for i := messageSize; i < len(line); i++ {
		pastFourLetters := line[i-messageSize : i]
		for _, l := range []rune(pastFourLetters) {
			if strings.Count(pastFourLetters, string(l)) > 1 {
				continue bulk
			}
		}
		ans = i
		break
	}
	return ans
}

/*
func p2(lines []string) int {
	ans := 0

	return ans
}
*/
