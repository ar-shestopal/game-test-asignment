package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type UIPresenter interface {
	// DrawBoard draws the board on the screen
	DrawBoard(b Board)
	GetAction() (int, int, string)
	FinishGame()
}

type ConsoleUI struct{}

// Some code duplication but this section is not mandatory.
func (ui *ConsoleUI) DrawBoard(b Board) {
	// Clear the screen
	ui.Clear()
	// Draw column numbers
	fmt.Printf("    ")
	for i := 0; i < len(b.Fields().([][]Cell)); i++ {
		fmt.Printf("%3d", i)
	}
	fmt.Println()

	// Draw horizontal separator
	fmt.Print("   ")
	for i := 0; i < len(b.Fields().([][]Cell))*3; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	// Draw row numbers and board cells
	for i, row := range b.Fields().([][]Cell) {
		// Draw row number
		fmt.Printf("%2d |", i)

		// Draw cells in the row
		for _, cell := range row {
			if cell.IsOpen {
				fmt.Print("  |")
			} else if cell.IsFlag {
				fmt.Print(" F|")
			} else if cell.IsHole {
				fmt.Print(" X|")
			} else {
				fmt.Printf("%2d|", cell.AdjHoles)
			}
		}

		fmt.Println()

		// Draw horizontal separator
		fmt.Print("   ")
		for i := 0; i < len(b.Fields().([][]Cell))*3; i++ {
			fmt.Print("-")
		}
		fmt.Println()
	}
}

// Works only on Unix-like systems
func (ui *ConsoleUI) Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (ui *ConsoleUI) GetAction() (row, col int, action string) {
	fmt.Println("Enter cell coordinates (row, col) and action (open/flag):")
	_, err := fmt.Scanf("%d %d %s", &row, &col, &action)
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return 0, 0, ""
	}

	return
}

func (ui *ConsoleUI) FinishGame() {
	fmt.Println("Game finished.")
}

func (ui *ConsoleUI) ReadSize() int {
	// Read command-line arguments
	args := os.Args[1:] // Exclude the program name itself

	// Check if an argument is provided
	if len(args) < 1 {
		fmt.Printf("Creating game with %d*%d board size", DEFAULT_BOARD_SIZE, DEFAULT_BOARD_SIZE)
		return DEFAULT_BOARD_SIZE
	}

	// Parse the argument as an integer
	size, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Invalid size argumen. Creating game with %d*%d board size \n", DEFAULT_BOARD_SIZE, DEFAULT_BOARD_SIZE)
		return DEFAULT_BOARD_SIZE
	}

	if size < 1 {
		fmt.Printf("Invalid size argumen. Creating game with %d*%d board size \n", DEFAULT_BOARD_SIZE, DEFAULT_BOARD_SIZE)
		return DEFAULT_BOARD_SIZE
	}

	fmt.Printf("Creating game with %d*%d board size \n", size, size)
	return size
}
