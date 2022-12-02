package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed *
var f embed.FS

// 1 for rock, 2 for paper, 3 for scissors
// 6 for win
// 3 for draw
// 0 for loss

type outcome int

const (
	win  outcome = 6
	draw         = 3
	loss         = 0
)

type symbol string

const (
	rock     symbol = "A"
	paper    symbol = "B"
	scissors symbol = "C"
)

var expectedOutcome = map[string]outcome{
	"X": loss,
	"Y": draw,
	"Z": win,
}

func (s symbol) battle(opponent symbol) outcome {
	if opponent == s {
		return draw
	}
	switch s {
	case rock:
		switch opponent {
		case scissors:
			return win
		case paper:
			return loss
		}
	case paper:
		switch opponent {
		case rock:
			return win
		case scissors:
			return loss
		}
	case scissors:
		switch opponent {
		case rock:
			return loss
		case paper:
			return win
		}
	}
	return draw
}
func (s symbol) worth() int {
	switch s {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}
	return 0
}

func cheat(wantedOutcome outcome, opponent symbol) symbol {
	for _, s := range []symbol{rock, paper, scissors} {
		if s.battle(opponent) == wantedOutcome {
			return s
		}
	}
	return symbol("")
}

func partOne() {
	data, _ := f.ReadFile("input")
	dataString := strings.ReplaceAll(string(data), "X", "A")
	dataString = strings.ReplaceAll(dataString, "Y", "B")
	dataString = strings.ReplaceAll(dataString, "Z", "C")

	points := 0
	for _, game := range strings.Split(dataString, "\n") {
		players := strings.Split(game, " ")
		points += int(symbol(players[1]).battle(symbol(players[0])))
		points += symbol(players[1]).worth()
	}
	fmt.Println(points)
}

func partTwo() {
	data, _ := f.ReadFile("input")

	points := 0
	for _, game := range strings.Split(string(data), "\n") {
		tabs := strings.Split(game, " ")
		strategy := cheat(expectedOutcome[tabs[1]], symbol(tabs[0]))
		points += int(strategy.battle(symbol(tabs[0])))
		points += symbol(strategy).worth()
	}
	fmt.Println(points)
}
func main() {
	partOne()
	fmt.Println("------")
	partTwo()
}
