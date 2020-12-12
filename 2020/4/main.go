package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

type PassportPart1 struct {
	Ecl bool
	Pid bool
	Eyr bool
	Hcl bool
	Byr bool
	Iyr bool
	Cid bool
	Hgt bool
}

type PassportPart2 struct {
	Ecl string
	Pid string
	Eyr string
	Hcl string
	Byr string
	Iyr string
	Cid string
	Hgt string
}

func main() {
	in := ReadFromStdIn()

	fmt.Println("Part 1")
	passportsPart1 := []*PassportPart1{{}}
	valid := 0
	activePassportIndex := 0
	for _, line := range in {
		active := passportsPart1[activePassportIndex]
		if line == "" {

			if active.Ecl && active.Pid && active.Eyr && active.Hcl && active.Byr && active.Iyr && active.Hgt {
				valid++
			}
			activePassportIndex++
			passportsPart1 = append(passportsPart1, &PassportPart1{})
			continue
		}
		active.Ecl = active.Ecl || strings.Contains(line, "ecl")
		active.Pid = active.Pid || strings.Contains(line, "pid")
		active.Eyr = active.Eyr || strings.Contains(line, "eyr")
		active.Hcl = active.Hcl || strings.Contains(line, "hcl")
		active.Byr = active.Byr || strings.Contains(line, "byr")
		active.Iyr = active.Iyr || strings.Contains(line, "iyr")
		active.Cid = active.Cid || strings.Contains(line, "cid")
		active.Hgt = active.Hgt || strings.Contains(line, "hgt")
	}
	fmt.Println(valid)

	fmt.Println("Part 2")
	passports := []*PassportPart2{{}}
	valid = 0
	activePassportIndex = 0

	for _, line := range in {
		active := passports[activePassportIndex]

		if line == "" {
			passValid := true

			// byr
			i, err := strconv.Atoi(active.Byr)
			if err != nil {
				passValid = false
			}
			if i < 1920 || i > 2002 {
				passValid = false
			}

			// iyr
			i, err = strconv.Atoi(active.Iyr)
			if err != nil {
				passValid = false
			}
			if i < 2010 || i > 2020 {
				passValid = false
			}

			// eyr
			i, err = strconv.Atoi(active.Eyr)
			if err != nil {
				passValid = false
			}
			if i < 2020 || i > 2030 {
				passValid = false
			}

			// hgt
			if strings.Contains(active.Hgt, "in") {
				heightStr := strings.Trim(active.Hgt, "in")
				i, err = strconv.Atoi(heightStr)
				if err != nil {
					passValid = false
				}
				if i < 59 || i > 76 {
					passValid = false
				}

			} else if strings.Contains(active.Hgt, "cm") {
				heightStr := strings.Trim(active.Hgt, "cm")
				i, err = strconv.Atoi(heightStr)
				if err != nil {
					passValid = false
				}
				if i < 150 || i > 193 {
					passValid = false
				}
			} else {
				passValid = false
			}

			// hcl
			match, err := regexp.MatchString("#[0-9a-z]{6}", active.Hcl)
			if err != nil || !match {
				passValid = false
			}

			// ecl
			switch active.Ecl {
			case "amb":
			case "blu":
			case "brn":
			case "gry":
			case "grn":
			case "hzl":
			case "oth":
			default:
				passValid = false
			}

			// pid
			if len(active.Pid) != 9 {
				passValid = false
			}
			if _, err := strconv.Atoi(active.Pid); err != nil {
				passValid = false
			}

			if passValid {
				valid++
			}

			activePassportIndex++
			passports = append(passports, &PassportPart2{})
		}

		vals := strings.Split(line, " ")

		for _, val := range vals {
			keyAndVal := strings.Split(val, ":")

			key := keyAndVal[0]
			val := ""
			if len(keyAndVal) > 1 {
				val = keyAndVal[1]
			}
			switch key {
			case "ecl":
				active.Ecl = val
			case "pid":
				active.Pid = val
			case "eyr":
				active.Eyr = val
			case "hcl":
				active.Hcl = val
			case "byr":
				active.Byr = val
			case "iyr":
				active.Iyr = val
			case "cid":
				active.Cid = val
			case "hgt":
				active.Hgt = val
			}
		}
	}
	fmt.Println(valid)

}
