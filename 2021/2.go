package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func loadPasswords() []string {
	file, err := os.Open("resources/2.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var passwords []string

	for scanner.Scan() {
		password := scanner.Text()
		if err != nil {
			panic(err)
		}
		passwords = append(passwords, password)
	}
	return passwords
}

func parsePassswordEntry(entry string) (int, int, rune, string) {
	comps := strings.Split(entry, ":")
	rule, password := comps[0], comps[1]
	ruleComps := strings.Split(rule, " ")
	interval, letter := ruleComps[0], ruleComps[1]
	bounds := strings.Split(interval, "-")
	minBound, maxBound := bounds[0], bounds[1]
	min, err := strconv.Atoi(minBound)
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(maxBound)
	if err != nil {
		panic(err)
	}
	return min, max, []rune(letter)[0], password
}

func validatePasswords1() (valid int) {
	passwords := loadPasswords()
	var count int
	for _, entry := range passwords {
		count = 0
		min, max, letter, password := parsePassswordEntry(entry)
		for _, char := range password {
			if char == letter {
				count++
			}
		}
		if count >= min && count <= max {
			valid++
		}
	}
	return
}

func validatePasswords2() (valid int) {
	passwords := loadPasswords()
	for _, entry := range passwords {
		min, max, letter, password := parsePassswordEntry(entry)
		passwordRunes := []rune(password)[1:]
		left, right := passwordRunes[min - 1] == letter, passwordRunes[max - 1] == letter
		if (left || right) && !(left && right) {
			valid++
		}
	}
	return
}

func main() {
	fmt.Printf("Found %d valid passwords\n", validatePasswords1())
	fmt.Printf("Found %d valid passwords\n", validatePasswords2())
}