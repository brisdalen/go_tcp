package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

		position := handleInput(input)
		if turnP1 {
			board1D[position] = 1
			turnP1 = !turnP1
		} else {
			board1D[position] = -1
			turnP1 = !turnP1
		}

		displayBoard(board1D)
	}
}

func handleInput(input string) int {
	// i.e. "1,3"
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
		fmt.Print("|")
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func printChar(char string) {
	fmt.Printf("|%2v", char)
}
