package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()
	// in := ReadFromStdIn()

	rules, messages := Parse(in)
	matcher := regexp.MustCompile(GenerateRegexpString(rules, 0))
	totalMatched := 0
	for _, message := range messages {
		if matcher.MatchString(message) {
			totalMatched++
		}
	}
	fmt.Println("Part 1:", totalMatched)
}

func GenerateRegexpString(rulesMap map[int]string, ruleNumber int) string {
	currentRule := rulesMap[ruleNumber]

	if currentRule == "a" || currentRule == "b" {
		return currentRule
	}

	matchingString := []string{}

	for _, part := range strings.Split(currentRule, " | ") {
		stringPart := ""
		for _, ruleNumber := range StringsToInts(strings.Split(part, " ")) {
			stringPart += strings.Trim(GenerateRegexpString(rulesMap, ruleNumber), "$^")
		}
		matchingString = append(matchingString, stringPart)
	}

	return "^(" + strings.Join(matchingString, "|") + ")$"

}

func Parse(in []string) (map[int]string, []string) {
	ruleToStr := map[int]string{}
	messages := []string{}
	isParsingRules := true
	for _, line := range in {
		if line == "" {
			isParsingRules = false
			continue
		}
		if isParsingRules {
			split := strings.Split(line, ": ")
			n, _ := strconv.Atoi(split[0])
			ruleToStr[n] = strings.ReplaceAll(split[1], "\"", "")
		} else {
			messages = append(messages, line)
		}
	}
	return ruleToStr, messages
}

func ReadFromStdIn() []string {
	lines := []string{}
	reader := bufio.NewReader(os.Stdin)

read_loop:
	for {
		text, _ := reader.ReadString('\n')
		if text == "go\n" {
			break read_loop
		}
		lines = append(lines, strings.TrimSpace(text))
	}

	return lines
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
