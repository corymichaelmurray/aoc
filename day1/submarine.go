package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check_delta(firstDepth int, secondDepth int) bool {
	if firstDepth < secondDepth {
		return true
	} else {
		return false
	}
}

func sum_window(window []int) int {
	result := 0
	for _, v := range window {
		result += v
	}
	return result
}

func simple_check(depths []int) int {
	increases := 0
	for i, depth := range depths {
		if i+1 == len(depths) {
			break
		}
		if check_delta(depth, depths[i+1]) {
			increases = increases + 1
		}
	}
	return increases
}

func window_check(depths []int) int {
	increases := 0
	for i, depth := range depths {
		if i+2 == len(depths) {
			break
		}
		if i-1 < 0 {
			continue
		}

		windowOne := []int{depths[i-1], depth, depths[i+1]}
		windowTwo := []int{depth, depths[i+1], depths[i+2]}
		if check_delta(sum_window(windowOne), sum_window(windowTwo)) {
			increases = increases + 1
		}
	}
	return increases
}

func main() {
	f, err := os.Open("input.txt")
	check(err)

	var depths []int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)
		depths = append(depths, depth)
	}

	simple_increases := simple_check(depths)
	window_increases := window_check(depths)

	fmt.Printf("Part 1: %v\n", simple_increases)
	fmt.Printf("Part 2: %v\n", window_increases)
}
