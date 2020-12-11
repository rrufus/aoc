package main

import (
	"bufio"
	"fmt"
	"os"
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

func CalculateTrees(lines []string, right int, down int) int {
	horizontal := 0
	trees := 0

	for i, line := range lines {
		if i == 0 || i%down != 0 {
			continue
		}

		horizontal += right
		horizontal = horizontal % len(line)

		if line[horizontal] == '#' {
			trees++
		}
	}
	return trees
}

func main() {
	in := ReadFromStdIn()

	fmt.Println("Part 1")
	horizontal := 0
	trees := 0
	for i, line := range in {
		if i == 0 {
			continue
		}

		horizontal += 3
		horizontal = horizontal % len(line)

		if line[horizontal] == '#' {
			trees++
		}
	}
	fmt.Println(trees)

	fmt.Println("Part 2")
	fmt.Println(CalculateTrees(in, 1, 1) * CalculateTrees(in, 3, 1) * CalculateTrees(in, 5, 1) * CalculateTrees(in, 7, 1) * CalculateTrees(in, 1, 2))
}
