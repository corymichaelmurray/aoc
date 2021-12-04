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

func main() {
	input, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(input)
	var gammaValues []string
	var epsilonValues []string
	var powerValues [][]string

	for scanner.Scan() {
		binaryString := scanner.Text()
		binarySlice := strings.Split(binaryString, "")
		powerValues = append(powerValues, binarySlice)
	}

	for i := 0; i < 12; i++ {
		zeros := 0
		ones := 0
		for _, binary := range powerValues {
			if binary[i] == "1" {
				ones++
			} else {
				zeros++
				fmt.Println("here")
			}
		}
		if ones > zeros {
			gammaValues = append(gammaValues, "1")
			epsilonValues = append(epsilonValues, "0")
		} else {
			gammaValues = append(gammaValues, "0")
			epsilonValues = append(epsilonValues, "1")
		}
	}

	gamma, err := strconv.ParseUint(strings.Join(gammaValues[:], ""), 2, 12)
	epsilon, err := strconv.ParseUint(strings.Join(epsilonValues[:], ""), 2, 12)
	fmt.Printf("gamma: %b\n", gamma)
	fmt.Printf("epsilon: %b\n", epsilon)
	fmt.Printf("power: %v\n", (gamma * epsilon))
}
