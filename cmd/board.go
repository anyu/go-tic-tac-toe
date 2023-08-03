package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		fmt.Println()
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
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Println("+---+---+---+")
	for _, row := range b.spaces {
		fmt.Printf("| %s | %s | %s |\n", row[0], row[1], row[2])
		fmt.Println("+---+---+---+")
		fmt.Println()
	}
}

func (b *Board) parseInput(input string) (int, int, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid input. Enter 2 numbers separated by a comma")
	}

	input1 := strings.TrimSpace(parts[0])
	input2 := strings.TrimSpace(parts[1])

	validNums := make(map[string]bool)
	for i, _ := range b.spaces {
		validNums[fmt.Sprintf("%d", i)] = true
	}

	if validNums[input1] && validNums[input2] {
		num1, _ := strconv.Atoi(input1)
		num2, _ := strconv.Atoi(input2)
		return num1, num2, nil
	}
	return 0, 0, errors.New("invalid input. Try again")
}

func (b *Board) GetPlayerMove(p *Player) (int, int) {
	fmt.Printf("%s's turn: Choose a spot (eg. '0,0' or '2,1')\n", p.name)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	rowMove, colMove, err := b.parseInput(input)
	if err != nil {
		fmt.Print(err)
	}

	return rowMove, colMove
}

func (b *Board) Update(colMove, rowMove int, symbol string) bool {
	// space already occupied
	if b.spaces[colMove][rowMove] != " " {
		return false
	}
	b.spaces[colMove][rowMove] = symbol
	return true
}
