package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
    board [][]int
    p1Moves []int
    p2Moves []int
    move int
    width int
    height int
    isP1Turn bool
)
func main() {
    createNewGame()

    // TODO: Play again functionality
}

func createNewGame() {
    for {
        fmt.Print("Enter width of the board: ");
        fmt.Scan(&width);
        fmt.Print("Enter height of the board: ");
        fmt.Scan(&height)

        if (width >= 7 && height >= 6) {break}
    }

    // fmt.Print(width, height)
    // TODO: Implement checks for max diff 2 and min size 6x8

	board = make([][]int, height)
	for i := 0; i < int(height); i++ {
		board[i] = make([]int, width)
	}
    
    p1Moves = make([]int, 0)
    p2Moves = make([]int, 0)
    isP1Turn = true
    for {
        printBoard()
        
        move = getNextMove() // Width passed through for user-message purpose

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
				board[i][move] = getPlayer()
				break
			}
		}

        // Track the moves
        if isP1Turn {
            p1Moves = append(p1Moves, move + 1)
        } else {
            p2Moves = append(p2Moves, move + 1)
        }

        if checkIfWon() {
            printBoard()
            fmt.Printf("\nPlayer %d has won the game!\n", getPlayer())
            break
        }

        if checkIfDraw() {
            printBoard()
            fmt.Println("\nGame is Drawn!")
            break
        }

        isP1Turn = !isP1Turn
    }
}

func getNextMove() (int) {
    var move string
    fmt.Printf("Enter column number (1 - %d)\n", width)
    fmt.Scan(&move)

    if strings.ToLower(move) == "save"{
        saveGame(board, p1Moves, p2Moves)
        os.Exit(0)
        return 0
    } else if strings.ToLower(move) == "load"{
        loadGame()
        return getNextMove()
    } else {
        m, err := strconv.Atoi(move)
        if err != nil {
            fmt.Println("Unkown command")
            return getNextMove()
        }
        return m - 1
    }
    
}

func printBoard() {
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

func checkIfWon() (bool) {
    movePlayed := getPlayer()

    // Check horizontal path
    for y := 0; y < width; y++ {
        for x := 0; x < height - 3; x++ {
            if board[x][y] == movePlayed && board[x+1][y] == movePlayed && board[x+2][y] == movePlayed && board[x+3][y] == movePlayed {
                return true
            }
        }
    }

    // Check vertical path
    for x := 0; x < height; x++ {
        for y := 0; y < width - 3; y++ {
            if board[x][y] == movePlayed && board[x][y+1] == movePlayed && board[x][y+2] == movePlayed && board[x][y+3] == movePlayed {
                return true
            }
        }
    }

    // check right diagonal path
    for x := 0; x < height - 3; x++ {
        for y := 3; y < width; y++ {
            if board[x][y] == movePlayed && board[x+1][y-1] == movePlayed && board[x+2][y-2] == movePlayed && board[x+3][y-3] == movePlayed {
                return true
            }
        }
    }

    // Check left diagonal path
    for x := 0; x < height - 3; x++{
        for y := 0; y < width - 3; y++ {
            if board[x][y] == movePlayed && board[x+1][y+1] == movePlayed && board[x+2][y+2] == movePlayed && board[x+3][y+3] == movePlayed {
                return true
            }
        }
    }

    return false
}

func saveGame(board [][]int, p1Moves []int, p2Moves []int) {
    file, err := os.Create("savedgame.txt")
    if err != nil {
        return
    }
    defer file.Close()

    for i := 0; i < height; i++ {
        file.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(board[i])), " "), "[]") + "\n")
    }
    file.WriteString(strings.Trim(strings.Join(strings.Fields("p1=" + fmt.Sprint(p1Moves)), " "), "[]") + "\n")
    file.WriteString(strings.Trim(strings.Join(strings.Fields("p2=" + fmt.Sprint(p2Moves)), " "), "[]") + "\n")
}

func loadGame() {
    data, err := os.ReadFile("savedgame.txt")
    if err != nil { return }

    boardData := strings.Split(string(data), "p1")[0]
    boardDataSplit := strings.Split(boardData, "\n")
    height = len(boardDataSplit) - 1

    for i := 0; i < len(boardDataSplit) - 1; i++ {
        boardDataSplitByRow := strings.Split(boardDataSplit[i], " ")
        for j := 0; j < len(boardDataSplitByRow); j++ {
            width = len(boardDataSplitByRow)
            board[i][j], _ = strconv.Atoi(boardDataSplitByRow[j])
        }
    }

    playerMovesData := strings.Split(string(data), "p1")[1]
    p1MovesStr := strings.Split(strings.Split(strings.Split(playerMovesData, "=[")[1], "p2")[0], " ")
    p2MovesStr := strings.Split(strings.TrimLeft(strings.Split(playerMovesData, "p2=[")[1], " "), " ")
    
    for i := 0; i < len(p1MovesStr); i++ {
        intValue, _ := strconv.Atoi(p1MovesStr[i])
        p1Moves = append(p1Moves, intValue)
    }

    for i := 0; i < len(p2MovesStr); i++ {
        intValue, _ := strconv.Atoi(p2MovesStr[i])
        p2Moves = append(p2Moves, intValue)
    }

    fmt.Println(p1MovesStr)
    fmt.Println(p2MovesStr)
}

func checkIfDraw() (bool) {
    for x := 0; x < height; x++ {
        for y := 0; y < width; y++ {
            if board[x][y] == 0 {return false}
        }
    }

    return true
}

func getPlayer() (int) {
    if isP1Turn { 
        return 1 
    }
    return 2
}