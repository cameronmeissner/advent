package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"./utils"
)

func loadOps() (ops []string) {
	file, err := os.Open("resources/8.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() { 
		ops = append(ops, scanner.Text())
	}
	return
}

func parseInstructions() int {
	ops, acc := loadOps(), 0
	visited := make(map[int]int)
	idx, ok := 0, false
	for !ok {
		sp := strings.Split(ops[idx], " ")
		opName, arg := sp[0], sp[1]
		argValue, err := strconv.Atoi(arg[1:])
		utils.Check(err)
		visited[idx] = 1
		switch opName {
		case "acc":
			if arg[0] == '-' {
				argValue = -argValue
			}
			acc += argValue
			idx++
		case "jmp":
			if arg[0] == '-' {
				argValue = -argValue
			}
			idx += argValue
		case "nop":
			idx++
		}
		_, ok = visited[idx]
	}
	return acc
}

func execute(ops []string) (int, bool) {
	visited := make(map[int]int)
	idx, acc, ok := 0, 0, false
	for !ok && idx < len(ops) {
		sp := strings.Split(ops[idx], " ")
		opName, arg := sp[0], sp[1]
		argValue, err := strconv.Atoi(arg[1:])
		visited[idx] = 1
		utils.Check(err)

		if arg[0] == '-' {
			argValue = -argValue
		}

		switch opName {
		case "acc":
			acc += argValue
			idx++
		case "jmp":
			idx += argValue
		case "nop":
			idx++
		}
		_, ok = visited[idx]
	}
	if idx < len(ops) {
		return acc, false
	} else {
		return acc, true
	}
}

func flipInstruction(idx int, ops []string) []string {
	ops = []string(ops)
	if idx < len(ops) {
		opName := strings.Split(ops[idx], " ")[0]
		if opName == "jmp" {
			rs := []rune(ops[idx])
			rs[0], rs[1], rs[2] = 'n', 'o', 'p'
			ops[idx] = string(rs)
		}
		if opName == "nop" {
			rs := []rune(ops[idx])
			rs[0], rs[1], rs[2] = 'j', 'm', 'p'
			ops[idx] = string(rs)
		}
	}
	return ops
}

func parseNewInstructions() int {
	ops := loadOps()
	visited := make(map[int]int)
	var instrs []int
	idx, ok := 0, false
	for !ok {
		sp := strings.Split(ops[idx], " ")
		opName, arg := sp[0], sp[1]
		argValue, err := strconv.Atoi(arg[1:])
		utils.Check(err)
		visited[idx] = 1
		if arg[0] == '-' {
			argValue = -argValue
		}

		switch opName {
		case "acc":
			idx++
		case "jmp":
			instrs = append(instrs, idx)
			idx += argValue
		case "nop":
			instrs = append(instrs, idx)
			idx++
		}
		_, ok = visited[idx]
	}
	
	for i := len(instrs) - 1; i >= 0; i-- {
		acc, halted := execute(flipInstruction(instrs[i], ops))
		if halted {
			return acc
		}
	}
	return -1
}

func main() {
	fmt.Printf("Value of accumulator: %d\n", parseInstructions())
	fmt.Printf("Value of accumulator: %d\n", parseNewInstructions())
}