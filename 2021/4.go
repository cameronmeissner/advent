package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
)


func validatePassports() (valid int) {
	file, err := os.Open("resources/4.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fields := make(map[string]int)
	var hasRequiredFields bool

	for scanner.Scan() { 
		line := scanner.Text()
		if len(line) == 0 {
			hasRequiredFields = true
			for _, field := range requiredFields {
				if _, ok := fields[field]; !ok {
					hasRequiredFields = false
					break
				}
			}
			if hasRequiredFields {
				valid++
			}
			fields = make(map[string]int)
		} else {
			for _, e := range strings.Split(line, " ") {
				field := strings.Split(e, ":")[0]
				if _, ok := fields[field]; !ok {
					fields[field] = 1
				}
			}
		}
	}
	hasRequiredFields = true
	for _, field := range requiredFields {
		if _, ok := fields[field]; !ok {
			hasRequiredFields = false
			break
		}
	}
	if hasRequiredFields {
		valid++
	}
	return
}

func byrValid(byr string) bool {
	yr, err := strconv.Atoi(byr)
	utils.Check(err)
	if yr >= 1920 && yr <= 2002 {
		return true
	}
	return false
}

func iyrValid(iyr string) bool {
	yr, err := strconv.Atoi(iyr)
	utils.Check(err)
	if yr >= 2010 && yr <= 2020 {
		return true
	}
	return false
}

func eyrValid(eyr string) bool {
	yr, err := strconv.Atoi(eyr)
	utils.Check(err)
	if yr >= 2020 && yr <= 2030 {
		return true
	}
	return false
}

func hgtValid(hgt string) bool {
	unit := hgt[len(hgt) - 2:]
	h := string(hgt[:len(hgt) - 2])
	if len(h) > 0 && len(string(unit)) == 2 {
		height, err := strconv.Atoi(h)
		utils.Check(err)
		if unit == "cm" && height >= 150 && height <= 193 {
			return true
		} else if unit == "in" && height >= 59 && height <= 76 {
			return true
		}
	}
	return false
}

func hclValid(hcl string) bool {
	l := 0
	if hcl[0] == '#' {
		for _, c := range hcl[1:] {
			if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
				return false
			}
			l++
		}
		return l == 6
	}
	return false
}

func eclValid(ecl string) bool {
	cls := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, cl := range cls {
		if ecl == cl {
			return true
		}
	}
	return false
}

func pidValid(pid string) bool {
	l := 0
	for _, c := range pid {
		if _, err := strconv.Atoi(string(c)); err != nil {
			return false
		}
		l++
	}
	return l == 9
}

func validatePassports2() (valid int) {
	file, err := os.Open("resources/4.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fields := make(map[string]int)
	var hasRequiredFields bool

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			hasRequiredFields = true
			for _, field := range requiredFields {
				if _, ok := fields[field]; !ok {
					hasRequiredFields = false
					break
				}
			}
			if hasRequiredFields {
				valid++
			}
			fields = make(map[string]int)
		} else {
			for _, e := range strings.Split(line, " ") {
				fsp := strings.Split(e, ":")
				field, value := fsp[0], fsp[1]
				if _, ok := fields[field]; !ok {
					switch field {
					case "byr":
						if byrValid(value) {
							fields[field] = 1
						}
					case "iyr":
						if iyrValid(value) {
							fields[field] = 1 
						}
					case "eyr":
						if eyrValid(value) {
							fields[field] = 1
						}
					case "hgt":
						if hgtValid(value) {
							fields[field] = 1
						}
					case "hcl":
						if hclValid(value) {
							fields[field] = 1
						}
					case "ecl":
						if eclValid(value) {
							fields[field] = 1
						}
					case "pid":
						if pidValid(value) {
							fields[field] = 1
						}
					}
				}
			}
		}
	}

	hasRequiredFields = true
	for _, field := range requiredFields {
		if _, ok := fields[field]; !ok {
			hasRequiredFields = false
			break
		}
	}
	if hasRequiredFields {
		valid++
	}
	return
}

func main() {
	fmt.Printf("Found %d valid passports\n", validatePassports())
	fmt.Printf("Found %d valid passports\n", validatePassports2())
}