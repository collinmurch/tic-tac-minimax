package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

type board struct {
	state [3][]string
	turn string
}

func main() {
	game := freshBoard()
	reader := bufio.NewReader(os.Stdin)
	var text string

	for {
		show(game)
		
		if winning(game, "X") || winning(game, "O") {	
			fmt.Printf("Winner! Game over.\n")
			break
		}

		fmt.Printf("Enter a move:\n")
		text, _ = reader.ReadString('\n')
		
		loc, err := strconv.ParseInt(text[:1], 10, 64)
		if err != nil {
			panic(err)
		}
		
		if check(game, int(loc)){
			game = update(game, int(loc))
		} else {
			fmt.Printf("Move taken. Try again.\n")
		}
		if !gameOver(game) {
			a, aiMove := minimax(game, 0)
			fmt.Printf("%d\n", a)
			game = update(game, aiMove)
		} else {
			show(game)
			fmt.Printf("It's a draw.\n")
			break
		}
	}
}

func winning(b board, p string) bool { 	
	switch p {
		// Go across diagonal, checking the row and column
		case b.state[0][0]:
			if b.state[0][1] == p && b.state[0][2] == p {
				return true
			}
			if b.state[1][0] == p && b.state[2][0] == p {
				return true
			}
			if b.state[1][1] == p && b.state[2][2] == p {
				return true
			}
		case b.state[1][1]:
			if b.state[1][0] == p && b.state[1][2] == p {
				return true
			}
			if b.state[0][1] == p && b.state[2][1] == p {
				return true
			}
		case b.state[2][2]:
			if b.state[2][0] == p && b.state[2][1] == p {
				return true
			}
			if b.state[0][2] == p && b.state[1][2] == p {
				return true
			}
	}

	// Check other diagonal
	if b.state[0][2] == p && b.state[1][1] == p && b.state[2][0] == p {
		return true
	} else {
		return false
	}
}

func gameOver(b board) bool {
	if winning(b, "X") || winning(b, "O") || len(getMoves(b)) == 0 {
		return true
	} else {
		return false
	}
}

func update(b board, l int) board {		
	newBoard := freshBoard()

	for i := 0; i < len(b.state); i++ {
		for j := 0; j < len(b.state[i]); j++ {
			newBoard.state[i][j] = b.state[i][j]
		}
	}
	newBoard.turn = b.turn
		
	row, loc := getLoc(newBoard, l)
	newBoard.state[row][loc] = newBoard.turn

	if newBoard.turn == "X" {
		newBoard.turn = "O"
	} else {
		newBoard.turn = "X"
	}

	return newBoard
}

func check(b board, l int) bool {	
	moves := getMoves(b)

	// Check if suggested move is in available moves
	for i := 0; i < len(moves); i++ {
		if l == moves[i] {
			return true
		}
	}
	
	return false
}

func getMoves(b board) []int {
	moves := make([]int, 0)

	// Go through array, numbered human-wise (1-9)
	// Append all empty squares
	for i := 1; i < 10; i++ {
		row, loc := getLoc(b, i)
		if b.state[row][loc] == "-" {
			moves = append(moves, i)
		}
	}

	return moves
}

func getLoc(b board, l int) (int, int) {
	var row int
	loc := int(l)

	if loc < 4 {
		loc -= 1
		row = 0
	} else if loc < 7 {
		loc -= 4
		row = 1
	} else if loc < 10 {
		loc -= 7
		row = 2
	}

	return row, loc
}

func show(b board) {
	for i := 0; i < len(b.state); i++ {
		fmt.Printf("%s\n", strings.Join(b.state[i], " "))
	}
	fmt.Printf("******\n")
}

func freshBoard() board {
	return board {
		state : [3][]string {
			[]string{"-", "-", "-"},
			[]string{"-", "-", "-"},
			[]string{"-", "-", "-"},
		},
		turn : "X",
	}

}
