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
	fmt.Println(p2(string(input)))
}

type forest map[int]tree

type tree struct {
	size       int
	onEdge     bool
	neighbours [4]int
}

func (t tree) isVisible(direction int, asking int, f forest) bool {
	if asking <= t.size {
		return false
	}
	if t.onEdge {
		return true
	}

	visible := false
	visible = f[t.neighbours[direction]].isVisible(direction, asking, f)
	return visible
}

func p1(lines string) int {
	ans := 0
	firstLine := strings.Split(lines, "\n")
	width := len(firstLine[0])
	lines = strings.ReplaceAll(lines, "\n", "")

	fr := make(forest)

	for i, _ := range lines {
		s, _ := strconv.Atoi(lines[i : i+1])
		edgy := false
		if i%width == 0 {
			edgy = true
		} else if (i+1)%(width) == 0 {
			edgy = true
		} else if i < width {
			edgy = true
		} else if i > len(lines)-width {
			edgy = true
		}

		tr := tree{
			size:       s,
			onEdge:     edgy,
			neighbours: [4]int{i - 1, i - width, i + 1, i + width},
		}
		fr[i] = tr
	}
	for _, v := range fr {
		if v.onEdge {
			ans++
			continue
		}
		vis := false
		for i, j := range v.neighbours {
			if fr[j].isVisible(i, v.size, fr) {
				vis = true
				break
			}
		}
		if vis {
			ans++
		}

	}

	return ans
}

func (t tree) isVisibleView(direction int, asking int, f forest) int {
	if asking <= t.size {
		return 1
	}
	if t.onEdge {
		return 1
	}
	visible := f[t.neighbours[direction]].isVisibleView(direction, asking, f)

	return visible + 1
}

func p2(lines string) int {
	ans := 0
	firstLine := strings.Split(lines, "\n")
	width := len(firstLine[0])
	lines = strings.ReplaceAll(lines, "\n", "")

	fr := make(forest)

	for i, _ := range lines {
		s, _ := strconv.Atoi(lines[i : i+1])
		edgy := false
		if i%width == 0 {
			edgy = true
		} else if (i+1)%(width) == 0 {
			edgy = true
		} else if i < width {
			edgy = true
		} else if i > len(lines)-width {
			edgy = true
		}

		tr := tree{
			size:       s,
			onEdge:     edgy,
			neighbours: [4]int{i - 1, i - width, i + 1, i + width},
		}
		fr[i] = tr
	}
	for _, v := range fr {
		if v.onEdge {
			continue
		}
		view := 0
		var views [4]int
		for i, j := range v.neighbours {
			views[i] = fr[j].isVisibleView(i, v.size, fr)
		}
		view = views[0] * views[1] * views[2] * views[3]

		if view > ans {
			ans = view
		}

	}

	return ans
}
