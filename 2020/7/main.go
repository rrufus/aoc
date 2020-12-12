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

type Rule struct {
	Colour string
	Number int
}

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1")
	parentBags := map[string][]string{}
	childBags := map[string][]*Rule{}
	for _, ruleString := range in {
		bagColour, containsBags := ParseRule(ruleString)

		childBags[bagColour] = containsBags

		for _, bag := range containsBags {
			if parentBags[bag.Colour] == nil {
				parentBags[bag.Colour] = []string{bagColour}
			} else {
				parentBags[bag.Colour] = append(parentBags[bag.Colour], bagColour)
			}
		}
	}

	fmt.Println(len(FindParentBags(parentBags, "shiny gold", map[string]bool{})))

	fmt.Println("Part 2")
	fmt.Println(FindChildBags(childBags, "shiny gold"))

}

func FindParentBags(parentBags map[string][]string, colour string, counted map[string]bool) map[string]bool {
	directColourParents := parentBags[colour]

	for _, parentColour := range directColourParents {
		counted[parentColour] = true
		FindParentBags(parentBags, parentColour, counted)
	}

	return counted
}

func FindChildBags(childBags map[string][]*Rule, colour string) int {
	total := 0
	rules, exists := childBags[colour]
	if exists {
		for _, rule := range rules {
			total += rule.Number + rule.Number*FindChildBags(childBags, rule.Colour)
		}
	}
	return total
}

func ParseRule(rule string) (string, []*Rule) {
	splitOnBagsContain := strings.Split(rule, "bags contain")
	colour := strings.TrimSpace(splitOnBagsContain[0])
	bagsStr := strings.Split(strings.Trim(splitOnBagsContain[1], " ."), ",")
	rules := []*Rule{}
	if len(bagsStr) == 1 && bagsStr[0] == "no other bags" {
		return colour, rules
	}

	for _, bag := range bagsStr {
		words := strings.Split(strings.TrimSpace(bag), " ")
		number, _ := strconv.Atoi(words[0])
		bagColour := strings.Join(words[1:3], " ")

		rules = append(rules, &Rule{Number: number, Colour: bagColour})
	}

	return colour, rules
}
