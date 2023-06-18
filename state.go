package main

import "fmt"

type Action string

const (
	OpenAction = "open"
	FlagAction = "flag"
)

type GameState interface {
	PerformAction(x, y int, action Action)
	GetBoard() Board
	IsFinished() bool
}

type State struct {
	GameBoard Board
	Won       bool
	Lost      bool
}

func NewGameState(b Board) State {
	s := State{GameBoard: b}
	s.Won = false
	s.Lost = false
	return s
}

func (s *State) GetBoard() Board {
	return s.GameBoard
}

func (s *State) IsFinished() bool {
	return s.Won || s.Lost
}

func (s *State) PerformAction(row, col int, action Action) {
	fmt.Println("Performing action", action, "on cell", row, col)

	if s.GameBoard.IsHole(row, col) && action == "open" {
		s.Lost = true
		return
	}

	// Here's a sample code snippet to demonstrate updating the IsOpen and IsFlag properties
	if action == "open" {
		s.GameBoard.OpenCell(row, col)
	} else if action == "flag" {
		s.GameBoard.FlagCell(row, col)
	}

	if s.GameBoard.GetHolesCount() == 0 || s.GameBoard.GetNotFoundCount() == 0 {
		s.Won = true
	}
}
