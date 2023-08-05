package main

import (
	"bufio"
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

type Player struct {
	name   string
	symbol string
}

func NewPlayer(name, symbol string) *Player {
	return &Player{
		name:   name,
		symbol: symbol,
	}
}

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
		return 0, 0, errInvalidInputFormat
	}

	input1 := strings.TrimSpace(parts[0])
	input2 := strings.TrimSpace(parts[1])

	num1, err := strconv.Atoi(input1)
	if err != nil {
		return 0, 0, errInvalidInputRange
	}

	num2, err := strconv.Atoi(input2)
	if err != nil {
		return 0, 0, errInvalidInputRange
	}

	if !isValidMove(num1) || !isValidMove(num2) {
		return 0, 0, errInvalidInputRange
	}
	return num1, num2, nil
}

func isValidMove(num int) bool {
	return num >= 0 && num < 3
}
