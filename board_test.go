package main

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	boardSize := 8
	minesCount := 10

	board := NewArrayBoard(boardSize, minesCount)

	// Assert the dimensions of the board
	if len(board.Cells) != boardSize {
		t.Errorf("Expected board size %d, got %d", boardSize, len(board.Cells))
	}
	for _, row := range board.Cells {
		if len(row) != boardSize {
			t.Errorf("Expected row size %d, got %d", boardSize, len(row))
		}
	}

	// Assert the number of mines in the board
	count := 0
	for _, row := range board.Cells {
		for _, cell := range row {
			if cell.IsMine {
				count++
			}
		}
	}
	if count != minesCount {
		t.Errorf("Expected mines count %d, got %d", minesCount, count)
	}
}

func TestOpenCell(t *testing.T) {
	boardSize := 8
	minesCount := 10

	board := NewArrayBoard(boardSize, minesCount)

	// Test opening a non-mine cell
	x, y := 0, 0
	board.OpenCell(x, y)
	if !board.Cells[x][y].IsOpen {
		t.Errorf("Expected cell (%d, %d) to be open, but it is not", x, y)
	}

	// Test opening a mine cell
	mineX, mineY := findMineCoordinates(board)
	board.OpenCell(mineX, mineY)
	if board.Cells[mineX][mineY].IsOpen {
		t.Errorf("Expected mine cell (%d, %d) to remain closed, but it is open", mineX, mineY)
	}
}

// This code canbe refactored to be more concise
func TestOpenAdjacentCells(t *testing.T) {
	// Create a test board with a size of 3x3
	board := NewArrayBoard(3, 0)

	// Set up the board with specific cell configurations
	board.Cells[0][0].IsMine = false
	board.Cells[0][1].IsMine = true
	board.Cells[0][2].IsMine = false
	board.Cells[1][0].IsMine = false
	board.Cells[1][1].IsMine = false
	board.Cells[1][2].IsMine = false
	board.Cells[2][0].IsMine = false
	board.Cells[2][1].IsMine = false
	board.Cells[2][2].IsMine = false

	board.Cells[0][0].AdjMines = 1
	board.Cells[0][1].AdjMines = 0
	board.Cells[0][2].AdjMines = 1
	board.Cells[1][0].AdjMines = 1
	board.Cells[1][1].AdjMines = 1
	board.Cells[1][2].AdjMines = 1
	board.Cells[2][0].AdjMines = 0
	board.Cells[2][1].AdjMines = 0
	board.Cells[2][2].AdjMines = 0

	// 1 x 1     1 x 1
	// 1 1 1 =>  1 '' 1
	// 0 0 0     '' '' ''

	// ui := ConsoleUI{}

	board.OpenCell(1, 1)

	if !board.Cells[1][1].IsOpen {
		t.Errorf("Expected cell (1, 1) to be opened, but it is not")
	}
	if !board.Cells[2][0].IsOpen {
		t.Errorf("Expected cell (2, 0) to be opened, but it is not")
	}
	if !board.Cells[2][1].IsOpen {
		t.Errorf("Expected cell (2, 1) to be opened, but it is not")
	}
	if !board.Cells[2][2].IsOpen {
		t.Errorf("Expected cell (2, 2) to be opened, but it is not")
	}

	// Check that cells with mines or adjacent mines are not opened
	if board.Cells[0][1].IsOpen {
		t.Errorf("Expected cell (0, 1) to be unopened, but it is opened")
	}
	if board.Cells[1][0].IsOpen {
		t.Errorf("Expected cell (1, 0) to be unopened, but it is opened")
	}
	if board.Cells[1][2].IsOpen {
		t.Errorf("Expected cell (1, 2) to be unopened, but it is opened")
	}
}

func TestFlagCell(t *testing.T) {
	boardSize := 8
	minesCount := 10

	board := NewArrayBoard(boardSize, minesCount)

	// Test flagging a cell
	x, y := 0, 0
	board.FlagCell(x, y)
	if !board.Cells[x][y].IsFlag {
		t.Errorf("Expected cell (%d, %d) to be flagged, but it is not", x, y)
	}

	// Test unflagging a cell
	board.FlagCell(x, y)
	if board.Cells[x][y].IsFlag {
		t.Errorf("Expected cell (%d, %d) to be unflagged, but it is flagged", x, y)
	}
}

// Helper function to find the coordinates of a mine in the board
func findMineCoordinates(board ArrayBoard) (int, int) {
	for i, row := range board.Cells {
		for j, cell := range row {
			if cell.IsMine {
				return i, j
			}
		}
	}
	return -1, -1
}
