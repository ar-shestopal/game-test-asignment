package main

func main() {
	// For simplicity, we are using a fixed number of holes and board size
	boardSize := DEFAULT_BOARD_SIZE
	holesCount := DEFAULT_MINES

	game := NewConsoleGame(boardSize, holesCount)
	game.Run()

}
