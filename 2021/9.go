package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
	"./utils"
)


func loadEntries() (entries []int) {
	file, err := os.Open("resources/9.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		entry, err := strconv.Atoi(scanner.Text())
		utils.Check(err)
		entries = append(entries, entry)
	}
	return
}

func isValidEntry(target int, preamble []int) bool {
	m := make(map[int]int)
	for _, e := range preamble {
		if _, ok := m[target - e]; ok && target - e != e {
			return true
		}
		m[e] = 1
	}
	return false
}

func findInvalidEntry() int {
	entries := loadEntries()
	if len(entries) <= 25 {
		panic("Couldn't find enough entries")
	}

	for i := 25; i < len(entries); i++ {
		if !isValidEntry(entries[i], entries[i-25:i]) {
			return entries[i]
		}
	}
	return -1
}

func findContinuousSum(target int) int {
	entries := loadEntries()
	for start := 0; start < len(entries) - 1; start++ {
		sum := entries[start] + entries[start + 1]
		idx := start + 2
		for sum < target {
			sum += entries[idx]
			idx++
		}
		if sum == target {
			sort.Ints(entries[start:idx])
			return entries[start] + entries[idx - 1]
		}
	}
	return -1
}

func main() {
	invalidEntry := findInvalidEntry()
	fmt.Printf("%d is invalid\n", invalidEntry)
	fmt.Printf("Found encryption weakness: %d\n", findContinuousSum(invalidEntry))
}