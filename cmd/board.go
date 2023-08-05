package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	errInvalidInputFormat = errors.New("\ninvalid input. Enter 2 numbers separated by a comma\n")
	errInvalidInputRange  = errors.New("\ninvalid input. Enter 2 numbers between 0 and 2\n")
)

type Board struct {
	cells [][]string
}

func NewBoard(size int) *Board {
	grid := make([][]string, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			grid[i][j] = " "
		}
	}
	return &Board{cells: grid}
}

func (b *Board) Draw() {
	fmt.Println()
	blankSpace := strings.Repeat(" ", 4)
	for i := range b.cells {
		if i != 0 {
			blankSpace = strings.Repeat(" ", 3)
		}
		fmt.Printf("%s%d", blankSpace, i)
	}
	fmt.Println("\n  +---+---+---+")

	for i, row := range b.cells {
		fmt.Printf("%d | %s | %s | %s |\n", i, row[0], row[1], row[2])
		fmt.Println("  +---+---+---+")
	}
	fmt.Println()
	return
}

func (b *Board) GetPlayerInput(p *Player) (int, int, error) {
	fmt.Printf("%s's turn. Choose a spot (eg. '2,1' for top right corner):\n", p.name)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	colMove, rowMove, err := b.parseInput(input)
	if err != nil {
		return 0, 0, err
	}

	if b.cells[rowMove][colMove] != " " {
		return 0, 0, errors.New("spot already taken. Please choose another spot")
	}

	return colMove, rowMove, nil
}

func (b *Board) Update(col, row int, symbol string) {
	b.cells[row][col] = symbol
}

func (b *Board) IsFull() bool {
	for _, row := range b.cells {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}

func (b *Board) parseInput(input string) (int, int, error) {
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

	if !b.isValidMove(num1) || !b.isValidMove(num2) {
		return 0, 0, errInvalidInputRange
	}
	return num1, num2, nil
}

func (b *Board) isValidMove(num int) bool {
	return num >= 0 && num < len(b.cells)
}
