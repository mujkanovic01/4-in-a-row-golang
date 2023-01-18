package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
    "math"
    "errors"
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
    var playAgain, loadGameInput string
    for {
        if _, err := os.Stat("savedgame.txt"); errors.Is(err, os.ErrNotExist) {  
            // save does not exist
            setupNewGame()
        } else { 
            // save exists
            fmt.Println("Do you want to load the saved game?(y/n)")
            fmt.Scan(&loadGameInput)
            if (strings.ToLower(loadGameInput) == "y") {
                loadGame()
            } else {
                setupNewGame()
            }
        }
        createNewGame()
        fmt.Println("Do you want to play again?(y/n)")
        fmt.Scan(&playAgain)
        if (strings.ToLower(playAgain) != "y") {
            os.Exit(0)
        }
    }
}

func setupNewGame() {
    for {
        fmt.Print("Enter width of the board: ");
        fmt.Scan(&width);
        fmt.Print("Enter height of the board: ");
        fmt.Scan(&height)

        if (width >= 7 && height >= 6 && math.Abs(float64(width-height)) <= 2) {
            break
        }
    }

	board = make([][]int, height)
	for i := 0; i < int(height); i++ {
		board[i] = make([]int, width)
	}
    
    p1Moves = make([]int, 0)
    p2Moves = make([]int, 0)
    isP1Turn = true
}

func createNewGame() {
    for {
        displayTheBoard()
        move = getNextMove()

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

        if isP1Turn {
            p1Moves = append(p1Moves, move + 1)
        } else {
            p2Moves = append(p2Moves, move + 1)
        }

        if checkIfWon() {
            displayTheBoard()
            fmt.Printf("\nPlayer %d has won the game!\n", getPlayer())
            break
        }

        if checkIfDraw() {
            displayTheBoard()
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
        fmt.Println("Game successfully saved!")
        os.Exit(0)
        return 0
    } else {
        mv, err := strconv.Atoi(move)
        if err != nil {
            fmt.Println("Unkown command")
            return getNextMove()
        }
        return mv - 1
    }
    
}

func displayTheBoard() {
	for i := range board {
		for j := range board[i] {
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

    // Check right diagonal path
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
    if err != nil { return }
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
    boardDataSplitByRow := strings.Split(boardDataSplit[0], " ")

    height = len(boardDataSplit) - 1
    width = len(boardDataSplitByRow)

	board = make([][]int, height)
	for i := 0; i < int(height); i++ {
		board[i] = make([]int, width)
	}

    for i := 0; i < len(boardDataSplit) - 1; i++ {
        boardDataSplitByRow := strings.Split(boardDataSplit[i], " ")
        for j := 0; j < len(boardDataSplitByRow); j++ {
            board[i][j], _ = strconv.Atoi(boardDataSplitByRow[j])
        }
    }

    playerMovesData := strings.Split(string(data), "p1")[1]
    p1MovesStr := strings.Split(strings.Split(strings.Split(playerMovesData, "=[")[1], "\n")[0], " ")
    p2MovesStr := strings.Split(strings.Split(strings.Split(playerMovesData, "p2=[")[1], "\n")[0], " ")
    
    for i := 0; i < len(p1MovesStr); i++ {
        intValue, _ := strconv.Atoi(p1MovesStr[i])
        p1Moves = append(p1Moves, intValue)
    }

    for i := 0; i < len(p2MovesStr); i++ {
        intValue, _ := strconv.Atoi(p2MovesStr[i])
        p2Moves = append(p2Moves, intValue)
    }

    isP1Turn = len(p1Moves) == len(p2Moves)
    os.Remove("savedgame.txt")
}

func checkIfDraw() (bool) {
    for x := 0; x < height; x++ {
        for y := 0; y < width; y++ {
            if board[x][y] == 0 { return false }
        }
    }

    return true
}

func getPlayer() (int) {
    if isP1Turn { return 1 }

    return 2
}