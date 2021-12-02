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
		fmt.Printf("depth 1: %v\n", depth)
		fmt.Printf("depth 2: %v\n", depths[i+1])
		fmt.Printf("is deeper?: %t\n", check_delta(depth, depths[i+1]))
		if check_delta(depth, depths[i+1]) {
			increases = increases + 1
		}
	}

	fmt.Print(increases)
}
