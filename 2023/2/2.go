package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputPath = "resources/input.txt"

	red   = "red"
	green = "green"
	blue  = "blue"

	redStartingQuantity   = 12
	greenStartingQuantity = 13
	blueStartingQuantity  = 14
)

type cubeSet map[string]int

func newCubeSet() cubeSet {
	return cubeSet{
		red:   0,
		green: 0,
		blue:  0,
	}
}

type calculationFunc func(game string) (int, error)

func getSum(calcFunc calculationFunc) (int, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return 0, nil
	}
	defer f.Close()

	var sum int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result, err := calcFunc(scanner.Text())
		if err != nil {
			return 0, err
		}
		sum += result
	}

	return sum, nil
}

func minSetPower(game string) (int, error) {
	parts := strings.Split(game, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("game with string repr %q is malformed: meta and data are not separated by a ':'", game)
	}

	data := strings.TrimSpace(parts[1])

	cs := newCubeSet()
	rounds := strings.Split(data, ";")
	for _, round := range rounds {
		revelations := strings.Split(round, ",")
		for _, revelation := range revelations {
			revelationParts := strings.Split(strings.TrimSpace(revelation), " ")
			if len(revelationParts) != 2 {
				return 0, fmt.Errorf("game with string repr %q is malformed: cannot induce revelation parts", game)
			}
			quantityStr := revelationParts[0]
			color := strings.ToLower(revelationParts[1])

			minRequired, ok := cs[color]
			if !ok {
				return 0, fmt.Errorf("game with string repr %q is malformed: unrecognized color %q", game, color)
			}

			quantity, err := strconv.Atoi(quantityStr)
			if err != nil {
				return 0, fmt.Errorf("game with string repr %q is malformed: unable to convert quantity string to int: %q", game, quantityStr)
			}

			if quantity > minRequired {
				cs[color] = quantity
			}
		}
	}

	power := 1
	for color := range cs {
		power *= cs[color]
	}

	return power, nil
}

func validID(game string) (int, error) {
	parts := strings.Split(game, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("game with string repr %q is malformed: meta and data are not separated by a ':'", game)
	}

	meta := strings.TrimSpace(parts[0])
	data := strings.TrimSpace(parts[1])

	rounds := strings.Split(data, ";")
	valid := true
	for _, round := range rounds {
		revelations := strings.Split(round, ",")
		for _, revelation := range revelations {
			revelationParts := strings.Split(strings.TrimSpace(revelation), " ")
			if len(revelationParts) != 2 {
				return 0, fmt.Errorf("game with string repr %q is malformed: cannot induce revelation parts", game)
			}
			quantityStr := revelationParts[0]
			color := strings.ToLower(revelationParts[1])

			quantity, err := strconv.Atoi(quantityStr)
			if err != nil {
				return 0, fmt.Errorf("game with string repr %q is malformed: unable to convert quantity string to int: %q", game, quantityStr)
			}

			var remainder int
			switch color {
			case red:
				remainder = redStartingQuantity - quantity
			case green:
				remainder = greenStartingQuantity - quantity
			case blue:
				remainder = blueStartingQuantity - quantity
			default:
				return 0, fmt.Errorf("game with string repr %q is malformed: unrecognized color %q", game, color)
			}

			if remainder < 0 {
				valid = false
				break
			}
		}
	}
	if valid {
		metaParts := strings.Split(meta, " ")
		if len(metaParts) != 2 {
			return 0, fmt.Errorf("game with string repr %q is malformed: metadata is malformed", game)
		}
		gameID, err := strconv.Atoi(metaParts[1])
		if err != nil {
			return 0, fmt.Errorf("game with string repr %q is malformed: unable to convert game ID string to int", game)
		}
		return gameID, nil
	}

	return 0, nil
}

func main() {
	val, err := getSum(validID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[PART 1] valid ID sum is: %d\n", val)

	val, err = getSum(minSetPower)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[PART 2] min set power sum is: %d\n", val)
}
