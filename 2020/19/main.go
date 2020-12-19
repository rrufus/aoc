package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	rules, messages := Parse(ReadFromInput())

	matcher := regexp.MustCompile(GenerateRegexpString(rules, 0, true, false))
	matcherp2 := regexp.MustCompile(GenerateRegexpString(rules, 0, true, true))

	fmt.Println("Part 1:", CountMatching(matcher, messages))
	fmt.Println("Part 2:", CountMatching(matcherp2, messages))
}

func Parse(in []string) (map[int]string, []string) {
	ruleToStr := map[int]string{}
	for idx, line := range in {
		if line == "" {
			return ruleToStr, in[idx:]
		}
		split := strings.Split(line, ": ")
		n, _ := strconv.Atoi(split[0])
		ruleToStr[n] = strings.ReplaceAll(split[1], "\"", "")
	}
	return ruleToStr, []string{}
}

func GenerateRegexpString(rulesMap map[int]string, ruleNumber int, includeStartAndEnd bool, part2 bool) string {
	currentRule := rulesMap[ruleNumber]

	if currentRule == "a" || currentRule == "b" {
		return currentRule
	}

	matchingStringParts := []string{}

	if part2 && ruleNumber == 8 {
		matchingStringParts = append(matchingStringParts, GenerateRegexpString(rulesMap, 42, false, part2)+"+")
	} else if part2 && ruleNumber == 11 {
		before := GenerateRegexpString(rulesMap, 42, false, part2)
		after := GenerateRegexpString(rulesMap, 31, false, part2)

		matchingStringParts = append(matchingStringParts, GenerateRegexpUpToNBeforeAndAfter(before, after, 5))
	} else {
		for _, part := range strings.Split(currentRule, " | ") {
			stringPart := ""
			for _, loopedRuleNumber := range StringsToInts(strings.Split(part, " ")) {
				stringPart += GenerateRegexpString(rulesMap, loopedRuleNumber, false, part2)
			}

			matchingStringParts = append(matchingStringParts, stringPart)
		}
	}

	result := "(" + strings.Join(matchingStringParts, "|") + ")"
	if includeStartAndEnd {
		return "^" + result + "$"
	} else {
		return result
	}
}

func GenerateRegexpUpToNBeforeAndAfter(before string, after string, n int) string {
	possibilities := []string{}
	for i := 0; i < n; i++ {
		possibilities = append(possibilities, fmt.Sprintf("%v{%v}%v{%v}", before, i+1, after, i+1))
	}
	return "(" + strings.Join(possibilities, "|") + ")"
}

func CountMatching(r *regexp.Regexp, messages []string) (totalMatched int) {
	for _, message := range messages {
		if r.MatchString(message) {
			totalMatched++
		}
	}
	return
}

func StringsToInts(stringInputs []string) []int {
	ints := []int{}
	for _, str := range stringInputs {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}
	return ints
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
