package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// types for problem

type SeaFloor [][]int

// error check

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// load file
func loadFloor(file string) SeaFloor {
	var floor SeaFloor
	f, err := os.Open(file)
	check(err)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var heights []int
		line := scanner.Text()
		asciiHeights := strings.Split(line, "")
		for _, asciiHeight := range asciiHeights {
			heightValue, err := strconv.Atoi(asciiHeight)
			check(err)
			heights = append(heights, (heightValue))
		}
		floor = append(floor, heights)
	}

	return floor
}

func findSmoke(floor SeaFloor) []int {
	var lowPoints []int
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			// if on first row AND first column,
			// only compare lower row and right column
			if i == 0 && j == 0 {
				if floor[i][j] < floor[i+1][j] && floor[i][j] < floor[i][j+1] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if on first row, check left, right, down
			} else if i == 0 && j > 0 && j != len(floor[i])-1 {
				if floor[i][j] < floor[i+1][j] && floor[i][j] < floor[i][j+1] && floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if on last column and first row,
				// only compare lower row and left column
			} else if i == 0 && j == len(floor[i])-1 {
				if floor[i][j] < floor[i+1][j] && floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if on any row not  first or last, but first column,
				// check upper & lower rows & right column
			} else if j == 0 && i > 0 && i != len(floor)-1 {
				if floor[i][j] < floor[i-1][j] && floor[i][j] < floor[i+1][j] && floor[i][j] < floor[i][j+1] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if on any row not first or last, but last column
				// check upper & lower rows & left column
			} else if j == len(floor[i])-1 && i > 0 {
				if floor[i][j] < floor[i-1][j] && floor[i][j] < floor[i+1][j] && floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if on last row and first column
				// check upper row and right column
			} else if i == len(floor)-1 && j == 0 {
				if floor[i][j] < floor[i-1][j] && floor[i][j] < floor[i][j+1] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if on last row and last column
				// cehck left column and upper row
			} else if i == len(floor)-1 && j == len(floor[i])-1 {
				if floor[i][j] < floor[i-1][j] && floor[i][j] < floor[i][j-1] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if on last row,
				// check upper row & left & right columns
			} else if i == len(floor)-1 && j > 0 {
				if floor[i][j] < floor[i][j-1] && floor[i][j] < floor[i][j+1] && floor[i][j] < floor[i-1][j] {
					lowPoints = append(lowPoints, floor[i][j])
				}
				// if anywhere else, check all adjacent values
			} else {
				if floor[i][j] < floor[i][j-1] && floor[i][j] < floor[i][j+1] && floor[i][j] < floor[i-1][j] && floor[i][j] < floor[i+1][j] {
					lowPoints = append(lowPoints, floor[i][j])
				}
			}
		}
	}
	return lowPoints
}

func calculateDanger(heights []int) int {
	var dangerValues []int
	for _, height := range heights {
		dangerValue := height + 1
		dangerValues = append(dangerValues, dangerValue)
	}
	var danger int
	for _, dangerValue := range dangerValues {
		danger = danger + dangerValue
	}
	return danger
}

func main() {
	seaFloor := loadFloor("input.txt")
	lowPoints := findSmoke(seaFloor)
	danger := calculateDanger(lowPoints)
	fmt.Printf("danger: %v\n", danger)
}
