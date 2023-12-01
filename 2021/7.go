package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"./utils"
)


func loadRules1() map[string][]string {
	file, err := os.Open("resources/7.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	rules := make(map[string][]string)
	for scanner.Scan() {
		rule := scanner.Text()
		rule = rule[:len(rule) - 1]
		bag := rule[:strings.Index(rule, " bags ")]
		if _, ok := rules[bag]; !ok && strings.Index(rule, "contain no other bags") == -1 {
			var bags []string
			contained := strings.Split(rule, " bags contain ")[1]
			sp := strings.Split(contained, ", ")
			for _, s := range sp {
				sidx := strings.Index(s, " ")
				bidx := strings.Index(s, " bag")
				if sidx == -1 || bidx == -1 {
					panic("parsing error")
				}
				bags = append(bags, s[sidx+1:bidx])
			}
			rules[bag] = bags
		}
	}
	return rules
}

func findShinyGold() (colors int) {
	rules := loadRules1()
	contains := make(map[string]bool)
	for color, _ := range rules {
		containsShinyGold(color, rules, contains)
	}
	return len(contains)
}

func containsShinyGold(color string, rules map[string][]string, contains map[string]bool) bool {
	colorRules, ok := rules[color]
	if !ok {
		return false
	}
	for _, containedColor := range colorRules {
		if containedColor == "shiny gold" {
			contains[color] = true
			return true
		} else if _, ok := contains[containedColor]; ok {
			contains[color] = true
			return true
		} else if containsShinyGold(containedColor, rules, contains) {
			contains[color] = true
			return true
		}
	}
	return false
}

func loadRules2() map[string][]string {
	file, err := os.Open("resources/7.txt")
	utils.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	rules := make(map[string][]string)
	for scanner.Scan() {
		rule := scanner.Text()
		rule = rule[:len(rule) - 1]
		bag := rule[:strings.Index(rule, " bags ")]
		if _, ok := rules[bag]; !ok && strings.Index(rule, "contain no other bags") == -1 {
			var bags []string
			contained := strings.Split(rule, " bags contain ")[1]
			sp := strings.Split(contained, ", ")
			for _, s := range sp {
				sidx := strings.Index(s, " ")
				bidx := strings.Index(s, " bag")
				if sidx == -1 || bidx == -1 {
					panic("parsing error")
				}
				bags = append(bags, fmt.Sprintf("%s:%s", s[:sidx], s[sidx+1:bidx]))
			}
			rules[bag] = bags
		}
	}
	return rules
}

func bagCount() int {
	rules := loadRules2()
	return bagCountRec("shiny gold", rules)
}

func bagCountRec(color string, rules map[string][]string) int {
	rule, ok := rules[color]
	if !ok {
		return 0
	} else {
		s := 0
		for _, r := range rule {
			sp := strings.Split(r, ":")
			q, c := sp[0], sp[1]
			quantity, err := strconv.Atoi(q)
			utils.Check(err)
			s += (quantity * bagCountRec(c, rules)) + quantity
		}
		return s
	}
}

func main() {
	fmt.Printf("Found %d colors that can eventually contain shiny gold\n", findShinyGold())
	fmt.Printf("Found %d bags required inside of shiny gold\n", bagCount())
}