package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var getValues = regexp.MustCompile(`(\d+)] = (\d+)`)

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1")
	fmt.Println(Solve(in, false))
	fmt.Println("Part 2")
	fmt.Println(Solve(in, true))
}

func Solve(in []string, part2 bool) int64 {
	mask := map[int]string{}
	registerToValue := map[uint64]int64{}
	for _, line := range in {
		if strings.Contains(line, "mask") {
			mask = map[int]string{}
			maskStr := strings.Split(line, " = ")[1]
			for idx, character := range maskStr {
				switch character {
				case 'X':
					if part2 {
						mask[idx] = "X"
					}
					continue
				case '0':
					if !part2 {
						mask[idx] = "0"
					}
				case '1':
					mask[idx] = "1"
				}
			}
		} else {
			values := getValues.FindStringSubmatch(line)[1:]
			rawRegister, _ := strconv.ParseUint(values[0], 10, 64)
			rawValue, _ := strconv.ParseInt(values[1], 10, 64)
			rawBinary := strconv.FormatInt(rawValue, 2)
			if part2 {
				rawBinary = strconv.FormatUint(rawRegister, 2)
			}

			// copy mask in
			thisValue := map[int]string{}
			for idx, val := range mask {
				thisValue[idx] = val
			}

			// go over binary
			for idx, val := range rawBinary {
				realIdx := idx + 36 - len(rawBinary)
				if _, exists := thisValue[realIdx]; !exists {
					thisValue[realIdx] = string(val)
				}
			}

			result := ""
			for i := 0; i < 36; i++ {
				setVal, exists := thisValue[i]
				if exists {
					result += setVal
				} else {
					result += "0"
				}
			}

			if !part2 {
				decValue, _ := strconv.ParseInt(result, 2, 64)
				registerToValue[rawRegister] = decValue

			} else {
				nFloating := strings.Count(result, "X")
				combinations := math.Pow(2, float64(nFloating))
				for i := 0; i < int(combinations); i++ {
					inBinary := fmt.Sprintf("%0*v", nFloating, strconv.FormatInt(int64(i), 2))

					register := result
					for _, val := range inBinary {
						register = strings.Replace(register, "X", string(val), 1)
					}
					n, _ := strconv.ParseUint(register, 2, 64)

					registerToValue[n] = rawValue
				}
			}
		}
	}
	total := int64(0)
	for _, val := range registerToValue {
		total += val
	}
	return total
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
