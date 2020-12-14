package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var getValues = regexp.MustCompile(`(\d+)] = (\d+)`)

func main() {
	in := ReadFromInput()
	// in := ReadFromStdIn()

	// What is the sum of all values left in memory after it completes?
	fmt.Println("Part 1")
	mask := map[int]int{}
	registerToValue := map[int]uint64{}
	for _, line := range in {
		if strings.Contains(line, "mask") {
			mask = map[int]int{}
			maskStr := strings.Split(line, " = ")[1]
			for idx, character := range maskStr {
				switch character {
				case 'X':
					continue
				case '0':
					mask[idx] = 0
				case '1':
					mask[idx] = 1
				}
			}
		} else {
			values := getValues.FindStringSubmatch(line)[1:]
			register, _ := strconv.Atoi(values[0])
			rawValue, _ := strconv.ParseInt(values[1], 10, 64)
			rawBinary := strconv.FormatInt(rawValue, 2)

			thisValue := map[int]int{}
			for idx, val := range mask {
				thisValue[idx] = val
			}

			for idx, val := range rawBinary {
				realIdx := idx + 36 - len(rawBinary)
				if _, exists := thisValue[realIdx]; !exists {
					switch val {
					case '0':
						thisValue[realIdx] = 0
					case '1':
						thisValue[realIdx] = 1
					}
				}
			}
			fullBinary := ""
			for i := 0; i < 36; i++ {
				setVal := thisValue[i]
				fullBinary = fmt.Sprintf("%v%v", fullBinary, setVal)
			}
			decValue, _ := strconv.ParseUint(fullBinary, 2, 64)
			registerToValue[register] = decValue
		}
	}
	total := uint64(0)
	for _, val := range registerToValue {
		total += val
	}
	fmt.Println(total)

	fmt.Println("Part 2")

}

// func Parse(in []string) []*Program {

// }

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

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
