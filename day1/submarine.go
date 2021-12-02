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

	increases := 0

	for i, depth := range depths {
		if i+1 == len(depths) {
			break
		}

		if check_delta(depth, depths[i+1]) {
			increases = increases + 1
		}
	}

	fmt.Print(increases)
}
