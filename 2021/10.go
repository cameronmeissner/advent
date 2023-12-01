package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"./utils"
)

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}


func loadJoltages() (joltages map[int]int) {
	file, err := os.Open("resources/10.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	joltages = make(map[int]int)

	for scanner.Scan() {
		joltage, err := strconv.Atoi(scanner.Text())
		utils.Check(err)
		joltages[joltage] = 1
	}
	return
}

func findJoltageDiff() int {
	joltages := loadJoltages()
	j, ok := 0, true
	diff1, diff3 := 0, 0
	for ok {
		_, ok1 := joltages[j + 1]
		_, ok2 := joltages[j + 2]
		_, ok3 := joltages[j + 3]

		if ok1 {
			diff1++
			j = j + 1
		} else if ok2 {
			j = j + 2
		} else if ok3 {
			diff3++
			j = j + 3
		}
		ok = ok1 || ok2 || ok3
	}
	diff3++
	return diff1 * diff3
}

func totalConfigs() int {
	joltages := loadJoltages()
	memo := make(map[int]int)
	return configsOf(0, memo, joltages)
}

func configsOf(joltage int, memo, joltages map[int]int) (configs int) {
	for j := range joltages {
		if j >= joltage + 1 && j <= joltage + 3 {
			if _, ok := memo[j]; !ok {
				memo[j] = max(1, configsOf(j, memo, joltages))
			}
			configs += memo[j]
		}
	}
	return
}

func main() {
	fmt.Printf("Found joltage diff product: %d\n", findJoltageDiff())
	fmt.Printf("Found %d total configurations\n", totalConfigs())
}