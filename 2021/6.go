package main

import (
	"fmt"
	"bufio"
	"os"
)


func countAnsweredQuestions() (sum int) {
	file, err := os.Open("resources/6.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	groupAnswers := make(map[rune]int)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			sum += len(groupAnswers)
			groupAnswers = make(map[rune]int)
		} else {
			for _, r := range line {
				if _, ok := groupAnswers[r]; !ok {
					groupAnswers[r] = 0
				}
				groupAnswers[r]++
			}
		}
	}
	sum += len(groupAnswers)
	return
}

func countAnsweredQuestions2() (sum int) {
	file, err := os.Open("resources/6.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	groupAnswers := make(map[rune]int)
	groupLen := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			for _, count := range groupAnswers {
				if count == groupLen {
					sum++
				}
			}
			groupAnswers = make(map[rune]int)
			groupLen = 0
		} else {
			for _, r := range line {
				if _, ok := groupAnswers[r]; !ok {
					groupAnswers[r] = 0
				}
				groupAnswers[r]++
			}
			groupLen++
		}
	}
	for _, count := range groupAnswers {
		if count == groupLen {
			sum++
		}
	}
	return
}


func main() {
	fmt.Printf("Found %d total questions answered\n", countAnsweredQuestions())
	fmt.Printf("Found %d total questions answered\n", countAnsweredQuestions2())
}