package main

import (
	"embed"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//go:embed *
var f embed.FS

func main() {
	input, _ := f.ReadFile("input")
	lines := strings.Split(string(input), "\n")
	fmt.Println(p1(lines))
}

func p1(lines []string) int {
	ans := 0
	folders := make(map[string]int) // name + size
	openFolders := make([]string, 0, 0)
	for i, k := range lines {
		if strings.HasPrefix(k, "$ cd") {
			command := strings.Split(k, " ")[2]
			switch command {
			case "/":
				// Close all but top folder
				openFolders = []string{"/"}
			case "..":
				// Close recently opened folder
				openFolders = openFolders[:len(openFolders)-1]
			default:
				r := fmt.Sprintf("%s%d", command, rand.Int31())
				folders[r] = 0
				openFolders = append(openFolders, r)
			}
		}
		if k == "$ ls" {
			for j := i + 1; j < len(lines); j++ {
				if strings.HasPrefix(lines[j], "$") {
					break
				}
				if !strings.HasPrefix(lines[j], "dir") {
					atoi, _ := strconv.Atoi(strings.Split(lines[j], " ")[0])
					for _, f := range openFolders {
						folders[f] = folders[f] + atoi
					}
				}
			}
		}
	}

	// just did small tweaks here to deal with part 2

	currentAvailable := 70000000 - folders["/"]
	missing := 30000000 - currentAvailable
	smallest := 70000000
	for _, v := range folders {
		if v > missing && v < smallest {
			smallest = v
		}
	}
	fmt.Println(smallest)
	return ans
}
