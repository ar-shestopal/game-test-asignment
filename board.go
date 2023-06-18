package main

import (
	"math/rand"
)

type Board interface {
	OpenCell(x, y int)
	FlagCell(x, y int)
	IsHole(x, y int) bool
	Fields() interface{}
	GetHolesCount() int
	GetNotFoundCount() int
}

type ArrayBoard struct {
	Cells         [][]Cell
	HolesCount    int
	NotFoundCount int
}

type Cell struct {
	IsHole   bool
	IsFlag   bool
	IsOpen   bool
	AdjHoles int
	X        int
	Y        int
}

func NewArrayBoard(boardSize int, holesCount int) ArrayBoard {
	b := make([][]Cell, boardSize)
	for i := range b {
		b[i] = make([]Cell, boardSize)
	}

	// Distribute holes using Sprinkling Algorithm
	availableCells := make([][2]int, boardSize*boardSize)
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			availableCells[i*boardSize+j] = [2]int{i, j}
		}
	}

	board := ArrayBoard{Cells: b, NotFoundCount: boardSize*boardSize - holesCount}

	for i := 0; i < holesCount; i++ {
		index := rand.Intn(len(availableCells))
		cell := availableCells[index]
		availableCells = append(availableCells[:index], availableCells[index+1:]...)

		board.AddHole(cell[0], cell[1])
	}
	return board
}

func (b *ArrayBoard) Fields() interface{} {
	return b.Cells
}

func (b *ArrayBoard) IsHole(x, y int) bool {
	return b.Cells[x][y].IsHole
}

func (b *ArrayBoard) GetHolesCount() int {
	return b.HolesCount
}

func (b *ArrayBoard) GetNotFoundCount() int {
	return b.NotFoundCount
}

func (b *ArrayBoard) AddHole(x, y int) {
	// Add hole to the specified cell
	b.Cells[x][y].IsHole = true
	b.HolesCount++

	// Get the board size
	boardSize := len(b.Cells)

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
			// Increment the AdjHoles count of the adjacent cell
			b.Cells[nx][ny].AdjHoles++
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

	// Check if the cell is a hole
	if cell.IsHole {
		return
	}

	// Open the cell
	b.Cells[x][y].IsOpen = true
	b.NotFoundCount--

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

		if !cell.IsHole && !cell.IsOpen && !cell.IsFlag && cell.AdjHoles == 0 {
			b.OpenCell(nx, ny)
		}

	}
}

func (b *ArrayBoard) FlagCell(x, y int) {
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

	// Update the holes count
	if b.Cells[x][y].IsFlag && b.Cells[x][y].IsHole {
		b.HolesCount--
	} else if !b.Cells[x][y].IsFlag && b.Cells[x][y].IsHole {
		b.HolesCount++
	}
}
