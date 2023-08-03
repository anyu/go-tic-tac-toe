package main

import (
	"fmt"
)

func main() {
	board := NewBoard()
	board.Draw()
}

type Board struct {
	spaces [3][3]string
}

func NewBoard() *Board {
	var b Board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.spaces[i][j] = " "
		}
	}
	return &b
}

func (b *Board) Draw() {
	fmt.Println("+---+---+---+")
	for _, row := range b.spaces {
		fmt.Printf("| %s | %s | %s |\n", row[0], row[1], row[2])
		fmt.Println("+---+---+---+")
	}

	// Outputs:
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
}
