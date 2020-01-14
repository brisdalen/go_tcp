package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var moves int = 0
var winner int

func main() {
	board1D := make([]int, 9)
	for i := 0; i < len(board1D); i++ {
		board1D[i] = 0
	}

	reader := bufio.NewReader(os.Stdin)
	turnP1 := true

	displayBoard(board1D)

	for {
		fmt.Println("Enter x and y coordinate for move, i.e \"1,3\".")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if input == "exit\n" {
			break
		}

		moves++
		position := handleInput(input)
		if turnP1 {
			board1D[position] = 1
			turnP1 = !turnP1
		} else {
			board1D[position] = -1
			turnP1 = !turnP1
		}

		displayBoard(board1D)

		if checkGameStatus(board1D) {
			printWinner(winner)
			break
		}
	}
}

func checkGameStatus(board []int) bool {
	// You can't win before the 5th move
	if moves > 4 {
		// rows
		if abs(board[0]+board[1]+board[2]) == 3 ||
			abs(board[3]+board[4]+board[5]) == 3 ||
			abs(board[6]+board[7]+board[8]) == 3 {
			winner = board[0]
			return true
		}
		//columns
		if abs(board[0]+board[3]+board[6]) == 3 ||
			abs(board[1]+board[4]+board[7]) == 3 ||
			abs(board[2]+board[5]+board[8]) == 3 {
			winner = board[0]
			return true
		}
		// diagonals
		if abs(board[0]+board[4]+board[8]) == 3 ||
			abs(board[2]+board[4]+board[6]) == 3 {
			winner = board[4]
			return true
		}
	}

	return false
}

func printWinner(winnerID int) {
	var winner string
	if winnerID == 1 {
		winner = "1"
	} else {
		winner = "2"
	}
	fmt.Println("The winner is Player", winner, "\b!")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func handleInput(input string) int {
	parts := strings.Split(input, ",")

	s := parts[:]

	v1 := strings.TrimSpace(s[0])
	v2 := strings.TrimSpace(s[1])

	c1, err := strconv.Atoi(v1)
	if err != nil {
		log.Fatal(err)
	}

	c2, err2 := strconv.Atoi(v2)
	if err2 != nil {
		log.Fatal(err2)
	}
	// convert 1-index to 0-index and calculate slice index
	pos := calculateIndex(c1-1, c2-1)

	return pos
}

func calculateIndex(column, row int) int {
	// [1] + (width * [0]);
	return row + (3 * column)
}

func displayBoard(board []int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := board[calculateIndex(i, j)]

			if v == 1 {
				printChar("O")
			} else if v == -1 {
				printChar("X")
			} else {
				printChar("_")
			}
			fmt.Print(" ")
		}
		fmt.Print("|\n")
	}
	fmt.Print("\n")
}

func printChar(char string) {
	fmt.Printf("|%2v", char)
}
