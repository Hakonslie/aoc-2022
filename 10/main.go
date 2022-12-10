package main

import (
	"embed"
	"fmt"
	"math"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	input, _ := f.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	fmt.Println(p1(lines))
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
	return p.x * p.inCycle
}

func p1(lines []string) int {
	ans := 0
	pr := program{inCycle: 0, x: 1}
	for _, l := range lines {
		var instruction string
		var x int
		fmt.Sscanf(l, "%s %d", &instruction, &x)
		switch instruction {
		case "noop":
			pr.startCycle()
			if (pr.inCycle+20)%40 == 0 {
				ans += pr.strength()
			}
			pr.endCycle(nil)
		case "addx":
			pr.startCycle()
			if (pr.inCycle+20)%40 == 0 {
				ans += pr.strength()
			}
			pr.endCycle(nil)
			pr.startCycle()
			if (pr.inCycle+20)%40 == 0 {
				ans += pr.strength()
			}
			pr.endCycle(&x)
		}
	}

	return ans
}

type crt struct {
	monitor string
}

func (c *crt) tryDraw(cycle int, x int) {
	pointer := cycle - 1
	if math.Abs(float64(x-(pointer%40))) <= 1 {
		c.monitor += "#"
	} else {
		c.monitor += "."
	}
}

func p2(lines []string) {
	pr := program{inCycle: 0, x: 1}
	crt := crt{}

	for _, l := range lines {
		var instruction string
		var x int
		fmt.Sscanf(l, "%s %d", &instruction, &x)
		switch instruction {
		case "noop":
			pr.startCycle()
			crt.tryDraw(pr.inCycle, pr.x)
			pr.endCycle(nil)
		case "addx":
			pr.startCycle()
			crt.tryDraw(pr.inCycle, pr.x)
			pr.endCycle(nil)
			pr.startCycle()
			crt.tryDraw(pr.inCycle, pr.x)
			pr.endCycle(&x)
		}
	}
	for m := 0; m <= 220; m += 40 {
		fmt.Println(crt.monitor[m : m+40])
	}
}
