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
	headPosition [2]int
	tailPosition [2]int
	visited      map[string]any
}

func coordToString(c [2]int) string {
	return fmt.Sprintf("%d,%d", c[0], c[1])
}
func newPosition(firstKnot [2]int, secondKnot [2]int) (bool, [2]int) {
	didMove := false
	newPos := secondKnot
	if (secondKnot[0] < firstKnot[0] && secondKnot[1] < firstKnot[1]-1) || (secondKnot[0] < firstKnot[0]-1 && secondKnot[1] < firstKnot[1]) {
		// move downright
		newPos[0] = secondKnot[0] + 1
		newPos[1] = secondKnot[1] + 1
		didMove = true
	}
	if (secondKnot[0] > firstKnot[0]+1 && secondKnot[1] < firstKnot[1]) || (secondKnot[0] > firstKnot[0] && secondKnot[1] < firstKnot[1]-1) {
		// move downleft
		newPos[0] = secondKnot[0] - 1
		newPos[1] = secondKnot[1] + 1
		didMove = true
	}

	if (secondKnot[0] < firstKnot[0] && secondKnot[1] > firstKnot[1]+1) || (secondKnot[0] < firstKnot[0]-1 && secondKnot[1] > firstKnot[1]) {
		// move upright
		newPos[0] = secondKnot[0] + 1
		newPos[1] = secondKnot[1] - 1
		didMove = true
	}

	if (secondKnot[0] > firstKnot[0] && secondKnot[1] > firstKnot[1]+1) || (secondKnot[0] > firstKnot[0]+1 && secondKnot[1] > firstKnot[1]) {
		// move upleft
		newPos[0] = secondKnot[0] - 1
		newPos[1] = secondKnot[1] - 1
		didMove = true
	}

	// vertical/horizontal movement
	if secondKnot[0] < firstKnot[0]-1 {
		// move right
		newPos[0] = secondKnot[0] + 1
		didMove = true
	}
	if secondKnot[0] > firstKnot[0]+1 {
		// move left
		newPos[0] = secondKnot[0] - 1
		didMove = true
	}
	if secondKnot[1] > firstKnot[1]+1 {
		// move up
		newPos[1] = secondKnot[1] - 1
		didMove = true
	}
	if secondKnot[1] < firstKnot[1]-1 {
		// move down
		newPos[1] = secondKnot[1] + 1
		didMove = true
	}
	return didMove, newPos
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
	didMove, newPos := newPosition(h.headPosition, h.tailPosition)
	if didMove {
		h.visited[coordToString(h.tailPosition)] = "ok"
	}
	h.tailPosition = newPos
}

func main() {
	input, _ := f.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	fmt.Println(p2(lines))
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

type rope struct {
	knotPositions [][2]int
	visited       map[string]any
}

func (r *rope) move(d int) {
	switch d {
	// left
	case 0:
		r.knotPositions[0][0] = r.knotPositions[0][0] - 1
	// up
	case 1:
		r.knotPositions[0][1] = r.knotPositions[0][1] - 1
	// right
	case 2:
		r.knotPositions[0][0] = r.knotPositions[0][0] + 1
	// down
	case 3:
		r.knotPositions[0][1] = r.knotPositions[0][1] + 1
	}
	for i := 1; i < len(r.knotPositions); i++ {
		didMove, newPos := newPosition(r.knotPositions[i-1], r.knotPositions[i])
		if didMove {
			if i == len(r.knotPositions)-1 {
				r.visited[coordToString(r.knotPositions[i])] = "ok"
			}
			r.knotPositions[i] = newPos
		}
	}
}

func newRope(knots int) rope {
	r := rope{visited: make(map[string]any)}
	var knotSlice [][2]int
	for i := 0; i < knots; i++ {
		knotSlice = append(knotSlice, [2]int{0, 0})
	}
	r.knotPositions = knotSlice
	return r
}

func p2(lines []string) int {
	ans := 0
	r := newRope(10)

	for _, l := range lines {
		spl := strings.Split(l, " ")
		d := strings.Index("LURD", spl[0])
		speed, _ := strconv.Atoi(spl[1])
		for i := 0; i < speed; i++ {
			r.move(d)
		}
	}

	r.visited[coordToString(r.knotPositions[9])] = "ok"
	ans = len(r.visited)

	return ans
}
