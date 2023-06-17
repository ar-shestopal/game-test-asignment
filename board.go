package main

import (
	"math/rand"
)

type Board interface {
	OpenCell(x, y int)
	FlagCell(x, y int)
	Fields() interface{}
}

type ArrayBoard struct {
	Cells [][]Cell
}

type Cell struct {
	IsMine   bool
	IsFlag   bool
	IsOpen   bool
	AdjMines int
	X        int
	Y        int
}

func NewArrayBoard(boardSize int, minesCount int) ArrayBoard {
	b := make([][]Cell, boardSize)
	for i := range b {
		b[i] = make([]Cell, boardSize)
	}

	// Distribute mines using Sprinkling Algorithm
	availableCells := make([][2]int, boardSize*boardSize)
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			availableCells[i*boardSize+j] = [2]int{i, j}
		}
	}

	for i := 0; i < minesCount; i++ {
		index := rand.Intn(len(availableCells))
		cell := availableCells[index]
		availableCells = append(availableCells[:index], availableCells[index+1:]...)

		addMine(b, cell[0], cell[1])
	}
	return ArrayBoard{Cells: b}
}

func (b *ArrayBoard) Fields() interface{} {
	return b.Cells
}

func addMine(cells [][]Cell, x, y int) {
	// Add mine to the specified cell
	cells[x][y].IsMine = true

	// Get the board size
	boardSize := len(cells)

	// Define the eight directions to check for adjacent cells
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Iterate over the directions
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		// Check if the adjacent cell is within the board boundaries
		if nx >= 0 && nx < boardSize && ny >= 0 && ny < boardSize {
			// Increment the AdjMines count of the adjacent cell
			cells[nx][ny].AdjMines++
		}
	}
}

func (b ArrayBoard) OpenCell(x, y int) {
	// Get the board size
	boardSize := len(b.Cells)

	// Check if the cell is within the board boundaries
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		return
	}

	cell := b.Cells[x][y]
	// Check if the cell is already open
	if cell.IsOpen {
		return
	}

	// Check if the cell is a mine
	if cell.IsMine {
		return
	}

	// Open the cell
	b.Cells[x][y].IsOpen = true

	// Define the eight directions to check for adjacent cells
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Iterate over the directions
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]

		if nx < 0 || nx >= boardSize || ny < 0 || ny >= boardSize {
			continue
		}

		cell := b.Cells[nx][ny]

		if !cell.IsMine && !cell.IsOpen && !cell.IsFlag && cell.AdjMines == 0 {
			b.OpenCell(nx, ny)
		}

	}
}

func (b ArrayBoard) FlagCell(x, y int) {
	// Get the board size
	boardSize := len(b.Cells)

	// Check if the cell is within the board boundaries
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		return
	}

	// Check if the cell is already open
	if b.Cells[x][y].IsOpen {
		return
	}

	// Toggle the flag
	b.Cells[x][y].IsFlag = !b.Cells[x][y].IsFlag
}
