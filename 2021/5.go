package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)


func seatId(label string) int {
	front, back := 0, 127
	minCol, maxCol := 0, 7
	var row, col int
	for _, c := range label[:7] {
		if c == 'F' {
			if back - front == 1 {
				row = front
				break
			} else {
				back -= ((back - front) + 1) / 2
			}
		} else if c == 'B' {
			if back - front == 1 {
				row = back
				break
			} else {
				front += ((back - front) + 1) / 2
			}
		}
	}
	for _, c := range label[7:] {
		if c == 'L' {
			if maxCol - minCol == 1 {
				col = minCol
				break
			} else {
				maxCol -= ((maxCol - minCol) + 1) / 2
			}
		} else if c == 'R' {
			if maxCol - minCol == 1 {
				col = maxCol
				break
			} else {
				minCol += ((maxCol - minCol) + 1) / 2
			}
		}
	}
	return (row * 8) + col
}

func maxSeatId() int {
	file, err := os.Open("resources/5.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	maxSeatId := -1
	for scanner.Scan() {
		seatId := seatId(scanner.Text())
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	return maxSeatId
}

func mySeatId() int {
	file, err := os.Open("resources/5.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var ids []int
	for scanner.Scan() {
		ids = append(ids, seatId(scanner.Text()))
	}
	
	sort.Ints(ids)
	for idx, id := range ids {
		if idx - 1 >= 0 && idx + 1 < len(ids) {
			if id != ids[idx - 1] + 1 || id != ids[idx + 1] - 1 {
				return ids[idx + 1] - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Printf("Found max seat ID: %d\n", maxSeatId())
	fmt.Printf("Found my seat ID: %d\n", mySeatId())
}