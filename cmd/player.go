package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	p1Name   = "Player 1"
	p1Symbol = "X"
	p2Name   = "Player 2"
	p2Symbol = "O"
)

var errInvalidInput = errors.New("\nInvalid input. Enter 2 numbers between 0 and 2, separated by a comma.\n")

// Player represents the player playing the game.
type Player struct {
	name   string
	symbol string
}

// NewPlayer creates a new instance of Player with the specified name and symbol.
func NewPlayer(name, symbol string) *Player {
	return &Player{
		name:   name,
		symbol: symbol,
	}
}

// GetInput prompts the player for their move and reads in input from the standard input.
// It expects the player to enter the coordinates of their move in the format "column,row".
func (p *Player) GetInput() (int, int, error) {
	fmt.Printf("%s's turn. Choose a spot (eg. '2,1' for top right corner):\n", p.name)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	return parseInput(input)
}

func parseInput(input string) (int, int, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return 0, 0, errInvalidInput
	}

	input1 := strings.TrimSpace(parts[0])
	input2 := strings.TrimSpace(parts[1])

	num1, err := strconv.Atoi(input1)
	if err != nil {
		return 0, 0, errInvalidInput
	}

	num2, err := strconv.Atoi(input2)
	if err != nil {
		return 0, 0, errInvalidInput
	}

	if !isValidMove(num1) || !isValidMove(num2) {
		return 0, 0, errInvalidInput
	}
	return num1, num2, nil
}

func isValidMove(num int) bool {
	return num >= 0 && num < 3
}
