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
	fmt.Println(p2(string(input)))
}

type monkey struct {
	items   []int
	inspect func(int) int
	test    func(int) (int, int)
}

func (m *monkey) catch(it int) {
	m.items = append(m.items, it)
}

func (m *monkey) play(others []*monkey) int {
	aItems := len(m.items)
	for i, _ := range m.items {
		inspected := m.inspect(m.items[i])
		next, divided := m.test(inspected)
		others[next].catch(divided)
		fmt.Println(divided)
	}
	//15864280884851703964
	//18446744073709551615
	m.items = []int{}
	return aItems
}

/*
func p1(lines string) int {
	ans := 0

	var monkeys []*monkey

	c := regexp.MustCompile("Monkey \\d:\\n  Starting items: (.*)\\n  Operation: new = old (.*)\\n  Test: divisible by (.*)\\n    If true: throw to monkey (\\d)\\n    If false: throw to monkey (\\d)")
	m := c.FindAllStringSubmatch(lines, -1)
	for _, me := range m {
		var itemsN []int
		items := strings.Split(me[1], ", ")
		for _, i := range items {
			iN, _ := strconv.Atoi(i)
			itemsN = append(itemsN, iN)
		}
		var operation rune
		var target int
		fmt.Sscanf(me[2], "%c %d", &operation, &target)
		op := func(old int) int {
			if target == 0 {
				return old * old
			}
			switch operation {
			case '*':
				return old * target
			case '+':
				return old + target
			}
			return 0
		}

		m4, _ := strconv.Atoi(me[4])
		m5, _ := strconv.Atoi(me[5])
		var divisible int
		fmt.Sscanf(me[3], "%d", &divisible)
		test := func(in int) (int, int) {
			bored := in // For part 1 divide by 3 here
			if int(bored)%divisible == 0 {
				return m4, bored
			} else {
				return m5, bored
			}
		}
		m := monkey{
			items:   itemsN,
			inspect: op,
			test:    test,
		}
		monkeys = append(monkeys, &m)
	}

	inspects := map[int]int{}
	rounds := 20
	for i := 0; i < rounds; i++ {
		for j, v := range monkeys {
			t := v.play(monkeys)
			inspects[j] = inspects[j] + t
		}
	}
	fmt.Println("Returned:")
	fmt.Println(inspects)
	return ans
}
*/

func p2(lines string) int {
	ans := 0
	var monkeys []*monkey
	var divisibleList []int
	var supermodulo int
	c := regexp.MustCompile("Monkey \\d:\\n  Starting items: (.*)\\n  Operation: new = old (.*)\\n  Test: divisible by (.*)\\n    If true: throw to monkey (\\d)\\n    If false: throw to monkey (\\d)")
	m := c.FindAllStringSubmatch(lines, -1)
	for _, me := range m {
		var itemsN []int
		items := strings.Split(me[1], ", ")
		for _, i := range items {
			iN, _ := strconv.Atoi(i)
			itemsN = append(itemsN, iN)
		}
		var operation rune
		var target int
		fmt.Sscanf(me[2], "%c %d", &operation, &target)
		insp := func(old int) int {
			old = old % supermodulo
			if target == 0 {
				return old * old
			}
			switch operation {
			case '*':
				return old * target
			case '+':
				return old + target
			}
			return 0
		}

		m4, _ := strconv.Atoi(me[4])
		m5, _ := strconv.Atoi(me[5])
		var divisible int
		fmt.Sscanf(me[3], "%d", &divisible)
		test := func(in int) (int, int) {

			if in%int(divisible) == 0 {
				return m4, in
			} else {
				return m5, in
			}
		}
		m := monkey{
			items:   itemsN,
			inspect: insp,
			test:    test,
		}
		monkeys = append(monkeys, &m)
		divisibleList = append(divisibleList, divisible)
	}
	supermodulo = 1
	for i := 0; i < len(divisibleList); i++ {
		supermodulo = supermodulo * divisibleList[i]
	}

	inspects := map[int]int{}
	rounds := 10_000
	for i := 0; i < rounds; i++ {
		for j, v := range monkeys {
			t := v.play(monkeys)
			inspects[j] = inspects[j] + t
		}
	}
	fmt.Println("Returned:")
	fmt.Println(inspects)
	/*
		Monkey 0 inspected items 52166 times.
		Monkey 1 inspected items 47830 times.
		Monkey 2 inspected items 1938 times.
		Monkey 3 inspected items 52013 times.
	*/
	fmt.Println("Expected:")
	expected := map[int]int{
		0: 52166,
		1: 47830,
		2: 1938,
		3: 52013,
	}
	fmt.Println(expected)

	return ans
}

// It's decided that if it is divisible by 23, then it gets sent to monkey 2
// monkey 2 multiplies it by itself, and if it the is divisible by 13 send
// it to monkey 3. Who then adds 3
