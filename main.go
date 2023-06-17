package main

func main() {
	// For simplicity, we are using a fixed number of mines and board size
	boardSize := DEFAULT_BOARD_SIZE
	minesCount := DEFAULT_MINES

	game := NewConsoleGame(boardSize, minesCount)
	game.Run()

}
