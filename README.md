# My Game

This is a simple game written in Go.

## Description


Vendor Take Home Interview Questions
Proxx
[Rules and playable game](https://proxx.app/). Review the rules and familiarize yourself with the game. You don’t
need to implement the flag functionality.
There are three parts to the exercise. For each part, please include a working coded solution
along with an explanation for choosing a certain approach.

## Important commentary.
To emphasize low coupling my classes interact over interfaces as much as possible. Despite that in current state of the code there is only one implementation for each interface, as for the test assignment, I believe, it is beneficial to demonstrate ability to approach this problem.

Also from experience, real world application require more using interfaces then creating a new ones, so my approach is mostly dictated by the existing code style, aggreements and requirements of a project or feature.

Also I did QA the programm extensively, only basic cases of loosing and wining, there might be bugs.

Part 1:
Choose a data structure(s) to represent the game state. You need to keep track of the following:
- NxN board
- Location of black holes
- Counts of # of adjacent black holes
- Whether a cell is open

To represent the game state I created
```
type GameState interface {
	PerformAction(x, y int, action Action)
	IsFinished() bool
}
```
and
```
type State struct {
	GameBoard  Board
	Won        bool
	Lost       bool
```
Where the `Board` type is
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
```
Here `IsHole`, `IsFlag`(means we mark it but not open), and `AdjHoles` a number of holes in adjacent cells.

The Board is responsible for opening, flagging Cells, and maintaining a count of opened/flagged cells to determine Win/Loose conditions

```func (b ArrayBoard) OpenCell(x, y int) {
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
```

Part 2:
Populate your data structure with K black holes placed in random locations. Note that should
place exactly K black holes and their location should be random with a uniform distribution.

We populate the board in
```
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

	for i := 0; i < holesCount; i++ {
		index := rand.Intn(len(availableCells))
		cell := availableCells[index]
		availableCells = append(availableCells[:index], availableCells[index+1:]...)

		addHole(b, cell[0], cell[1])
	}
	return ArrayBoard{Cells: b}
}
```
I Gooogled and implemented specific algorithm to make uniform destributions of Holes (I called them Holes)

Part 3
For each cell without a black hole, compute and store the number of adjacent black holes. Note
that diagonals also count. E.g.
0 2 H
1 3 H
H 2 1

Method `addHole` adds Hole(Hole) on the board, and iterates over every adjacent cess in increments it's `AdjHoles`.

Part 4
Write the logic that updates which cells become visible when a cell is clicked. Note that if a cell
has zero adjacent black holes the game needs to automatically make the surrounding cells
visible.

As performing the action changes the state, I added `PerformAction` method to the state
```
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
```
For each Cell that the user has chosen, it can perform an Open or Flag action.
If the user opens a Hole, he is losing, if the user Flags all Holes he wins.
If the user opens all cells that are not Holes he wins

If the user flags the Cell that is already flagged, it toggles the flag.

Also this method `State.GetHolesCount()` field, to know when all Holes are Flagged and `s.GameBoard.GetNotFoundCount()` to know when all non-Holes Cells are opened.



Note that there’s no requirement to build a UI for the game. Only the logic for updating the data structure is needed.

Although UI is not required I added a simple one in the console.
It requests the user to input in format `row, col action(open/flag)`
Example: `0 7 open` or `1 3 flag`

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

with row/column numbers and the number of adjacent Holes (Holes) in every Cell. Holes(Holes) are marked with `X` to simplify testing.

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