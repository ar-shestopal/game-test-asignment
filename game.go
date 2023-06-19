package main

import "fmt"

// GameRunner is an interface that defines the Run method
type GameRunner interface {
	Run()
}

type ConsoleGame struct {
	State GameState
	UI    UIPresenter
}

func NewConsoleGame(ui UIPresenter, boardSize, holesCount int) ConsoleGame {
	board := NewArrayBoard(boardSize, holesCount)
	state := NewGameState(&board)
	// game := NewConsoleGame(&state, &ui)
	g := ConsoleGame{
		// Not sure if we need boardSize here,
		State: &state,
		UI:    ui,
	}
	return g
}

func (g *ConsoleGame) Run() {
	board := g.State.GetBoard()
	g.UI.DrawBoard(board)

	for !g.State.IsWon() && !g.State.IsLost() {
		row, col, action := g.UI.GetAction()
		g.State.PerformAction(row, col, Action(action))
		fmt.Println(g.State, g.State.IsLost(), g.State.IsWon())
		g.UI.DrawBoard(board)
	}

	if g.State.IsWon() {
		g.UI.GameWon()
	} else {
		g.UI.GameLost()
	}
}
