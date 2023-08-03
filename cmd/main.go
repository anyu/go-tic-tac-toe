package main

import (
	"fmt"
	"strings"
)

func main() {

	game := NewGame()
	game.Start()
}

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() {
	fmt.Println("Let's play tic tac toe!")
	board := NewBoard(3)

	board.Draw(true)

	fmt.Println("Player 1 (X): Choose a spot. Eg '0,0' or '2,1'")

	var player1move string
	if _, err := fmt.Scan(&player1move); err != nil {
		return
	}
	if !board.checkMoveValidity(player1move) {
		fmt.Print("Invalid move. Please enter a valid spot.")
	}
}

type Board struct {
	spaces [][]string
}

func NewBoard(size int) *Board {
	grid := make([][]string, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			grid[i][j] = " "
		}
	}
	return &Board{spaces: grid}
}

func (b *Board) Draw(withGuides bool) {
	if withGuides {
		blankSpace := strings.Repeat(" ", 4)
		for i := range b.spaces {
			if i != 0 {
				blankSpace = strings.Repeat(" ", 3)
			}
			fmt.Printf("%s%d", blankSpace, i)
		}
		fmt.Println("\n  +---+---+---+")

		for i, row := range b.spaces {
			fmt.Printf("%d | %s | %s | %s |\n", i, row[0], row[1], row[2])
			fmt.Println("  +---+---+---+")
		}
		return
	}

	fmt.Println("+---+---+---+")
	for _, row := range b.spaces {
		fmt.Printf("| %s | %s | %s |\n", row[0], row[1], row[2])
		fmt.Println("+---+---+---+")
	}
}

func (b *Board) checkMoveValidity(input string) bool {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return false
	}
	rowNumInput := strings.TrimSpace(parts[0])
	colNumInput := strings.TrimSpace(parts[1])

	validNums := make(map[string]bool)
	for i, _ := range b.spaces {
		validNums[fmt.Sprintf("%d", i)] = true
	}

	if validNums[rowNumInput] && validNums[colNumInput] {
		return true
	}
	return false
}

func (b *Board) MarkSpace(x, y int, mark string) {

}

func (b *Board) CheckRows() bool {
	return false
}

func (b *Board) CheckColumns() bool {
	return false
}

func (b *Board) CheckDiagonals() bool {
	return false
}

func (b *Board) CheckWin() bool {
	return false
}
