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

type program struct {
	inCycle int
	x       int
}

func (p *program) startCycle() {
	p.inCycle++
}
func (p *program) endCycle(i *int) {
	if i != nil {
		p.x += *i
	}

}
func (p *program) strength() int {
	fmt.Printf("x: %d, cycle: %d \n", p.x, p.inCycle)
	return p.x * p.inCycle
}

func p1(lines []string) int {
	ans := 0
	pr := program{inCycle: 0, x: 1}
	stopat := make(map[int]interface{})
	stopat[20] = ""
	stopat[60] = ""
	stopat[100] = ""
	stopat[140] = ""
	stopat[180] = ""
	stopat[220] = ""
	for _, l := range lines {
		var instruction string
		var x int
		fmt.Sscanf(l, "%s %d", &instruction, &x)
		switch instruction {
		case "noop":
			pr.startCycle()
			if _, ok := stopat[pr.inCycle]; ok {
				ans += pr.strength()
			}
			pr.endCycle(nil)
		case "addx":
			pr.startCycle()
			if _, ok := stopat[pr.inCycle]; ok {
				ans += pr.strength()
			}
			pr.endCycle(nil)
			pr.startCycle()
			if _, ok := stopat[pr.inCycle]; ok {
				ans += pr.strength()
			}
			pr.endCycle(&x)
		}
	}

	return ans
}

func p2(lines []string) int {
	ans := 0

	return ans
}
