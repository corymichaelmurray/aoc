package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check_travel(direction string, travel int, distance int, depth int) (int, int) {
	if direction == "forward" {
		return distance + travel, depth
	} else if direction == "up" {
		return distance, depth - travel
	} else {
		return distance, depth + travel
	}
}

func check_travel_with_aim(direction string, travel int, distance int, depth int, aim int) (int, int, int) {
	if direction == "forward" {
		depth = depth + (travel * aim)
		return distance + travel, depth, aim
	} else if direction == "up" {
		aim = aim - travel
		return distance, depth, aim
	} else {
		aim = aim + travel
		return distance, depth, aim
	}
}

func part_one() (int, int) {
	travel := 0
	depth := 0

	file, err := os.Open("movement.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instructions := strings.Fields(scanner.Text())
		depthString := instructions[1]
		depthTravel, err := strconv.Atoi(depthString)
		check(err)
		travel, depth = check_travel(instructions[0], depthTravel, travel, depth)
	}

	return travel, depth
}

func part_two() (int, int) {
	travel := 0
	depth := 0
	aim := 0

	file, err := os.Open("movement.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instructions := strings.Fields(scanner.Text())
		depthString := instructions[1]
		depthTravel, err := strconv.Atoi(depthString)
		check(err)
		travel, depth, aim = check_travel_with_aim(instructions[0], depthTravel, travel, depth, aim)
	}

	return travel, depth
}

func main() {
	travelOne, depthOne := part_one()
	travelTwo, depthTwo := part_two()
	fmt.Printf("part one:\n  depth: %v\n  travel: %v\n", depthOne, travelOne)
	fmt.Printf("part one answer: %v\n", (depthOne * travelOne))
	fmt.Printf("part two:\n  depth: %v\n  travel: %v\n", depthTwo, travelTwo)
	fmt.Printf("part two answer: %v\n", (depthTwo * travelTwo))
}
