package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()

	totalp1 := 0
	totalp2 := 0
	for _, problem := range in {
		totalp1 += CalculateResult(problem, false)
		totalp2 += CalculateResult(problem, true)
	}
	fmt.Println("Part 1:", totalp1)
	fmt.Println("Part 2:", totalp2)
}

func CalculateResult(input string, part2 bool) (total int) {

	input = DoAdditions(DoParentheses(input, part2), part2)

	items := strings.Split(input, " ")
	lastItem := ""
	for _, item := range items {
		n, err := strconv.Atoi(item)
		if err != nil {
			lastItem = item
			continue
		}
		if lastItem == "" {
			total = n
		} else if lastItem == "+" {
			total += n
		} else if lastItem == "*" {
			total *= n
		}
	}
	return
}

func DoParentheses(problem string, part2 bool) string {
	pointerIdx := 0
	parenthesesDepth := 0
	start := -1
	result := problem

	for pointerIdx < len(problem) {
		val := problem[pointerIdx]
		if val == ' ' || val == '*' || val == '+' || (48 <= val && val <= 57) {
			pointerIdx++
			continue
		}
		if val == '(' {
			if start == -1 {
				start = pointerIdx + 1
			}
			parenthesesDepth++
			pointerIdx++
			continue
		}
		if val == ')' {
			parenthesesDepth--
			if parenthesesDepth == 0 {
				subProblem := problem[start:pointerIdx]
				if strings.ContainsAny(subProblem, "()") {
					reducedSubProblem := DoParentheses(subProblem, part2)
					result = strings.Replace(result, fmt.Sprintf("(%v)", subProblem), fmt.Sprint(CalculateResult(reducedSubProblem, part2)), -1)
				} else {
					result = strings.Replace(result, fmt.Sprintf("(%v)", subProblem), fmt.Sprint(CalculateResult(subProblem, part2)), -1)
				}

				start = -1
			}
			pointerIdx++
			continue
		}
	}

	return result
}

func DoAdditions(input string, part2 bool) string {
	if !part2 {
		return input
	}
	for strings.Count(input, "+") > 0 {
		items := strings.Split(input, " ")
		for idx, item := range items {
			if item == "+" {
				before := items[idx-1]
				after := items[idx+1]
				a, _ := strconv.Atoi(before)
				b, _ := strconv.Atoi(after)

				input = strings.Replace(input, fmt.Sprintf("%v + %v", before, after), fmt.Sprint(a+b), 1)

				return DoAdditions(input, part2)
			}
		}
	}
	return input
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
