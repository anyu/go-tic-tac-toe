package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errInvalidInputFormat = errors.New("\nInvalid input. Enter 2 numbers separated by a comma.\n")
	errInvalidInputRange  = errors.New("\nInvalid input. Enter 2 numbers between 0 and 2.\n")
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
