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
	lines := strings.Split(string(input), "\n")
	fmt.Println(p2(lines))
}

func getPrio(r rune) int {
	var prio int
	if int(r) <= 95 {
		prio = (int(r) % 65) + 27
	} else {
		prio = (int(r) % 97) + 1
	}
	return prio
}

func p1(lines []string) int {
	ans := 0
lines:
	for _, line := range lines {
		compartmentOne := []rune(line)[:len(line)/2]
		compartmentTwo := string([]rune(line)[len(line)/2:])
		for _, r := range compartmentOne {
			if strings.ContainsRune(compartmentTwo, r) {
				ans += getPrio(r)
				continue lines
			}
		}
	}
	return ans
}

func p2(lines []string) int {
	ans := 0
lines:
	for i := 0; i <= len(lines)-3; i += 3 {
		elfOne := lines[i]
		elfTwo := lines[i+1]
		elfThree := lines[i+2]
		for _, r := range []rune(elfOne) {
			if strings.ContainsRune(elfTwo, r) && strings.ContainsRune(elfThree, r) {
				ans += getPrio(r)
				continue lines
			}
		}
	}
	return ans
}
