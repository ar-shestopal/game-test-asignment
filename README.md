# My Game

This is a simple game written in Go.

## Description


Vendor Take Home Interview Questions
Proxx
Rules and playable game. Review the rules and familiarize yourself with the game. You don’t
need to implement the flag functionality.
There are three parts to the exercise. For each part, please include a working coded solution
along with an explanation for choosing a certain approach.

Part 1:
Choose a data structure(s) to represent the game state. You need to keep track of the following:
- NxN board
- Location of black holes
- Counts of # of adjacent black holes
- Whether a cell is open

To represent game state I created
```
type GameState interface {
	PerformAction(x, y int, action Action)
	GetBoard() Board
	IsFinished() bool
}
```
and
```
type State struct {
	GameBoard  Board
	Won        bool
	Lost       bool
	MinesCount int
```
Where `Board` type is
```
type Board interface {
	OpenCell(x, y int)
	FlagCell(x, y int)
	Fields() interface{}
}
```
Where ```ArrayBoard``` implements it for our game.

```
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
```
Here `IsMine`, `IsFlag`(means we mark it but not open), `AdjMines` number of mines in adjaisent cells.

Part 2:
Populate your data structure with K black holes placed in random locations. Note that should
place exactly K black holes and their location should be random with a uniform distribution.

We populate the board in
```
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
```
I Gooogled and implemented specific algorithm to make uniform destributions of Holes (I called them Mines)

Part 3
For each cell without a black hole, compute and store the number of adjacent black holes. Note
that diagonals also count. E.g.
0 2 H
1 3 H
H 2 1

Method `addMine` adds Hole(Mine) on the board, and iterates over every adjaisent cess in increments it's `AdjMines`.

Part 4
Write the logic that updates which cells become visible when a cell is clicked. Note that if a cell
has zero adjacent black holes the game needs to automatically make the surrounding cells
visible.

As performing the action changes the state, I added `PerformAction` method to the state
```
func (s *State) PerformAction(row, col int, action Action) {
	fmt.Println("Performing action", action, "on cell", row, col)

	field := s.GameBoard.Fields().([][]Cell)[row][col]

	if field.IsMine && action == "open" {
		s.Lost = true
		return
	}

	// Here's a sample code snippet to demonstrate updating the IsOpen and IsFlag properties
	if action == "open" {
		s.GameBoard.OpenCell(row, col)
	} else if action == "flag" {
		s.GameBoard.FlagCell(row, col)

		field := s.GameBoard.Fields().([][]Cell)[row][col]
		if field.IsFlag && field.IsMine {
			s.MinesCount--
		} else if !field.IsFlag && field.IsMine {
			s.MinesCount++
		}
	}

	if s.MinesCount == 0 {
		s.Won = true
	}
}
```
For each cess that user has chosen, it can perform Open or Flag action.
If user opens Hole, he is loosing, if user Flags all Holes he wins.
We could also add condition to win if use opens all cells but Holes, but I ommited it.

If User flags cell that is already flagged, it toggles the flag.

Also this method on `State.MinesCount` field, to know when all Mines are Flagged.

Note that there’s no requirement to build a UI for the game. Only the logic for updating the data structure is needed.

Althos UI is not required I added simple one in the console.
It request user to input in format `row, col action(open/flag)`
Examle: `0 7 open` or `1 3 flag`

The board looks like this
```
      0  1  2  3  4  5  6  7
   ------------------------
 0 | 0| 0| 0| 0| 0| 0| 0| 0|
   ------------------------
 1 | 0| 0| 0| 0| 0| 0| 0| 0|
   ------------------------
 2 | 0| 0| 0| 0| 0| 0| 0| 0|
   ------------------------
 3 | 1| 1| 1| 0| 0| 1| 1| 1|
   ------------------------
 4 | 1| X| 1| 0| 0| 1| X| 1|
   ------------------------
 5 | 1| 1| 1| 1| 1| 2| 1| 1|
   ------------------------
 6 |  |  |  | 1| X| 1| 0| 0|
   ------------------------
 7 |  |  |  | 1| 1| 1| 0| 0|
   ------------------------
```

with row/column numbers and number of adjecent Holes (Mines) in every Cell. Holes(Mines) are marked with `X` to simplify testing.

Also there is no validation of user input.

## Setup

1. Install Go on your machine by following the instructions at: https://golang.org/doc/install
2. Clone the repository: `git clone https://github.com/ar-shestopal/proxx-console-game.git`
3. Navigate to the project directory: `cd proxx-console-game`

## Start

To start the game, run the following command:
```
go run .
```