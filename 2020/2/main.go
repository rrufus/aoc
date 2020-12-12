package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
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
	in := ReadFromInput()

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
