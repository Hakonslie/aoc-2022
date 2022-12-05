package main

import (
	"embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	input, _ := f.ReadFile("input")
	inputz := strings.ReplaceAll(string(input), "   ", "[ ]")
	lines := strings.Split(inputz, "\n")
	fmt.Println(p2(lines))
}

type stack struct {
	boxes []string
}

func (s *stack) put(r string) {
	s.boxes = append(s.boxes, r)
}
func (s *stack) pop() string {
	var length = len(s.boxes) - 1
	ru := s.boxes[length]
	s.boxes = s.boxes[:length]
	return ru
}

func p1(lines []string) int {
	ans := 0
	var stacks []stack
	for _, l := range lines {
		if strings.Contains(l, "1") {
			i, _ := strconv.Atoi(string([]rune(l)[len(l)-1]))
			for j := 0; j < i; j++ {
				stacks = append(stacks, stack{})
			}
			break
		}
	}

	for _, l := range lines {
		if !strings.HasPrefix(l, "[") {
			break
		}
		l = l + " "
		onStack := 0

		for i := 0; i <= len(l)-4; i += 4 {
			box := l[i : i+3]
			match, _ := regexp.MatchString("[A-Z]", box)
			if match {
				stacks[onStack].put(box[1:2])
			}
			onStack++
		}
	}

	// flip stacks
	fmt.Println(stacks)
	for i, _ := range stacks {
		var ordered []string
		for j := len(stacks[i].boxes) - 1; j >= 0; j-- {
			ordered = append(ordered, stacks[i].boxes[j])
		}
		stacks[i].boxes = ordered
	}
	fmt.Println(stacks)

	for _, l := range lines {
		if !strings.HasPrefix(l, "move") {
			continue
		}
		re := regexp.MustCompile("\\d+")
		steps := re.FindAllString(l, -1)
		fmt.Println(steps)
		move, _ := strconv.Atoi(steps[0])
		from, _ := strconv.Atoi(steps[1])
		to, _ := strconv.Atoi(steps[2])

		fmt.Printf("moving %d from %d to %d \n", move, from, to)
		for i := 0; i < move; i++ {
			stacks[to-1].put(stacks[from-1].pop())
		}
		fmt.Println(stacks)
	}
	for _, s := range stacks {
		fmt.Println(s.boxes)
	}
	return ans
}

func p2(lines []string) int {
	ans := 0
	var stacks []stack
	for _, l := range lines {
		if strings.Contains(l, "1") {
			i, _ := strconv.Atoi(string([]rune(l)[len(l)-1]))
			for j := 0; j < i; j++ {
				stacks = append(stacks, stack{})
			}
			break
		}
	}

	for _, l := range lines {
		if !strings.HasPrefix(l, "[") {
			break
		}
		l = l + " "
		onStack := 0

		for i := 0; i <= len(l)-4; i += 4 {
			box := l[i : i+3]
			match, _ := regexp.MatchString("[A-Z]", box)
			if match {
				stacks[onStack].put(box[1:2])
			}
			onStack++
		}
	}

	// flip stacks
	for i, _ := range stacks {
		var ordered []string
		for j := len(stacks[i].boxes) - 1; j >= 0; j-- {
			ordered = append(ordered, stacks[i].boxes[j])
		}
		stacks[i].boxes = ordered
	}

	for _, l := range lines {
		if !strings.HasPrefix(l, "move") {
			continue
		}
		re := regexp.MustCompile("\\d+")
		steps := re.FindAllString(l, -1)
		move, _ := strconv.Atoi(steps[0])
		from, _ := strconv.Atoi(steps[1])
		to, _ := strconv.Atoi(steps[2])

		fmt.Println(stacks)
		var holding []string
		for i := 0; i < move; i++ {
			holding = append(holding, stacks[from-1].pop())
		}
		fmt.Println(holding)
		for i := len(holding) - 1; i >= 0; i-- {
			stacks[to-1].put(holding[i])
		}
		fmt.Println(stacks)
	}
	for _, s := range stacks {
		fmt.Println(s.boxes)
	}
	return ans
}
