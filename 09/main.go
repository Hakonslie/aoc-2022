package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed *
var f embed.FS

type head struct {
	knotPositions [][2]int // x,y
	headPosition  [2]int
	tailPosition  [2]int
	visited       map[string]any
}

func coordToString(c [2]int) string {
	return fmt.Sprintf("%d,%d", c[0], c[1])
}

func (h *head) move(direction int) {
	switch direction {
	// left
	case 0:
		h.headPosition[0] = h.headPosition[0] - 1
	// up
	case 1:
		h.headPosition[1] = h.headPosition[1] - 1
	// right
	case 2:
		h.headPosition[0] = h.headPosition[0] + 1
	// down
	case 3:
		h.headPosition[1] = h.headPosition[1] + 1
	}

	// Potential Diagonal Movement
	if (h.tailPosition[0] < h.headPosition[0] && h.tailPosition[1] < h.headPosition[1]-1) || (h.tailPosition[0] < h.headPosition[0]-1 && h.tailPosition[1] < h.headPosition[1]) {
		// move downright
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[0] = h.tailPosition[0] + 1
		h.tailPosition[1] = h.tailPosition[1] + 1
	}
	if (h.tailPosition[0] > h.headPosition[0]+1 && h.tailPosition[1] < h.headPosition[1]) || (h.tailPosition[0] > h.headPosition[0] && h.tailPosition[1] < h.headPosition[1]-1) {
		// move downleft
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[0] = h.tailPosition[0] - 1
		h.tailPosition[1] = h.tailPosition[1] + 1
	}

	if (h.tailPosition[0] < h.headPosition[0] && h.tailPosition[1] > h.headPosition[1]+1) || (h.tailPosition[0] < h.headPosition[0]-1 && h.tailPosition[1] > h.headPosition[1]) {
		// move upright
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[0] = h.tailPosition[0] + 1
		h.tailPosition[1] = h.tailPosition[1] - 1
	}

	if (h.tailPosition[0] > h.headPosition[0] && h.tailPosition[1] > h.headPosition[1]+1) || (h.tailPosition[0] > h.headPosition[0]+1 && h.tailPosition[1] > h.headPosition[1]) {
		// move upleft
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[0] = h.tailPosition[0] - 1
		h.tailPosition[1] = h.tailPosition[1] - 1
	}

	// vertical/horizontal movement
	if h.tailPosition[0] < h.headPosition[0]-1 {
		// move right
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[0] = h.tailPosition[0] + 1
	}
	if h.tailPosition[0] > h.headPosition[0]+1 {
		// move left
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[0] = h.tailPosition[0] - 1
	}
	if h.tailPosition[1] > h.headPosition[1]+1 {
		// move up
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[1] = h.tailPosition[1] - 1
	}
	if h.tailPosition[1] < h.headPosition[1]-1 {
		// move down
		h.visited[coordToString(h.tailPosition)] = "ok"
		h.tailPosition[1] = h.tailPosition[1] + 1
	}
}

func main() {
	input, _ := f.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	fmt.Println(p1(lines))
}

func p1(lines []string) int {
	ans := 0
	h := head{
		headPosition: [2]int{},
		tailPosition: [2]int{},
		visited:      make(map[string]interface{}),
	}
	for _, l := range lines {
		spl := strings.Split(l, " ")
		d := strings.Index("LURD", spl[0])
		speed, _ := strconv.Atoi(spl[1])
		for i := 0; i < speed; i++ {
			h.move(d)
		}
	}

	h.visited[coordToString(h.tailPosition)] = "ok"
	ans = len(h.visited)

	return ans
}

/*
func p2(lines []string) int {
	ans := 0

	return ans
}
*/
