package main

import (
	"bufio"
	"fmt"
	"os"
)

// Game represents a tic-tac-toe game, keeping track of the board, players, and game state.
type Game struct {
	board           *Board
	players         []*Player
	activePlayerIdx int
	gameOver        bool
}

// NewGame represents a new instance of Game with the specified players.
func NewGame(p []*Player) *Game {
	return &Game{
		board:   NewBoard(3),
		players: p,
	}
}

// Start begins the game loop for players to take turns until the game is over.
// Once the game is over, players have the option to play again.
func (g *Game) Start() {
	fmt.Println()
	fmt.Println("Let's play tic tac toe!")

	g.board.Draw()
	for !g.gameOver {
		g.takeTurn()
	}
	reset(g.players)
}

func (g *Game) takeTurn() {
	for {
		activePlayer := g.players[g.activePlayerIdx]
		colIndex, rowIndex, err := activePlayer.GetInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if g.board.cells[rowIndex][colIndex] != " " {
			fmt.Println("\nSpot already taken. Please choose another spot.")
			continue
		}

		g.board.Update(colIndex, rowIndex, activePlayer.symbol)
		g.board.Draw()

		if winner := g.checkWinner(); winner != nil {
			fmt.Printf("Winner is: %s (%s)\n", winner.name, winner.symbol)
			g.gameOver = true
		} else if g.board.IsFull() {
			fmt.Println("Tie game!")
			g.gameOver = true
		}
		g.activePlayerIdx = g.activePlayerIdx ^ 1 // toggle between 0 and 1

		break
	}
}

func (g *Game) checkWinner() *Player {
	if p := g.checkRows(); p != nil {
		return p
	}
	if p := g.checkColumns(); p != nil {
		return p
	}
	if p := g.checkDiagonals(); p != nil {
		return p
	}
	return nil
}

func (g *Game) checkRows() *Player {
	for _, row := range g.board.cells {
		if row[0] == row[1] && row[1] == row[2] {

			switch row[0] {
			case g.players[0].symbol:
				return g.players[0]
			case g.players[1].symbol:
				return g.players[1]
			}
		}
	}
	return nil
}

func (g *Game) checkColumns() *Player {
	cells := g.board.cells

	for col := range cells {
		if cells[0][col] == cells[1][col] &&
			cells[1][col] == cells[2][col] {

			switch g.board.cells[0][col] {
			case g.players[0].symbol:
				return g.players[0]
			case g.players[1].symbol:
				return g.players[1]
			}
		}
	}
	return nil
}

func (g *Game) checkDiagonals() *Player {
	cells := g.board.cells

	// check both diagonals
	if cells[0][0] != " " && cells[0][0] == cells[1][1] && cells[1][1] == cells[2][2] ||
		cells[2][0] != " " && cells[2][0] == cells[1][1] && cells[1][1] == cells[0][2] {

		switch cells[1][1] { // check middle symbol
		case g.players[0].symbol:
			return g.players[0]
		case g.players[1].symbol:
			return g.players[1]
		}
	}
	return nil
}

func reset(players []*Player) {
	fmt.Println("\nWant to play again? (y/n):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input == "yes" || input == "y" {
		game := NewGame(players)
		game.Start()
		return
	}

	if input == "no" || input == "n" {
		fmt.Println("Exiting game. Goodbye.")
		return
	}
}
