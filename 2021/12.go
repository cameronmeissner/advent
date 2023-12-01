package main

import (
	"fmt"
	"bufio"
	"os"
	"./utils"
	"strconv"
)

func loadShipCommands() []string {
	file, err := os.Open("resources/12.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	shipCommands := make([]string, 0)
	for scanner.Scan() {
		shipCommands = append(shipCommands, scanner.Text())
	}
	return shipCommands
}

func parseCommand(command string) (byte, int) {
	value, err := strconv.Atoi(command[1:])
	utils.Check(err)
	return command[0], value
}

func manhattanDistanceTo(xpos, ypos int) int {
	if xpos < 0 {
		xpos = -xpos
	}
	if ypos < 0 {
		ypos = -ypos
	}
	return xpos + ypos
}

func shipManhattanPart1() int {
	shipCommands := loadShipCommands()
	xpos, ypos := 0, 0
	heading := 0
	for _, command := range shipCommands {
		cmd, value := parseCommand(command)
		switch cmd {
		case 'F':
			switch heading % 360 {
			case 0:
				xpos += value
			case 180:
				xpos -= value
			case -180:
				xpos -= value
			case 90:
				ypos += value
			case -90:
				ypos -= value
			case 270:
				ypos -= value
			case -270:
				ypos += value
			default:
				fmt.Printf("Unrecognized heading: %d\n", heading)
			}
		case 'L':
			heading += value
		case 'R':
			heading -= value
		case 'E':
			xpos += value
		case 'W':
			xpos -= value
		case 'N':
			ypos += value
		case 'S':
			ypos -= value
		default:
			fmt.Printf("Unrecognized command: %s\n", cmd)
		}
	}
	return manhattanDistanceTo(xpos, ypos)
}

func rotateWpRBy90(x, y int) (int, int) {
	return y, -x
}

func rotateWpLBy90(x, y int) (int, int) {
	return -y, x
}

func shipManhattanPart2() int {
	shipCommands := loadShipCommands()
	shipXPos, shipYPos := 0, 0
	wpXPos, wpYPos := 10, 1
	for _, command := range shipCommands {
		cmd, value := parseCommand(command)
		switch cmd {
		case 'F':
			shipXPos += wpXPos * value
			shipYPos += wpYPos * value
		case 'L':
			for i := 0; i < value / 90; i++ {
				wpXPos, wpYPos = rotateWpLBy90(wpXPos, wpYPos)
			}
		case 'R':
			for i := 0; i < value / 90; i++ {
				wpXPos, wpYPos = rotateWpRBy90(wpXPos, wpYPos)
			}
		case 'E':
			wpXPos += value
		case 'W':
			wpXPos -= value
		case 'N':
			wpYPos += value
		case 'S':
			wpYPos -= value
		default:
			fmt.Printf("Unrecognized command: %s\n", cmd)
		}
	}
	return manhattanDistanceTo(shipXPos, shipYPos)
}

func main() {
	fmt.Printf("Part 1 Manhattan distance from starting location: %d\n", shipManhattanPart1())
	fmt.Printf("Part 2 Manhattan distance from starting location: %d\n", shipManhattanPart2())
}