package main

func main() {
	// For simplicity, we are using a fixed number of holes and board size
	ui := StdOutUI{}

	// boardSize := ui.ReadSize()
	boardSize := DEFAULT_BOARD_SIZE
	holesCount := DEFAULT_MINES

	game := NewConsoleGame(&ui, boardSize, holesCount)
	game.Run()

}
