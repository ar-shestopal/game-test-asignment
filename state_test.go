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

	// Perform fman open action on a non-hole cell
	state.PerformAction(0, 0, OpenAction)
	if !state.GameBoard.Fields().([][]Cell)[0][0].IsOpen {
		t.Errorf("Expected cell (0, 0) to be open, but it is not")
	}

	// Perform an open action on a hole cell
	state.PerformAction(0, 1, OpenAction)
	if state.GameBoard.Fields().([][]Cell)[0][1].IsOpen {
		t.Errorf("Expected hole cell (0, 1) to remain closed, but it is open")
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

	// Perform a flag action on a non-hole cell
	state.PerformAction(0, 0, FlagAction)
	if !state.GameBoard.Fields().([][]Cell)[0][0].IsFlag {
		t.Errorf("Expected cell (0, 0) to be flagged, but it is not")
	}

	if state.HolesCount != 1 {
		t.Errorf("Expected holes count to be 1, but it is %d", state.HolesCount)
	}

	// Perform a flag action on a hole cell
	// Add one more hole to the board
	state.GameBoard.Fields().([][]Cell)[0][0].IsHole = true
	state.HolesCount = 2

	state.PerformAction(0, 1, FlagAction)
	if !state.GameBoard.Fields().([][]Cell)[0][1].IsFlag {
		t.Errorf("Expected hole cell (0, 1) to be flagged, but it is not")
	}

	if state.HolesCount != 1 {
		t.Errorf("Expected holes count to be 1, but it is %d", state.HolesCount)
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

	if state.HolesCount != 2 {
		t.Errorf("Expected holes count to be 2, but it is %d", state.HolesCount)
	}
}

func TestPerformAction_FlagAllHoles(t *testing.T) {
	board := createTestBoard()
	// Create a test instance of State
	state := NewGameState(&board)

	state.GameBoard.Fields().([][]Cell)[0][2].IsHole = true
	// Perform a flag action on all hole cells
	state.PerformAction(0, 1, FlagAction)
	state.PerformAction(0, 2, FlagAction)

	fmt.Println(state.GameBoard.Fields().([][]Cell)[0][1].IsHole)
	fmt.Println(state.GameBoard.Fields().([][]Cell)[0][1].IsHole)

	if state.GameBoard.GetHolesCount() != 0 {
		t.Errorf("Expected holes count to be 0, but it is %d", state.HolesCount)
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
	board.Cells[0][0].IsHole = false
	board.Cells[0][1].IsHole = true
	board.Cells[0][2].IsHole = false
	board.Cells[1][0].IsHole = false
	board.Cells[1][1].IsHole = false
	board.Cells[1][2].IsHole = false
	board.Cells[2][0].IsHole = false
	board.Cells[2][1].IsHole = false
	board.Cells[2][2].IsHole = false

	board.Cells[0][0].AdjHoles = 1
	board.Cells[0][1].AdjHoles = 0
	board.Cells[0][2].AdjHoles = 1
	board.Cells[1][0].AdjHoles = 1
	board.Cells[1][1].AdjHoles = 1
	board.Cells[1][2].AdjHoles = 1
	board.Cells[2][0].AdjHoles = 0
	board.Cells[2][1].AdjHoles = 0
	board.Cells[2][2].AdjHoles = 0

	// 1 x 1
	// 1 1 1
	// 0 0 0

	return board
}
