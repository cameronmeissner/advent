package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../utils"
)

func loadShuttleData() (int, string) {
	file, err := os.Open("resources/13.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	notes := make([]string, 0)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	departureTime, err := strconv.Atoi(notes[0])
	utils.Check(err)

	return departureTime, notes[1]
}

func getShuttleDepartureTimes(schedule string) []int {
	scheduleEntries := strings.Split(schedule, ",")
	departureTimes := make([]int, 0)
	for _, entry := range scheduleEntries {
		if entry != "x" {
			departureTime, err := strconv.Atoi(entry)
			utils.Check(err)
			departureTimes = append(departureTimes, departureTime)
		}
	}
	return departureTimes
}

func shuttleSearchPart1() int {
	minDepartureTime, schedule := loadShuttleData()
	shuttleDepartureTimes := getShuttleDepartureTimes(schedule)
	time := minDepartureTime
	var shuttleId, delay int
	var foundShuttle bool

	for {
		for _, departureTime := range shuttleDepartureTimes {
			if time%departureTime == 0 {
				shuttleId = departureTime
				delay = time - minDepartureTime
				foundShuttle = true
				break
			}
		}
		if foundShuttle {
			break
		}
		time++
	}
	return shuttleId * delay
}

func shuttleSearchPart2() int {
	_, schedule := loadShuttleData()
	scheduleLen := len(schedule)
	busIdOffsets := make(map[int]int)

	for idx, busId := range strings.Split(schedule, ",") {
		if busId != "x" {
			busIdNumeric, err := strconv.Atoi(busId)
			utils.Check(err)
			busIdOffsets[idx] = busIdNumeric
		}
	}

	var isViable bool
	timestamp := 0
	for {
		isViable = true
		for t := timestamp; t < timestamp+scheduleLen; t++ {
			offset := t - timestamp
			busId, ok := busIdOffsets[offset]
			if ok {
				if t%busId != 0 {
					isViable = false
					break
				}
			} else {
				continue
			}
		}
		if isViable {
			return timestamp
		}
		timestamp += busIdOffsets[0]
	}
}

func main() {
	fmt.Printf("Shuttle ID - delay product: %d\n", shuttleSearchPart1())
	fmt.Printf("Min matching timestamp: %d\n", shuttleSearchPart2())
}
