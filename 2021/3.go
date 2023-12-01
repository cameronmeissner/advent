package main

import (
	"fmt"
	"bufio"
	"os"
)

func traverse1() (trees int) {
	file, err := os.Open("resources/3.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var x int

	for scanner.Scan() {
		if x != 0 {
			level := scanner.Text()
			if level[x % 31] == '#' {
				trees++
			}
		}
		x += 3
	}
	return
}

func traverse2(vShift, hShift int) (trees int) {
	file, err := os.Open("resources/3.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var x, y int

	for scanner.Scan() {
		if y % vShift == 0 {
			if x != 0 {
				level := scanner.Text()
				if level[x % 31] == '#' {
					trees++
				}
			}
			x += hShift
		}
		y++
	}
	return
}

func treeTraversalProduct() int {
	prod := 1
	slopes := [][]int{[]int{1, 1}, []int{3, 1}, []int{5, 1}, []int{7, 1}, []int{1, 2}}
	for _, slope := range slopes {
		prod *= traverse2(slope[1], slope[0])
	}
	return prod
}

func main() {
	fmt.Printf("Traversed %d trees\n", traverse1())
	fmt.Printf("Tree Traversal Product: %d\n", treeTraversalProduct())
}