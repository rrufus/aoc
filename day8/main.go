package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ACC = "acc"
	JMP = "jmp"
	NOP = "nop"
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

func RunDaRiddem(program []string) (int, bool) {
	accumulator := 0
	executionRegister := 0
	executedLines := map[int]bool{}

	for {
		if executionRegister == len(program) {
			// program executed successfully
			return accumulator, true
		}

		_, exists := executedLines[executionRegister]
		if exists == true {
			return accumulator, false
		} else {
			executedLines[executionRegister] = true
		}

		if executionRegister > len(program) || executionRegister < 0 {
			return accumulator, false
		}

		line := program[executionRegister]

		instructionAndValue := strings.Split(line, " ")
		instruction, valueStr := instructionAndValue[0], instructionAndValue[1]
		if instruction == NOP {
			executionRegister++
			continue
		}
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatalf("Converting value for instruction [%v] caused error: [%v]", instruction, err)
		}
		switch instruction {
		case ACC:
			accumulator += value
			executionRegister++
			continue
		case JMP:
			executionRegister += value
			continue
		default:
			log.Fatalf("Should not get here, bad instruction [%v]", instruction)
		}
	}
}

func main() {
	in := ReadFromStdIn()

	fmt.Println("Part 1")
	part1Accumulator, _ := RunDaRiddem(in)
	fmt.Println(part1Accumulator)

	fmt.Println("Part 2")
	for idx, line := range in {
		instructionAndValueStr := strings.Split(line, " ")
		instruction, valueStr := instructionAndValueStr[0], instructionAndValueStr[1]
		newProgram := make([]string, len(in))
		if instruction == NOP {
			copy(newProgram, in)
			newProgram[idx] = fmt.Sprintf("%v %v", JMP, valueStr)
		} else if instruction == JMP {
			copy(newProgram, in)
			newProgram[idx] = fmt.Sprintf("%v %v", NOP, valueStr)
		}
		if newProgram[0] != "" {
			part2Accumulator, success := RunDaRiddem(newProgram)
			if success == true {
				fmt.Println(part2Accumulator)
				break
			}
		}

	}

}
