package main

import (
	"fmt"
	"bufio"
	"os"
	"./utils"
)

func loadSeatMap() [][]byte {
	file, err := os.Open("resources/11.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	seatMap := make([][]byte, 0)
	for scanner.Scan() {
		text := []byte(scanner.Text())
		if len(text) > 0 {
			seatMap = append(seatMap, text)
		}
	}
	return seatMap
}

func numOccupiedAdjacentPart1(x, y int, seatMap [][]byte) (adjacent int) {
	height, width := len(seatMap), len(seatMap[0])
	xcords := []int{x - 1, x, x + 1}
	ycords := []int{y - 1, y, y + 1}
	for _, xcord := range xcords {
		for _, ycord := range ycords {
			if xcord >= 0 && xcord < width && ycord >= 0 && ycord < height {
				if seatMap[ycord][xcord] == '#' && (xcord != x || ycord != y) {
					adjacent++
				}
			}
		}
	}
	return
}

func numOccupiedAdjacentPart2(x, y int, seatMap [][]byte) (adjacent int) {
	height, width := len(seatMap), len(seatMap[0])
	straightStartCords := [][]int{[]int{x - 1, y}, []int{x + 1, y}, []int{x, y - 1}, []int{x, y + 1}}
	diagStartCords := [][]int{[]int{x - 1, y - 1}, []int{x - 1, y + 1}, []int{x + 1, y - 1}, []int{x + 1, y + 1}}

	for _, startCord := range straightStartCords {
		xcord, ycord := startCord[0], startCord[1]
		for xcord >= 0 && xcord < width && ycord >= 0 && ycord < height {
			if seatMap[ycord][xcord] == '#' {
				adjacent++
				break
			}
			if seatMap[ycord][xcord] == 'L' {
				break
			}
			if startCord[0] < x {
				xcord--
			} else if startCord[0] > x {
				xcord++
			} else if startCord[1] < y {
				ycord--
			} else {
				ycord++
			}
		}
	}

	for _, startCord := range diagStartCords {
		if adjacent >= 5 {
			return
		}
		xcord, ycord := startCord[0], startCord[1]
		for xcord >= 0 && xcord < width && ycord >= 0 && ycord < height {
			if seatMap[ycord][xcord] == '#' {
				adjacent++
				break
			}
			if seatMap[ycord][xcord] == 'L' {
				break
			}
			if startCord[0] < x && startCord[1] < y {
				xcord--
				ycord--
			} else if startCord[0] > x && startCord[1] < y {
				xcord++
				ycord--
			} else if startCord[0] < x && startCord[1] > y {
				xcord--
				ycord++
			} else {
				xcord++
				ycord++
			}
		}
	}
	return
}

func getOccupied(seatMap [][]byte) (occupied int) {
	for xcord := 0; xcord < len(seatMap[0]); xcord++ {
		for ycord := 0; ycord < len(seatMap); ycord++ {
			if seatMap[ycord][xcord] == '#' {
				occupied++
			}
		}
	}
	return
}

func runSeatSystem(seatMap [][]byte, adjFunc func(int, int, [][]byte) int, adjThreshold int) int {
	height, width := len(seatMap), len(seatMap[0])
	currentSeatMap := seatMap
	var stateChanged bool

	for {
		nextSeatMap := make([][]byte, height)
		stateChanged = false
		for ycord := 0; ycord < height; ycord++ {
			nextSeatMap[ycord] = make([]byte, width)
			for xcord := 0; xcord < width; xcord++ {
				if currentSeatMap[ycord][xcord] == 'L' && adjFunc(xcord, ycord, currentSeatMap) == 0 {
					nextSeatMap[ycord][xcord] = '#'
					if !stateChanged {
						stateChanged = true
					}
				} else if currentSeatMap[ycord][xcord] == '#' && adjFunc(xcord, ycord, currentSeatMap) >= adjThreshold {
					nextSeatMap[ycord][xcord] = 'L'
					if !stateChanged {
						stateChanged = true
					}
				} else {
					nextSeatMap[ycord][xcord] = currentSeatMap[ycord][xcord]
				}
			}
		}
		if !stateChanged {
			break
		}
		currentSeatMap = nextSeatMap
	}
	return getOccupied(currentSeatMap)
}

func main() {
	seatMap := loadSeatMap()
	fmt.Printf("Total seats occupied part 1: %d\n", runSeatSystem(seatMap, numOccupiedAdjacentPart1, 4))
	fmt.Printf("Total seats occupied part 2: %d\n", runSeatSystem(seatMap, numOccupiedAdjacentPart2, 5))
}