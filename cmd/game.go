package main

import (
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
			rowMove, colMove := g.board.GetPlayerMove(p)
			if ok := g.board.Update(colMove, rowMove, p.symbol); !ok {
				fmt.Println("Spot already taken. Please choose another spot.")
				break
			}
			g.board.Draw(true)
			
			winner := g.CheckWinner()
			if winner != nil {
				fmt.Printf("Winner is:%s\n", winner.name)
				g.End()
			}
		}
	}
}

func (g *Game) End() {
	os.Exit(1)
}

func (g *Game) CheckWinner() *Player {
	p := g.CheckRows()
	if p != nil {
		return p
	}
	p = g.CheckColumns()
	if p != nil {
		return p
	}

	p = g.CheckDiagonals()
	if p != nil {
		return p
	}
	return nil
}

func (g *Game) CheckRows() *Player {
	for _, col := range g.board.spaces {
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
	for col := 0; col < len(g.board.spaces); col++ {
		if g.board.spaces[0][col] != " " &&
			g.board.spaces[0][col] == g.board.spaces[1][col] &&
			g.board.spaces[1][col] == g.board.spaces[2][col] {

			switch g.board.spaces[0][col] {
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
	if g.board.spaces[0][0] != " " &&
		g.board.spaces[0][0] == g.board.spaces[1][1] &&
		g.board.spaces[1][1] == g.board.spaces[2][2] {

		switch g.board.spaces[0][0] {
		case g.players[0].symbol:
			return g.players[0]
		case g.players[1].symbol:
			return g.players[1]
		}
	}

	if g.board.spaces[2][0] != " " &&
		g.board.spaces[2][0] == g.board.spaces[1][1] &&
		g.board.spaces[1][1] == g.board.spaces[0][2] {

		switch g.board.spaces[2][0] {
		case g.players[0].symbol:
			return g.players[0]
		case g.players[1].symbol:
			return g.players[1]
		}
	}
	return nil
}
