package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	matchingStringsPart1 []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	matchingStringsPart2 []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1")
	values := []string{}
	for _, line := range in {
		var firstRune rune = 'u'
		var lastRune rune = 'u'
		var currentRune rune = '0'
		for _, r := range line {
			if r == '0' || r == '1' || r == '2' || r == '3' || r == '4' || r == '5' || r == '6' || r == '7' || r == '8' || r == '9' {
				currentRune = r
				if firstRune == 'u' {
					firstRune = r
				}
			}
		}
		lastRune = currentRune

		value := fmt.Sprintf("%s%s", string(firstRune), string(lastRune))

		values = append(values, value)
	}
	ints := StringsToInts(values)
	total := 0
	for _, entry := range ints {
		total += entry
	}
	fmt.Println(total)

	fmt.Println("Part 2")

	values2 := []string{}
	for _, line := range in {
		lowestIndex := len(line)
		lowestFoundIndexValue := ""
		highestIndex := -1
		highestFoundIndexValue := ""
		for _, matching := range matchingStringsPart2 {
			foundFirstIndex := strings.Index(line, matching)
			foundLastIndex := strings.LastIndex(line, matching)
			if foundFirstIndex != -1 && foundFirstIndex < lowestIndex {
				lowestIndex = foundFirstIndex
				if len(matching) != 1 {
					lowestFoundIndexValue = NumberedStringToNumber(matching)
				} else {
					lowestFoundIndexValue = matching
				}
			}
			if foundLastIndex != -1 && foundLastIndex > highestIndex {
				highestIndex = foundLastIndex
				if len(matching) != 1 {
					highestFoundIndexValue = NumberedStringToNumber(matching)
				} else {
					highestFoundIndexValue = matching
				}
			}
		}

		value := fmt.Sprintf("%s%s", lowestFoundIndexValue, highestFoundIndexValue)
		values2 = append(values2, value)
	}
	ints2 := StringsToInts(values2)
	total2 := 0
	for _, entry2 := range ints2 {
		total2 += entry2
	}
	fmt.Println(total2)

}

func NumberedStringToNumber(input string) string {
	switch input {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "zero":
		return "0"
	}
	return ""
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
