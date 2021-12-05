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

func markCard(callouts []uint, card BingoCard) BingoCard {
	for _, callout := range callouts {
		for _, row := range card {
			for _, number := range row {
				// fmt.Printf("callout: %v - %T\n", callout, callout)
				// fmt.Printf(" number: %v - %T \n  mark: %v - %T\n", number.number, number.number, number.marked, number.marked)
				if number.number == callout {
					number.marked = true
					fmt.Println(number)
				}
			}
		}
	}

	return card
}

func isWinner(callouts []uint, card BingoCard) {
	return
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

// check cards
// mark cards that won

func main() {
	callouts, cards := getBingoData()

	fmt.Printf("%v\n", callouts)
	fmt.Printf("unmarked\n-------------\n")
	fmt.Printf("%v\n", cards[0])

	markedCard := markCard(callouts, cards[0])
	fmt.Printf("marked\n-------------\n")
	fmt.Printf("%v\n", markedCard)
}
