package main

func main() {
	p1 := NewPlayer(p1Name, p1Symbol)
	p2 := NewPlayer(p2Name, p2Symbol)
	players := []*Player{p1, p2}

	game := NewGame(players)
	game.Start()
}
