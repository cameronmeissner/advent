package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)


func loadEntries() []int {
	file, err := os.Open("resources/1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var entries []int

	for scanner.Scan() {
		entry, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		entries = append(entries, entry)
	}
	return entries
}


func reportPair1() int {
	entries := loadEntries()
	entryMap := make(map[int]int)

	for _, entry := range entries {
		if complement, ok := entryMap[2020 - entry]; ok {
			return complement * entry
		} else {
			entryMap[entry] = entry
		}
	}
	return -1
}

func reportPair2() int {
	entries := loadEntries()
	entryMap := make(map[int][2]int)

	for i := range entries {
		if comps, ok := entryMap[2020 - entries[i]]; ok {
			return comps[0] * comps[1] * entries[i]
		}
		for j := range entries {
			if i != j {
				if _, ok := entryMap[entries[i] + entries[j]]; !ok {
					entryMap[entries[i] + entries[j]] = [2]int{entries[i], entries[j]}
				}
			}
		}
	}
	return -1
}

func reportPair2Slow() int {
	entries := loadEntries()
	for i := range entries {
		for j := range entries {
			for k := range entries {
				if i != j && i != k && j != k {
					if entries[i] + entries[j] + entries[k] == 2020 {
						return entries[i] * entries[j] * entries[k]
					}
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Printf("Report Pair 1: %d\nReport Pair 2: %d\n",
	 reportPair1(), reportPair2())
}