package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in := ReadFromInput()

	fmt.Println("Part 1")
	charactersInPositions := CountCharactersInPositions(in)
	mostCommon := SelectMostCommon(charactersInPositions)
	leastCommon := SelectLeastCommon(charactersInPositions)
	gammaRate, _ := strconv.ParseInt(mostCommon, 2, 64)
	epsilonRate, _ := strconv.ParseInt(leastCommon, 2, 64)
	fmt.Println(gammaRate * epsilonRate)

	fmt.Println("Part 2")
	oxygenRating := ""
	co2Rating := ""
	co2RatingFound := false
	for idx := range in[0] {

		oxygenRatingCandidates := SelectStringsWithPrefix(in, oxygenRating)
		co2RatingCandidates := SelectStringsWithPrefix(in, co2Rating)

		if len(co2RatingCandidates) == 1 {
			co2Rating = co2RatingCandidates[0]
			co2RatingFound = true
		}

		oCharacters := CountCharactersInPositions(oxygenRatingCandidates)
		oZeroCount := oCharacters[idx]["0"]
		oOneCount := oCharacters[idx]["1"]

		if oZeroCount == oOneCount || oOneCount > oZeroCount {
			oxygenRating += "1"
		} else {
			oxygenRating += "0"
		}

		if !co2RatingFound {
			cCharacters := CountCharactersInPositions(co2RatingCandidates)
			cZeroCount := cCharacters[idx]["0"]
			cOneCount := cCharacters[idx]["1"]

			if cZeroCount == cOneCount || cZeroCount < cOneCount {
				co2Rating += "0"
			} else {
				co2Rating += "1"
			}
		}
	}
	co2RatingDec, _ := strconv.ParseInt(co2Rating, 2, 64)
	o2RatingDec, _ := strconv.ParseInt(oxygenRating, 2, 64)
	fmt.Println(co2RatingDec * o2RatingDec)

}

func SelectMostCommon(counts []map[string]int) string {
	result := ""
	for _, count := range counts {
		if count["0"] > count["1"] {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func SelectLeastCommon(counts []map[string]int) string {
	result := ""
	for _, count := range counts {
		if count["0"] > count["1"] {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func CountCharactersInPositions(input []string) []map[string]int {
	results := make([]map[string]int, len(input[0]))

	for _, line := range input {
		numbers := strings.Split(line, "")
		for col, nStr := range numbers {
			if results[col] == nil {
				results[col] = map[string]int{}
			}
			results[col][nStr]++
		}
	}

	return results
}

func SelectStringsWithPrefix(input []string, prefix string) []string {
	result := []string{}
	for _, line := range input {
		if strings.HasPrefix(line, prefix) {
			result = append(result, line)
		}
	}
	return result
}

func ReadFromInput() []string {
	bytes, _ := ioutil.ReadFile("input")
	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
