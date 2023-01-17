package main

import (
	"fmt"
)

func main() {
    createNewGame()
}

func createNewGame() {
    var width, height int

    fmt.Print("Enter width of the board: ")
    fmt.Scan(&width)
    fmt.Print("Enter height of the board: ")
    fmt.Scan(&height)

    // fmt.Print(width, height)
    // TODO: Implement checks for max diff 2 and min size 6x8

	board := make([][]int, height)
	for i := 0; i < int(height); i++ {
		board[i] = make([]int, width)
	}
    
    // board := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
    p1Moves := make([]int, 0)
    p2Moves := make([]int, 0)
    isP1Turn := true
    for {
        printBoard(board, p1Moves, p2Moves)
        
        move := getNextMove(width) // Width passed through for user-message purpose

        // Check if the move is within rules of the game
		if move < 0 || move >= width {
			fmt.Println("Invalid move!")
			continue
		}
		if board[0][move] != 0 {
			fmt.Println("Column is full!")
			continue
		}
        
        // Place the move
		for i := height - 1; i >= 0; i-- {
			if board[i][move] == 0 {
				board[i][move] = func() int { if isP1Turn { return 1 } else { return 2 } }() // Makeshift ternary operator
				break
			}
		}

        // Track the moves
        if isP1Turn {
            p1Moves = append(p1Moves, move + 1)
        } else {
            p2Moves = append(p2Moves, move + 1)
        }

        isP1Turn = !isP1Turn
    }
}

func getNextMove(width int) (int) {
    var move int
    fmt.Printf("Enter column number (1 - %d)\n", width)
    fmt.Scan(&move)
    return move - 1
}

func printBoard(board [][]int, p1Moves []int, p2Moves []int) {
	for i := range board {
		for j := range board[i] {
            // fmt.Print(board[i][j])
			switch board[i][j] {
			case 0:
				fmt.Print("_ ")
			case 1:
				fmt.Print("X ")
			case 2:
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
    fmt.Println("Player 1 moves:", p1Moves)
    fmt.Println("Player 2 moves:", p2Moves)
}
