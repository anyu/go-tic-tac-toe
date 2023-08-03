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
	cells [][]string
}

func NewBoard(size int) *Board {
	grid := make([][]string, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			grid[i][j] = " "
		}
	}
	return &Board{cells: grid}
}

func (b *Board) Draw(withGuides bool) {
	if withGuides {
		fmt.Println()
		blankSpace := strings.Repeat(" ", 4)
		for i := range b.cells {
			if i != 0 {
				blankSpace = strings.Repeat(" ", 3)
			}
			fmt.Printf("%s%d", blankSpace, i)
		}
		fmt.Println("\n  +---+---+---+")

		for i, row := range b.cells {
			fmt.Printf("%d | %s | %s | %s |\n", i, row[0], row[1], row[2])
			fmt.Println("  +---+---+---+")
		}
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Println("+---+---+---+")
	for _, row := range b.cells {
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
	for i, _ := range b.cells {
		validNums[fmt.Sprintf("%d", i)] = true
	}

	if validNums[input1] && validNums[input2] {
		num1, _ := strconv.Atoi(input1)
		num2, _ := strconv.Atoi(input2)
		return num1, num2, nil
	}
	return 0, 0, errors.New("invalid input. Try again")
}

func (b *Board) GetPlayerInput(p *Player) (int, int, error) {
	fmt.Printf("%s's turn. Choose a spot (eg. '0,0' or '2,1'):\n", p.name)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	rowMove, colMove, err := b.parseInput(input)
	if err != nil {
		return 0, 0, err
	}

	if b.cells[colMove][rowMove] != " " {
		return 0, 0, errors.New("spot already taken. Please choose another spot")
	}

	return rowMove, colMove, nil
}

func (b *Board) Update(colMove, rowMove int, symbol string) {
	b.cells[colMove][rowMove] = symbol
}
