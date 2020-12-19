package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()

	rules, messages := Parse(in)
	matcher := regexp.MustCompile(GenerateRegexpString(rules, 0, true, false))
	fmt.Println("Part 1:", CountMatching(matcher, messages))

	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"
	matcherp2String := GenerateRegexpString(rules, 0, true, true)
	matcherp2 := regexp.MustCompile(matcherp2String)
	fmt.Println("Part 2:", CountMatching(matcherp2, messages))
}

func CountMatching(r *regexp.Regexp, messages []string) (totalMatched int) {
	for _, message := range messages {
		if r.MatchString(message) {
			totalMatched++
		}
	}
	return
}

func GenerateRegexpUpToNBeforeAndAfter(before string, after string, n int) string {
	possibilities := []string{}
	for i := 0; i < n; i++ {
		possibilities = append(possibilities, fmt.Sprintf("%v{%v}%v{%v}", before, i+1, after, i+1))
	}
	return "(" + strings.Join(possibilities, "|") + ")"
}

func GenerateRegexpString(rulesMap map[int]string, ruleNumber int, includeStartAndEnd bool, part2 bool) string {
	currentRule := rulesMap[ruleNumber]

	if currentRule == "a" || currentRule == "b" {
		return currentRule
	}

	matchingString := []string{}

	if part2 && ruleNumber == 8 {
		matchingString = append(matchingString, GenerateRegexpString(rulesMap, 42, false, part2)+"+")
	} else if part2 && ruleNumber == 11 {
		before := GenerateRegexpString(rulesMap, 42, false, part2)
		after := GenerateRegexpString(rulesMap, 31, false, part2)

		matchingString = append(matchingString, GenerateRegexpUpToNBeforeAndAfter(before, after, 5))
	} else {
		for _, part := range strings.Split(currentRule, " | ") {
			stringPart := ""
			for _, loopedRuleNumber := range StringsToInts(strings.Split(part, " ")) {
				stringPart += GenerateRegexpString(rulesMap, loopedRuleNumber, false, part2)
			}

			matchingString = append(matchingString, stringPart)

		}
	}

	result := "^(" + strings.Join(matchingString, "|") + ")$"
	if includeStartAndEnd {
		return result
	} else {
		return strings.Trim(result, "$^")
	}

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

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		if str == "" {
			continue
		}
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
