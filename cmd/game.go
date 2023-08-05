package main

import (
	"bufio"
	"fmt"
	"os"
)

type Game struct {
	board    *Board
	players  []*Player
	gameOver bool
}

func NewGame(p []*Player) *Game {
	return &Game{
		board:   NewBoard(3),
		players: p,
	}
}

func (g *Game) Start() {
	fmt.Println()
	fmt.Println("Let's play tic tac toe!")

	g.board.Draw()

	for !g.gameOver {
		for _, p := range g.players {
			if g.gameOver {
				break
			}
			g.startPlayerTurn(p)
		}
	}
	reset(g.players)
}

func (g *Game) startPlayerTurn(p *Player) {
	for {
		colMove, rowMove, err := g.board.GetPlayerInput(p)
		if err != nil {
			fmt.Println(err)
			continue
		}

		g.board.Update(colMove, rowMove, p.symbol)
		g.board.Draw()

		if winner, found := g.checkWinner(); found {
			fmt.Printf("Winner is: %s (%s)\n", winner.name, winner.symbol)
			g.gameOver = true
		} else if g.board.IsFull() {
			fmt.Println("Tie game!")
			g.gameOver = true
		}

		break
	}
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

func (g *Game) checkWinner() (*Player, bool) {
	if p := g.checkRows(); p != nil {
		return p, true
	}
	if p := g.checkColumns(); p != nil {
		return p, true
	}

	if p := g.checkDiagonals(); p != nil {
		return p, true
	}
	return nil, false
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
