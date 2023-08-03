package main

import "fmt"

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() {
	fmt.Println()
	fmt.Println("Let's play tic tac toe!")

	board := NewBoard(3)
	board.Draw(true)

	p1 := NewPlayer(p1Name, p1Symbol)
	p2 := NewPlayer(p2Name, p2Symbol)
	players := []*Player{p1, p2}

	for {
		for _, p := range players {
			rowMove, colMove := board.GetPlayerMove(p)
			if ok := board.Update(colMove, rowMove, p.symbol); !ok {
				fmt.Println("Spot already taken. Please choose another spot.")
				break
			}
			board.Draw(true)
			board.CheckWinner()
		}
	}
}
