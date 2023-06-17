package main

const DEFAULT_BOARD_SIZE = 8
const DEFAULT_MINES = 3

// GameRunner is an interface that defines the Run method
type GameRunner interface {
	Run()
}

type ConsoleGame struct {
	State GameState
	UI    UIPresenter
}

func NewConsoleGame(boardSize, holesCount int) ConsoleGame {
	board := NewArrayBoard(boardSize, holesCount)
	state := NewGameState(&board)
	ui := ConsoleUI{}
	// game := NewConsoleGame(&state, &ui)
	g := ConsoleGame{
		// Not sure if we need boardSize here,
		State: &state,
		UI:    &ui,
	}
	return g
}

func (g *ConsoleGame) Run() {
	board := g.State.GetBoard()
	g.UI.DrawBoard(board)

	for !g.State.IsFinished() {
		x, y, action := g.UI.GetAction()
		g.State.PerformAction(x, y, Action(action))
		g.UI.DrawBoard(board)
	}

	g.UI.FinishGame()
}
