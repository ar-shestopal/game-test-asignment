package main

import (
	"fmt"
	"testing"
)

func TestPerformAction_Open(t *testing.T) {
	board := createTestBoard()
	// Create a test instance of State
	state := NewGameState(&board)

	fmt.Println(state.GameBoard.Fields().([][]Cell)[0][0])

	// Perform fman open action on a non-mine cell
	state.PerformAction(0, 0, OpenAction)
	if !state.GameBoard.Fields().([][]Cell)[0][0].IsOpen {
		t.Errorf("Expected cell (0, 0) to be open, but it is not")
	}

	// Perform an open action on a mine cell
	state.PerformAction(0, 1, OpenAction)
	if state.GameBoard.Fields().([][]Cell)[0][1].IsOpen {
		t.Errorf("Expected mine cell (0, 1) to remain closed, but it is open")
	}

	if !state.Lost {
		t.Errorf("Expected game to be lost, but it is not")
	}

	if state.Won {
		t.Errorf("Expected game to not be won, but it is")
	}
}

func TestPerformAction_Flag(t *testing.T) {
	board := createTestBoard()
	// Create a test instance of State
	state := NewGameState(&board)
	state.MinesCount = 1

	// Perform a flag action on a non-mine cell
	state.PerformAction(0, 0, FlagAction)
	if !state.GameBoard.Fields().([][]Cell)[0][0].IsFlag {
		t.Errorf("Expected cell (0, 0) to be flagged, but it is not")
	}

	if state.MinesCount != 1 {
		t.Errorf("Expected mines count to be 1, but it is %d", state.MinesCount)
	}

	// Perform a flag action on a mine cell
	// Add one more mine to the board
	state.GameBoard.Fields().([][]Cell)[0][0].IsMine = true
	state.MinesCount = 2

	state.PerformAction(0, 1, FlagAction)
	if !state.GameBoard.Fields().([][]Cell)[0][1].IsFlag {
		t.Errorf("Expected mine cell (0, 1) to be flagged, but it is not")
	}

	if state.MinesCount != 1 {
		t.Errorf("Expected mines count to be 1, but it is %d", state.MinesCount)
	}

	if state.Lost {
		t.Errorf("Expected game to not be lost, but it is")
	}

	if state.Won {
		t.Errorf("Expected game to not be won, but it is")
	}

	// Perform a flag action on a flagged cell
	state.PerformAction(0, 0, FlagAction)
	if state.GameBoard.Fields().([][]Cell)[0][0].IsFlag {
		t.Errorf("Expected cell (0, 0) to be unflagged, but it is flagged")
	}

	if state.MinesCount != 2 {
		t.Errorf("Expected mines count to be 2, but it is %d", state.MinesCount)
	}
}

func TestPerformAction_FlagAllMines(t *testing.T) {
	board := createTestBoard()
	// Create a test instance of State
	state := NewGameState(&board)
	state.MinesCount = 2

	state.GameBoard.Fields().([][]Cell)[0][2].IsMine = true
	// Perform a flag action on all mine cells
	state.PerformAction(0, 1, FlagAction)
	state.PerformAction(0, 2, FlagAction)

	fmt.Println(state.GameBoard.Fields().([][]Cell)[0][1].IsMine)
	fmt.Println(state.GameBoard.Fields().([][]Cell)[0][1].IsMine)

	if state.MinesCount != 0 {
		t.Errorf("Expected mines count to be 0, but it is %d", state.MinesCount)
	}

	if !state.Won {
		t.Errorf("Expected game to be won, but it is not")
	}

	if state.Lost {
		t.Errorf("Expected game to not be lost, but it is")
	}
}

func createTestBoard() ArrayBoard {
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

	// 1 x 1
	// 1 1 1
	// 0 0 0

	return board
}
