package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoNumber struct {
	number uint
	marked bool
}

type BingoCard [][]BingoNumber

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printCard(card BingoCard) {
	mark := "O"
	for _, row := range card {
		for _, number := range row {
			if number.marked {
				mark = "X"
			}
			fmt.Printf("[%v - %v]  ", number.number, mark)
		}
		fmt.Printf("\n")
	}
}

func checkCard(callouts []uint, card BingoCard) (BingoCard, uint) {
	var turns uint
	for turn, callout := range callouts {
		for _, row := range card {
			for index, number := range row {
				if number.number == callout {
					row[index].marked = true
				}
			}
		}
		if isWinner(card) {
			turns = uint(turn)
			break
		}
	}

	return card, turns
}

func checkHorizontals(card BingoCard) bool {
	for _, row := range card {
		for i := 0; i < len(row); i++ {
			if row[i].marked == false {
				return false
			}
		}
	}
	return true
}

func checkVerticals(card BingoCard) bool {
	for i := 0; i < len(card); i++ {
		for _, row := range card {
			if row[i].marked == false {
				return false
			}
		}
	}
	return true
}

func checkLeftDiagonal(card BingoCard) bool {
	for i := 0; i < len(card); i++ {
		for j := 0; j < len(card[i]); j++ {
			if i == j {
				if card[i][j].marked == false {
					return false
				}
			}
		}
	}
	return true
}

func checkRightDiagonal(card BingoCard) bool {
	for i := 0; i < len(card); i++ {
		for j := 0; j < len(card[i]); j++ {
			if i+j == len(card)-1 {
				if card[i][j].marked == false {
					return false
				}
			}
		}
	}
	return true
}

func isWinner(card BingoCard) bool {
	if checkHorizontals(card) {
		return true
	} else if checkVerticals(card) {
		return true
	} else if checkLeftDiagonal(card) {
		return true
	} else if checkRightDiagonal(card) {
		return true
	} else {
		return false
	}
}

func getBingoData() ([]uint, []BingoCard) {
	var bingoNum BingoNumber
	var card BingoCard
	var callouts []uint
	var cards []BingoCard

	calloutFile, _ := os.Open("callouts.txt")
	cardsFile, _ := os.Open("cards.txt")

	calloutRead := csv.NewReader(calloutFile)

	calloutStrings, err := calloutRead.ReadAll()

	check(err)

	for _, call := range calloutStrings[0] {
		calloutNumber, err := strconv.Atoi(call)
		check(err)
		callouts = append(callouts, uint(calloutNumber))
	}
	fmt.Printf("length: %v\n", len(callouts))
	cardScanner := bufio.NewScanner(cardsFile)

	// get bingo cards
	for cardScanner.Scan() {
		var row []BingoNumber

		line := cardScanner.Text()
		// Fields is fuckin' tite
		rawRow := strings.Fields(line)

		if len(rawRow) == 0 {
			cards = append(cards, card)
			card = nil
			continue
		}

		for _, number := range rawRow {
			bingoNumber, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			bingoNum.number = uint(bingoNumber)
			row = append(row, bingoNum)
		}

		card = append(card, row)
	}

	return callouts, cards
}

func main() {
	callouts, cards := getBingoData()

	fmt.Printf("%v\n", callouts)
	fmt.Printf("unmarked\n-------------\n")
	printCard(cards[5])

	markedCard, turns := checkCard(callouts, cards[5])
	fmt.Printf("marked\n-------------\n")
	printCard(markedCard)
	fmt.Printf("\nturns: %v\n", turns)
}
