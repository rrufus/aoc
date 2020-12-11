package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFromStdIn() []string {
	lines := []string{}
	reader := bufio.NewReader(os.Stdin)

read_loop:
	for {
		text, _ := reader.ReadString('\n')
		if text == "\n" {
			break read_loop
		}
		lines = append(lines, strings.TrimSpace(text))
	}

	return lines
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}

	return ints
}

type RuleAndPassword struct {
	Password  string
	Character string
	Lower     int
	Upper     int
}

func GetRuleAndPassword(line string) *RuleAndPassword {
	r := &RuleAndPassword{}

	splitBySemicolon := strings.Split(line, ":")
	r.Password = strings.TrimSpace(splitBySemicolon[1])

	splitBySpace := strings.Split(splitBySemicolon[0], " ")
	r.Character = splitBySpace[1]

	splitByHyphen := strings.Split(splitBySpace[0], "-")

	r.Lower, _ = strconv.Atoi(splitByHyphen[0])
	r.Upper, _ = strconv.Atoi(splitByHyphen[1])

	return r
}

func main() {
	in := ReadFromStdIn()

	fmt.Println("part 1")
	number := 0
	for _, line := range in {
		item := GetRuleAndPassword(line)

		nTimesCharacterPresent := strings.Count(item.Password, item.Character)

		if item.Lower <= nTimesCharacterPresent && nTimesCharacterPresent <= item.Upper {
			number++
		}
	}
	fmt.Println(number)

	fmt.Println("part 2")
	newNumber := 0
	for _, line := range in {
		item := GetRuleAndPassword(line)

		matches := 0

		if string(item.Password[item.Lower-1]) == item.Character {
			matches++
		}
		if string(item.Password[item.Upper-1]) == item.Character {
			matches++
		}
		if matches == 1 {
			newNumber++
		}
	}
	fmt.Println(newNumber)
}
