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

	if state.GameBoard.GetHolesCount() != 1 {
		t.Errorf("Expected holes count to be 1, but it is %d", state.GameBoard.GetHolesCount())
	}

	// Perform a flag action on a hole cell
	// Add one more hole to the board
	state.GameBoard.Fields().([][]Cell)[0][0].IsHole = true
	state.GameBoard.(*ArrayBoard).HolesCount = 2

	state.PerformAction(0, 1, FlagAction)
	if !state.GameBoard.Fields().([][]Cell)[0][1].IsFlag {
		t.Errorf("Expected hole cell (0, 1) to be flagged, but it is not")
	}

	if state.GameBoard.GetHolesCount() != 1 {
		t.Errorf("Expected holes count to be 1, but it is %d", state.GameBoard.GetHolesCount())
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

	if state.GameBoard.GetHolesCount() != 2 {
		t.Errorf("Expected holes count to be 2, but it is %d", state.GameBoard.GetHolesCount())
	}
}

func TestPerformAction_FlagAllHoles(t *testing.T) {
	board := createTestBoard()
	// Create a test instance of State
	state := NewGameState(&board)
	state.GameBoard.(*ArrayBoard).AddHole(0, 2)

	state.PerformAction(0, 1, FlagAction)
	state.PerformAction(0, 2, FlagAction)

	fmt.Println(state.GameBoard.Fields().([][]Cell)[0][1].IsHole)
	fmt.Println(state.GameBoard.Fields().([][]Cell)[0][1].IsHole)

	if state.GameBoard.GetHolesCount() != 0 {
		t.Errorf("Expected holes count to be 0, but it is %d", state.GameBoard.GetHolesCount())
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
	board.AddHole(0, 1)

	// 1 x 1
	// 1 1 1
	// 0 0 0

	return board
}
