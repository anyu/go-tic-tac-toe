package main

import (
	"fmt"
	"strings"
)

// Board represents the game board with cells arranged in a 2D grid.
type Board struct {
	cells [][]string
}

// NewBoard creates a new instance of Board with the specified size.
// Each cell is initialized with an empty space.
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

// Draw prints the current state of the game board, along with indices for easy reference.
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

// Update updates the board by placing the input symbol at the specified
// row and column on the board.
func (b *Board) Update(col, row int, symbol string) {
	b.cells[row][col] = symbol
}

// IsFull returns true if the game board cells are all occupied, false otherwise.
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
