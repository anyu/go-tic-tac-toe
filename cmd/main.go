package main

func main() {
	game := setupGame()
	game.Start()
}

func setupGame() *Game {
	p1 := NewPlayer(p1Name, p1Symbol)
	p2 := NewPlayer(p2Name, p2Symbol)
	players := []*Player{p1, p2}

	return NewGame(players)
}
