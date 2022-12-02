package main

import (
	"embed"
	"fmt"
)

//go:embed *
var f embed.FS

func main() {
	input, _ := f.ReadFile("input")
	fmt.Println(p1)
}

func p1(string input) {

}

func p2(string input) {

}
