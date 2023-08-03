package main

import (
	"bufio"
	"fmt"
	"os"
)

type Game struct {
	board   *Board
	players []*Player
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

	g.board.Draw(true)

	for {
		for _, p := range g.players {
			g.StartPlayerTurn(p)
		}
	}
}

func (g *Game) StartPlayerTurn(p *Player) {
	for {
		rowMove, colMove, err := g.board.GetPlayerInput(p)
		if err != nil {
			fmt.Println(err)
			continue
		}

		g.board.Update(colMove, rowMove, p.symbol)
		g.board.Draw(true)

		winner, isTie := g.CheckWinner()
		if isTie {
			fmt.Println("Tie game!")
			g.End()
		}
		if winner != nil {
			fmt.Printf("Winner is:%s\n", winner.name)
			g.End()
		}
		break
	}
}

func (g *Game) End() {
	fmt.Println("Want to play again? (y/n):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input == "yes" || input == "y" {
		p1 := NewPlayer(p1Name, p1Symbol)
		p2 := NewPlayer(p2Name, p2Symbol)
		players := []*Player{p1, p2}

		game := NewGame(players)
		game.Start()
	}

	if input == "no" || input == "n" {
		fmt.Println("Exiting game. Goodbye!")
		os.Exit(1)
	}
}

func (g *Game) CheckWinner() (*Player, bool) {
	if g.isTieGame() {
		return nil, true
	}
	p := g.CheckRows()
	if p != nil {
		return p, false
	}
	p = g.CheckColumns()
	if p != nil {
		return p, false
	}

	p = g.CheckDiagonals()
	if p != nil {
		return p, false
	}
	return nil, false
}

func (g *Game) CheckRows() *Player {
	for _, col := range g.board.cells {
		if col[0] != " " && col[0] == col[1] && col[1] == col[2] {
			switch col[0] {
			case g.players[0].symbol:
				return g.players[0]
			case g.players[1].symbol:
				return g.players[1]
			}
		}
	}
	return nil
}

func (g *Game) CheckColumns() *Player {
	for col := 0; col < len(g.board.cells); col++ {
		if g.board.cells[0][col] != " " &&
			g.board.cells[0][col] == g.board.cells[1][col] &&
			g.board.cells[1][col] == g.board.cells[2][col] {

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

func (g *Game) CheckDiagonals() *Player {
	if g.board.cells[0][0] != " " &&
		g.board.cells[0][0] == g.board.cells[1][1] &&
		g.board.cells[1][1] == g.board.cells[2][2] {

		switch g.board.cells[0][0] {
		case g.players[0].symbol:
			return g.players[0]
		case g.players[1].symbol:
			return g.players[1]
		}
	}

	if g.board.cells[2][0] != " " &&
		g.board.cells[2][0] == g.board.cells[1][1] &&
		g.board.cells[1][1] == g.board.cells[0][2] {

		switch g.board.cells[2][0] {
		case g.players[0].symbol:
			return g.players[0]
		case g.players[1].symbol:
			return g.players[1]
		}
	}
	return nil
}

func (g *Game) isTieGame() bool {
	for _, row := range g.board.cells {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}
