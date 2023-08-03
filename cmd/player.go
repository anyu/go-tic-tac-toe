package main

const (
	p1Name   = "Player 1"
	p1Symbol = "X"
	p2Name   = "Player 2"
	p2Symbol = "O"
)

type Player struct {
	name   string
	symbol string
}

func NewPlayer(name, symbol string) *Player {
	return &Player{
		name:   name,
		symbol: symbol,
	}
}
